package api

import (
	"fmt"
	"log"
	"encoding/json"
	"strings"
)

const URL_TIME_HISTORICO_ID = "/time/id/%d/%d"

type TimeHistorico TimeCompleto

func (t *TimeHistorico) GetTime(){
	request := Request{}
	res, err := request.Get(t.MountUrl(), 10)

	if err != nil {
		log.Println(err)
	}
	if res.StatusCode() != 200 {
		log.Println("time id diferente de 200", res.StatusCode(), "response: ", string(res.Body()))
	}else{
		json.Unmarshal(res.Body(), &t)
		t.ChangeFormatDefault()
	}
}

func (t *TimeHistorico) MountUrl() string {
	return fmt.Sprintf(URL_TIME_HISTORICO_ID, t.TimeCompleto.TimeId, t.RodadaAtual)
}

func (a *TimeHistorico) ChangeFormatDefault() {
	for i, des := range a.Atletas {
		a.Atletas[i].Foto = strings.Replace(des.Foto, "FORMATO", "140x140", 3)
	}
}