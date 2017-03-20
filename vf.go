package httpvf

import (
	"github.com/toukii/goutils"

	"bytes"
	"fmt"
	"github.com/toukii/jsnm"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

func verify(req *Req) *Msg {
	msg := newMsg(req)
	var resp *http.Response
	var request *http.Request
	var err error
	// fmt.Printf("[%s]%s\n", req.Method, req.URL)
	if len(req.Upload) > 0 {
		splt := strings.Split(req.Upload, "@")
		tag := "filename"
		filename := splt[0]
		if len(splt) > 1 {
			tag = splt[0]
			filename = splt[1]
		}
		request, err = newfileUploadRequest(req.URL, nil, tag, filename)
		if goutils.CheckErr(err) {
			msg.Append(FATAL, err.Error())
			buf := reqBody(req.Body)
			request, err = http.NewRequest(req.Method, req.URL, buf)
		}
	} else {
		buf := reqBody(req.Body)
		request, err = http.NewRequest(req.Method, req.URL, buf)
	}
	if goutils.CheckErr(err) {
		msg.Append(FATAL, err.Error())
	}

	for k, v := range req.Header {
		request.Header.Add(k, v)
	}

	//  start
	start := time.Now()
	c := http.Client{}
	resp, _ = c.Do(request)

	// end
	duration := time.Now().Sub(start)

	cost := int(duration.Nanoseconds() / 1e6)
	if req.Resp.Cost != 0 {

		if cost > req.Resp.Cost {
			msg.Append(ERROR, fmt.Sprintf("time cost: %d ms more than %d ms;", cost, req.Resp.Cost))
		} else if cost > req.Resp.Cost*3/4 {
			msg.Append(WARN, fmt.Sprintf("time cost: %d ms near by %d ms;", cost, req.Resp.Cost))
		} else {
			msg.Append(INFO, fmt.Sprintf("time cost: %d ms / %d ms;", cost, req.Resp.Cost))
		}
	}
	msg.Req.Resp.RealCost = cost
	if resp == nil {
		msg.Append(ERROR, "nil response")
	} else {
		if req.Resp.Code != 0 && req.Resp.Code != resp.StatusCode {
			msg.Append(ERROR, fmt.Sprintf("error code::%d gotten, %d wanted", resp.StatusCode, req.Resp.Code))
		}
		bs, respReadErr := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if len(req.Resp.Body) > 0 {
			if goutils.CheckErr(respReadErr) {
				msg.Append(ERROR, respReadErr.Error())
			}
			if !strings.EqualFold(req.Resp.Body, goutils.ToString(bs)) {
				msg.Append(ERROR, fmt.Sprintf("response body is: %s, not wanted: %s", goutils.ToString(bs), req.Resp.Body))
			}
		}

		if len(req.Resp.Regex) > 0 {
			if matched, errg := regexp.Match(req.Resp.Regex, bs); !matched || goutils.LogCheckErr(errg) {
				msg.Append(ERROR, fmt.Sprintf("response body is: %s, not wanted regexp: %s", goutils.ToString(bs), req.Resp.Regex))
			}
		}
		if len(req.Resp.Json) > 0 {
			vfJson(bs, req.Resp.Json, msg)
		}
	}
	return msg
}

func reqBody(body string) *bytes.Buffer {
	if strings.HasPrefix(body, "@") {
		return bytes.NewBuffer(goutils.ReadFile(strings.TrimPrefix(body, "@")))
	}
	return bytes.NewBufferString(body)
}

func vfJson(bs []byte, kvs map[string]string, msg *Msg) {
	js := jsnm.BytesFmt(bs)
	for ks, wv := range kvs {
		k := js.ArrGet(strings.Split(ks, ",")...).RawData().String()
		if k != wv {
			msg.Append(ERROR, fmt.Sprintf("response body: <%s> is goten, <%s> is wanted.", k, wv))
		}
	}
}

func verifys(reqs []*Req, isSync bool) {
	if len(reqs) <= 0 {
		return
	}
	var wg sync.WaitGroup
	tickerMap := make(map[string]*time.Ticker)
	runtineMap := make(map[string]chan struct{})
	for _, it := range reqs {
		it.Prapare()
		// fmt.Println("***", it.URL)
		if it.Interval > 0 {
			ticker := time.NewTicker(time.Duration(it.Interval * 1e6))
			tickerMap[it.MapKey()] = ticker
			// fmt.Println(it.MapKey(), it.Interval)
		}
		runtineMap[it.MapKey()] = make(chan struct{}, it.Runtine)
		wg.Add(1)
		if isSync {
			i := 0
			cost := 0
			var tps string
			logs := make([]*Log, 0, 64)
			index := make(chan struct{}, 1)
			for {
				go func() {
					index <- struct{}{}
					i++
					<-index
					msg := verify(it)
					cost += msg.Req.Resp.RealCost
					logs = append(logs, msg.Logs()...)
					if i >= it.N {
						fmt.Println()
						tps += fmt.Sprint("avg cost: ", cost/i, " ms")
						msg = newMsg(it)
						msg.Append(CONCLUSION, tps)
						msg.AppendLogs(logs)
						fmt.Println(msg)
					}
					runtineMap[it.MapKey()] <- struct{}{}
				}()
				<-runtineMap[it.MapKey()]
				verifys(it.Then, it.Sync)
				if i >= it.N {
					break
				}
				if ticker, ok := tickerMap[it.MapKey()]; it.Interval > 0 && ok {
					<-ticker.C
				}
			}
			wg.Done()
		} else {
			go func(it *Req) {
				i := 0
				cost := 0
				var tps string
				logs := make([]*Log, 0, 64)
				index := make(chan struct{}, 1)
				for {
					go func() {
						index <- struct{}{}
						i++
						<-index
						msg := verify(it)
						cost += msg.Req.Resp.RealCost
						logs = append(logs, msg.Logs()...)
						if i >= it.N {
							fmt.Println()
							tps += fmt.Sprint("avg cost: ", cost/i, " ms")
							msg = newMsg(it)
							msg.Append(CONCLUSION, tps)
							msg.AppendLogs(logs)
							fmt.Println(msg)
						}
						runtineMap[it.MapKey()] <- struct{}{}
					}()
					<-runtineMap[it.MapKey()]
					verifys(it.Then, it.Sync)
					if i >= it.N {
						break
					}
					if ticker, ok := tickerMap[it.MapKey()]; it.Interval > 0 && ok {
						<-ticker.C
					}
				}
				wg.Done()
			}(it)
		}
	}
	wg.Wait()
}

func Verify(vf string) {
	reqs, errvf := Reqs(vf)
	if goutils.LogCheckErr(errvf) {
		return
	}
	verifys(reqs, false)
}

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	return req, err
}
