package httpvf

import (
	"fmt"
	"github.com/toukii/goutils"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"strings"
)

type Resp struct {
	Code     int
	Cost     int
	RealCost int
	Body     string
	Regex    string // regex body
	Json     map[string]string
}

const (
	GET  = "GET"
	POST = "POST"
)

type Req struct {
	N        int
	URL      string
	Method   string
	Header   map[string]string
	Param    map[string]string
	Body     string
	Resp     Resp
	Upload   string
	Interval int
	Runtine  int
	Timeout  int64
	Then     []*Req
	Sync     bool // 根请求（第一个请求）为异步模式，默认同一层请求为异步模式
}

func (req *Req) Do() ([]byte, error) {
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
			buf := reqBody(req.Body)
			request, err = http.NewRequest(req.Method, req.URL, buf)
		}
	} else {
		buf := reqBody(req.Body)
		request, err = http.NewRequest(req.Method, req.URL, buf)
	}
	if goutils.CheckErr(err) {
		return nil, err
	}

	for k, v := range req.Header {
		request.Header.Add(k, v)
	}

	c := http.Client{}
	resp, err = c.Do(request)
	if goutils.CheckErr(err) {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func Reqs(filename string) (reqs []*Req, err error) {
	in := goutils.ReadFile(filename)
	reqs = make([]*Req, 0, 1)
	if len(in) > 0 && in[0] != byte('-') {
		var req1 *Req
		req1, err = ReqFmt(in)
		if goutils.CheckErr(err) {
			return nil, err
		}
		reqs = append(reqs, req1)
		return
	}
	err = yaml.Unmarshal(in, &reqs)
	if goutils.CheckErr(err) {
		return nil, err
	}
	return reqs, nil
}

func ReqFmt(in []byte) (*Req, error) {
	var req Req
	err := yaml.Unmarshal(in, &req)
	if goutils.CheckErr(err) {
		return nil, err
	}

	return &req, nil
}

func (req *Req) Prapare() {

	if len(req.Param) > 0 {
		ss := []string{}
		for k, v := range req.Param {
			ss = append(ss, fmt.Sprintf("%s=%s", k, v))
		}
		query := strings.Join(ss, "&")

		if strings.Contains(req.URL, "?") {
			req.URL += "&" + query
		} else {
			req.URL += "?" + query
		}
	}
	if req.Runtine <= 0 {
		req.Runtine = 1
	}
	if req.N <= 0 {
		req.N = 1
	}
}

func (req Req) MapKey() string {
	return fmt.Sprintf("%s-%s", req.Method, req.URL)
}

func Test() {
	var req Req
	req.URL = "http://upload.daoapp.io/upload/a.json"
	req.Body = fmt.Sprintf(`{"name":"toukii"}`)
	req.Resp.Body = "world"
	req.Header = make(map[string]string)
	req.Header["Content-Type"] = "application/json"
	req.Method = GET
	req.Resp.Json = map[string]string{"v": "p1,p2,0,p31"}
	reqs := []Req{req, req}
	bs2, err := yaml.Marshal(reqs)
	goutils.CheckErr(err)
	fmt.Println(goutils.ToString(bs2))

	reqs2 := []Req{}
	err = yaml.Unmarshal(bs2, &reqs2)
	goutils.CheckErr(err)
	fmt.Println(reqs2)

}
