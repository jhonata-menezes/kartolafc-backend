package api

import (
	"log"
	"encoding/json"
)

const URL_LIGAS = "/ligas?q="

type Ligas struct {
		LigaId int `json:"liga_id"`
		Nome string `json:"nome"`
		Descricao string `json:"descricao"`
		Slug string `json:"slug"`
		Imagem string `json:"imagem"`
		QuantidadeTimes int `json:"quantidade_times"`
		MataMata bool `json:"mata_mata"`
		Tipo string `json:"tipo"`
}

type PesquisaLigas []Ligas

func (l *PesquisaLigas) GetPesquisaLigas(nome string) {
	request := Request{}
	res, err := request.Get(URL_LIGAS + UrlEncode(nome), 10)
	if err != nil {
		log.Println("ligas", err)
	}
	if res.StatusCode() != 200 {
		log.Println("ligas status ", res.StatusCode())
	}else{
		json.Unmarshal(res.Body(), &l)
	}

}


