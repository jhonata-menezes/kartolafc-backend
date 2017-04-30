package kartola

import (
	"github.com/valyala/fasthttp"
	"time"
)

type Request struct{}

func (r *Request) Get(url string, timeout int) (*fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36")
	req.URI().Update(url)
	req.Header.SetMethod("GET")
	res := fasthttp.AcquireResponse()
	client := fasthttp.Client{}
	err := client.DoTimeout(req, res, time.Duration(timeout) * time.Second)
	return res, err
}

