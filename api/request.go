package api

import (
	"github.com/valyala/fasthttp"
	"time"
	"net/url"
	"github.com/jhonata-menezes/kartolafc-backend/cmd"
	"encoding/json"
	"log"
)

const BASE_URL_API = "https://api.cartolafc.globo.com"

type Request struct{}

func (r *Request) Get(uri string, timeout int) (*fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://cartolafc.globo.com/")
	req.Header.Set("Origin", "https://cartolafc.globo.com/")
	req.Header.Set("Accept-Language", "pt-BR,pt;q=0.8,en-US;q=0.6,en;q=0.4,es;q=0.2")
	req.Header.Set("X-GLB-Token", cmd.Config.TokenGlb)
	req.URI().Update(BASE_URL_API + uri)
	res := fasthttp.AcquireResponse()
	client := fasthttp.Client{}
	err := client.DoTimeout(req, res, time.Duration(timeout) * time.Second)
	return res, err
}

func (r *Request) GetSimple(url string, timeout int) (*fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36")
	req.Header.Set("Accept-Language", "pt-BR,pt;q=0.8,en-US;q=0.6,en;q=0.4,es;q=0.2")
	req.URI().Update(url)
	res := fasthttp.AcquireResponse()
	client := fasthttp.Client{}
	err := client.DoTimeout(req, res, time.Duration(timeout) * time.Second)
	return res, err
}

func (r *Request) Post(url string, form interface{}, timeout int) (*fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36")
	req.Header.Set("Accept-Language", "pt-BR,pt;q=0.8,en-US;q=0.6,en;q=0.4,es;q=0.2")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.SetMethod("post")
	formJson, _ := json.Marshal(form)
	log.Printf("%#v", form)
	req.SetBody(formJson)
	req.URI().Update(url)
	res := fasthttp.AcquireResponse()
	client := fasthttp.Client{}
	err := client.DoTimeout(req, res, time.Duration(timeout) * time.Second)
	return res, err
}

func UrlEncode(path string) string {
	u := &url.URL{Path: path}
	return u.String()
}

