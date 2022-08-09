package pay

import (
	"crypto/tls"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"time"
)

var ct = client()

func client() fasthttp.Client {
	return fasthttp.Client{
		Name:                     "2pay",
		NoDefaultUserAgentHeader: true,
		TLSConfig:                &tls.Config{InsecureSkipVerify: true},
		MaxConnsPerHost:          2000,
		MaxIdleConnDuration:      5 * time.Second,
		MaxConnDuration:          5 * time.Second,
		ReadTimeout:              5 * time.Second,
		WriteTimeout:             5 * time.Second,
		MaxConnWaitTimeout:       5 * time.Second,
	}
}

func PostJson(url string, args map[string]interface{}) ([]byte, error) {

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.SetRequestURI(url)
	marshal, _ := json.Marshal(args)
	req.SetBody(marshal)

	if err := ct.Do(req, resp); err != nil {
		return nil, err
	}

	return resp.Body(), nil
}
