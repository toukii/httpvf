package httpvf

import (
	"github.com/astaxie/beego/httplib"
	"github.com/toukii/goutils"

	"fmt"
	"time"
	//"io/ioutil"
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func Verify(req *Req) *Msg {
	msg := newMsg(req)
	var resp *http.Response
	var err error
	beegoReq := httplib.NewBeegoRequest(req.URL, req.Method)

	//  start
	start := time.Now()
	if len(req.Filename) > 0 {
		request, err := newfileUploadRequest(req.URL, nil, "filename", req.Filename)
		//request.Header.Set("Content-Type","")
		if !goutils.CheckErr(err) {
			c := http.Client{}
			resp, _ = c.Do(request)
		}
		//beegoReq.PostFile("filename",req.Filename)
		goto END
	}
	resp, err = beegoReq.DoRequest()
END:
	if goutils.CheckErr(err) {
		msg.Append(FATAL, err.Error())
	}

	// end
	duration := time.Now().Sub(start)

	// fmt.Printf("Request[%s] cost:%d ms\n", req.URL, duration.Nanoseconds()/1e6)
	cost := int(duration.Nanoseconds() / 1e6)
	if cost > req.Resp.Cost {
		msg.Append(ERROR, fmt.Sprintf("time cost: %d ms more than %d ms;", cost, req.Resp.Cost))
	} else if cost > req.Resp.Cost*3/4 {
		msg.Append(WARN, fmt.Sprintf("time cost: %d ms near by %d ms;", cost, req.Resp.Cost))
	} else {
		msg.Append(INFO, fmt.Sprintf("time cost: %d ms / %d ms;", cost, req.Resp.Cost))
	}
	if resp == nil {
		msg.Append(ERROR, "nil response")
	} else if req.Resp.Code != resp.StatusCode {
		msg.Append(ERROR, fmt.Sprintf("error code::%d gotten, %d wanted", resp.StatusCode, req.Resp.Code))
	}
	return msg
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
