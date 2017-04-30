package api

import (
	"github.com/valyala/fasthttp"
	"time"
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
	req.URI().Update(BASE_URL_API + uri)
	res := fasthttp.AcquireResponse()
	client := fasthttp.Client{}
	err := client.DoTimeout(req, res, time.Duration(timeout) * time.Second)
	return res, err
}

//X-GLB-Token: 1af163e2480eeffb9d494c99c62c2b4ad7339584269786f375a4e414252347939474b537942363545624574574d313249616d695a4243365452474d53425559684163686f555671327849446c5a314236504f35704b50734b63714975796659463341334170413d3d3a303a75776b6e6e666d7175796f66656b707463666d71