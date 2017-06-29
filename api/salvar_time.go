package api

import "encoding/json"

const URI_SALVAR_TIME = "/auth/time/salvar"

type SalvarTime struct {
	Esquema int `json:"esquema,omitempty"`
	Atleta []int `json:"atleta,omitempty"`
	Mensagem string `json:"mensagem,omitempty"`
}

func (s *SalvarTime) Post(token string) {
	request := Request{}
	form, _ := json.Marshal(s)
	resp, _ := request.PostToken(URI_SALVAR_TIME, 10, form, token)
	if (resp.StatusCode() != 200) {
		s.Mensagem = "Houve algum problema e seu time não foi escalado. Tente novamente!"
		return
	}
	json.Unmarshal(resp.Body(), s)
}
