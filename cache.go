package kartolafc

import (
	"github.com/jhonata-menezes/kartolafc-backend/api"
	"time"
	"log"
)

var CacheKartolaAtletas api.Atletas
var CacheStatus api.Status
var CacheDestaques api.Destaques
var CachePontuados api.Pontuados
var CacheRankingPontuados TimesRankingFormated
var CacheRankingPontuadosMelhores []TimeRankingFormated
var CacheRankingTimeIdPontuados []TimeIdRanking
var CachePartidas []api.Partidas

func UpdateStatus() {
	status := api.Status{}
	status.GetStatus()
	if status.RodadaAtual != 0 {
		CacheStatus = status
	}
	SleepCacheSecond(60)
	UpdateStatus()
}

func UpdateDestaques() {
	destaques := api.Destaques{}
	destaques.GetDestaques()
	if len(destaques) > 0 {
		//log.Printf("%#v", destaques)
		CacheDestaques = destaques
	}

	SleepCacheSecond(60)
	UpdateDestaques()
}

func UpdateMercado() {
	mercado := api.Atletas{}
	mercado.GetAtletas()
	if len(mercado.Atleta) > 0 {
		CacheKartolaAtletas = mercado
	}

	SleepCacheSecond(60)
	UpdateMercado()
}

func UpdatePontuados() {
	pontuados := api.Pontuados{}
	pontuados.GetPontuados()

	if pontuados.Rodada != 0 {
		CachePontuados = pontuados
	}

	SleepCacheSecond(60)
	UpdatePontuados()
}

func UpdatePartidas() {
	CachePartidas = make([]api.Partidas, 21)

	// se pegou todas rodadas anteriores, atualiza apenas a rodada atual
	if CachePartidas[0].Rodada > 0 {
		tmp := api.Partidas{}
		tmp.Get(0)

		// atualiza o cache da rodada 0 e da rodada retornada
		if tmp.Rodada > 0 {
			CachePartidas[0] = tmp
			CachePartidas[tmp.Rodada] = tmp
		} else {
			log.Println("rodada atual retornada como 0")
		}
		CachePartidas[0] = tmp

		SleepCacheSecond(10)
		UpdatePartidas()
	}


	for i:=0; i<=20; i++ {
		tmp := api.Partidas{}
		if i == 0 {
			tmp.Get(i)
			CachePartidas[i] = tmp
			continue
		}
		tmp.Get(i)
		CachePartidas[i] = tmp
	}

	SleepCacheSecond(3600)
	UpdatePartidas()
}

func UpdateCache() {
	go UpdateStatus()
	go UpdateDestaques()
	go UpdateMercado()
	go UpdatePontuados()
	go UpdatePartidas()
}

func SleepCacheSecond(t int) {
	time.Sleep(time.Duration(t) * time.Second)
}
