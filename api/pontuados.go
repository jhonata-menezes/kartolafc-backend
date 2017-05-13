package api

import (
	"encoding/json"
	"strings"
	"log"
)

const URL_PONTUADOS = "/atletas/pontuados"

type Pontuado struct {
	Apelido string `json:"apelido"`
	Pontuacao float32 `json:"pontuacao"`
	Scout map[string]int `json:"scout"`
	Foto string `json:"foto"`
	PosicaoId int `json:"posicao_id"`
	ClubeId int `json:"clube_id"`
}

type Clube struct {
	Id int `json:"id"`
	Nome string `json:"nome"`
	Abreviacao string `json:"abreviacao"`
	Escudos map[string]string
}

type Posicoes struct {
	Id int `json:"id"`
	Nome string `json:"nome"`
	Abreviacao string `json:"abreviacao"`
}

type Pontuados struct {
	Rodada int `json:"rodada"`
	Atletas map[int]Pontuado `json:"atletas"`
	Clubes map[int]Clube `json:"clubes"`
	Posicoes map[int]Posicoes `json:"posicoes"`
	TotalAtletas int `json:"total_atletas"`
}

func (p *Pontuados) GetPontuados() {
	request := Request{}
	res, err := request.Get(URL_PONTUADOS, 10)
	if err != nil {
		log.Println(err)
	}

	if err = json.Unmarshal(res.Body(),&p); err != nil {
		log.Println("json pontuados", err)
	}else{
		p.ChangeFormatDefault()
	}

}

func (p *Pontuados) ChangeFormatDefault() {
	for i, des := range p.Atletas {
		foto := p.Atletas[i]
		foto.Foto = strings.Replace(des.Foto, "FORMATO", "140x140", 3)
		p.Atletas[i] = foto
	}
}
