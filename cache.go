package kartolafc

import (
	"github.com/jhonata-menezes/kartolafc-backend/api"
	"time"
)

var CacheKartolaAtletas api.Atletas
var CacheStatus api.Status
var CacheDestaques api.Destaques


func UpdateStatus() {
	status := api.Status{}
	status.GetStatus()
	if status.RodadaAtual != 0 {
		CacheStatus = status
	}
	SleepCacheSecond(60)
}

func UpdateDestaques() {
	destaques := api.Destaques{}
	destaques.GetDestaques()
	if len(destaques) > 0 {
		CacheDestaques = destaques
	}

	SleepCacheSecond(60)
}

func UpdateMercado() {
	mercado := api.Atletas{}
	mercado.GetAtletas()
	if len(mercado.Atleta) > 0 {
		CacheKartolaAtletas = mercado
	}

	SleepCacheSecond(60)
}

func UpdateCache() {
	go UpdateStatus()
	go UpdateDestaques()
	go UpdateMercado()
}

func SleepCacheSecond(t int) {
	time.Sleep(time.Duration(t) * time.Second)
}
