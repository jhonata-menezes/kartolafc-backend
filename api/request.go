package api

import (
	"net/url"
	"github.com/jhonata-menezes/kartolafc-backend/cmd"
	"github.com/parnurzeal/gorequest"
)

const BASE_URL_API = "https://api.cartolafc.globo.com"

type Request struct{}

func (r *Request) Get(uri string, timeout int) (gorequest.Response, error) {
	request := gorequest.New()
	resp, _, errs := request.Get(BASE_URL_API + uri).Set("Accept", "application/json, text/plain, */*").
		Set("Referer", "https://cartolafc.globo.com/").
		Set("Origin", "https://cartolafc.globo.com/").
		Set("Accept-Language", "pt-BR,pt;q=0.8,en-US;q=0.6,en;q=0.4,es;q=0.2").
		Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36").
		Set("X-GLB-Token", cmd.Config.TokenGlb).End()
	if len(errs) > 0 {
		return resp, errs[0]
	}
	return resp, nil
}

func UrlEncode(path string) string {
	u := &url.URL{Path: path}
	return u.String()
}