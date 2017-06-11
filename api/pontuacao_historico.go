package api

import (
	"fmt"
	"log"
	"encoding/json"
)

const URL_HISTORICO_RODADA = "/auth/mercado/atleta/%d/pontuacao"

type PontuacaoHistorico []struct {
	AtletaId  int `json:"atleta_id"`
	RodadaId int `json:"rodada_id"`
	Pontos float32 `json:"pontos"`
	Preco float32 `json:"preco"`
	Variacao float32 `json:"variacao"`
	Media float32 `json:"media"`
}

func (p *PontuacaoHistorico) Get(atletaId int) {
	request := Request{}
	res, err := request.Get(p.MountUrl(atletaId), 10)

	if err != nil || res.StatusCode() != 200 {
		//log.Println("erro na requisicao de historico", atletaId, err)
		return
	}

	if err = json.Unmarshal(res.Body(), &p); err != nil {
		log.Println("nao decodificou o json do historico", atletaId)
	}

}

func (p *PontuacaoHistorico) MountUrl(atletaId int) string {
	return fmt.Sprintf(URL_HISTORICO_RODADA, atletaId)
}
