package api

import (
	"log"
	"encoding/json"
)

const URL_LOGIN = "https://login.globo.com/api/authentication"

type Login struct {
	Id string `json:"id"`
	UserMessage string `json:"userMessage"`
	GblId string `json:"glbId"`
}

type payload struct {
	Payload struct {
		Email string `json:"email"`
		Password string `json:"password"`
		ServiceId int `json:"serviceId"`
	} `json:"payload"`
	Captcha string `json:"captcha"`
}


func (l *Login) Login(email, pass string) {
	f := payload{}
	f.Payload.Email = email
	f.Payload.Password = pass
	f.Payload.ServiceId = 438

	request := Request{}
	res, err := request.Post(URL_LOGIN, f, 10)
	if err != nil {
		log.Println(err)
	} else {
		json.Unmarshal(res.Body(), &l)
	}
}
