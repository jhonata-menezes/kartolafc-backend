package cmd

import (
	"io/ioutil"
	"log"
	"encoding/json"
)

type configTemplate struct {
	ServerBind string `json:"bind"`
	MongoDB string `json:"mongodb"`
	VapidPrivate string `json:"vapidPrivate"`
	BotKey string `json:"botKey"`
	BotIdClient int `json:"botClientId"`
	JobsNotification int `json:"jobsNotification"`
}

var Config configTemplate

func init() {
	//"Kartola FC é um wrapper da API do cartolafc"
	c, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Println("arquivo de configuraçao nao existe")
	}

	err = json.Unmarshal(c, &Config)
	if err != nil {
		panic(err)
	}
}

