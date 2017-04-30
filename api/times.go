package api

import (
	"fmt"
	"log"
	"encoding/json"
	"net/url"
)

type Times struct {
	Pesquisa string `json:"pesquisa"`
	Times []Time `json:"times"`
}

const URL_TIMES_PESQUISA = "/times?q=%s"

func (t *Times) GetTimes(){
	request := Request{}
	res, err := request.Get(t.MountUrl(), 20)

	if err != nil {
		log.Println("pesquisa de times", t.Pesquisa, err)
	}
	if res.StatusCode() != 200{
		log.Println("pesquisa diferente de 200", res.StatusCode(), "response:", string(res.Body()))
	}else{
		json.Unmarshal(res.Body(), &t.Times)
	}
}

func (t *Times) MountUrl() string{
	pesquisaEncoded := &url.URL{Path: t.Pesquisa}
	return fmt.Sprintf(URL_TIMES_PESQUISA, pesquisaEncoded.String())
}
