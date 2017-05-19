package api

import (
	"log"
	"encoding/json"
	"strings"
)

const URL_ATLETAS = "/atletas/mercado"

type Atleta struct {
	Nome string `json:"nome"`
	Apelido string `json:"apelido"`
	Foto string `json:"foto"`
	AtletaId int `json:"atleta_id"`
	RodadaId int `json:"rodada_id"`
	ClubeId int `json:"clube_id"`
	PosicaoId int `json:"posicao_id"`
	StatusId int `json:"status_id"`
	PontosNum int `json:"pontos_num"`
	PrecoNum int `json:"preco_num"`
	VariacaoNum int `json:"variacao_num"`
	MediaNum int `json:"media_num"`
	JogosNum int `json:"jogos_num"`
	Scout map[string]int `json:"scout"`
}

type Atletas struct{
	Atleta []Atleta `json:"atletas"`
}

func (a *Atletas) GetAtletas(){
	request := Request{}
	res, err := request.Get(URL_ATLETAS, 10)
	if err != nil {
		log.Println("atletas", err)
	}
	if res.StatusCode() != 200 {
		log.Println("atletas status:", res.StatusCode())
	}else{
		json.Unmarshal(res.Body(), &a)
		a.ChangeFormatDefault()
	}
}

func (a *Atletas) ChangeFormatDefault() {
	for i, des := range a.Atleta {
		a.Atleta[i].Foto = strings.Replace(des.Foto, "FORMATO", "140x140", 3)
	}
}
