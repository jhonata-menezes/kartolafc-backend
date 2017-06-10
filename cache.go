package kartolafc

import (
	"github.com/jhonata-menezes/kartolafc-backend/api"
	"time"
	"log"
	"gopkg.in/mgo.v2"
)

var CacheKartolaAtletas api.Atletas
var CacheStatus api.Status
var CacheDestaques api.Destaques
var CachePontuados api.Pontuados
var CacheRankingPontuados TimesRankingFormated
var CacheRankingPontuadosMelhores []TimeRankingFormated
var CacheRankingPontuadosMelhoresPro []TimeRankingFormated
var CacheRankingTimeIdPontuados []TimeIdRanking
var CachePartidas []api.Partidas
var CacheHistoricoAtleta = make([]api.PontuacaoHistorico, 100000)

// collection para query de time
var ChannelCollectionTime = make(chan *mgo.Collection, 2000)

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
	// para nao ficar enviando muitas requisicoes, ira apenas atualizar quando o mercado estiver fechado
	for (CacheStatus.StatusMercado != 2 && CacheStatus.StatusMercado != 0) {
		SleepCacheSecond(60)
	}

	pontuados := api.Pontuados{}
	pontuados.GetPontuados()

	if pontuados.Rodada != 0 {
		CachePontuados = pontuados
	}

	// verificar se tem jogo rolando para diminuir o tempo de requisicao
	if TemJogoComPartidaRolando() {
		SleepCacheSecond(10)
	} else {
		SleepCacheSecond(60)
	}
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

	SleepCacheSecond(30)
	UpdatePartidas()
}

func UpdateHistoricoPontuacao() {
	SleepCacheSecond(10)
	for _, d := range CacheKartolaAtletas.Atleta {
		historico := api.PontuacaoHistorico{}
		historico.Get(d.AtletaId)
		CacheHistoricoAtleta[d.AtletaId] = historico
	}
	SleepCacheSecond(7200)
	UpdateHistoricoPontuacao()
}

func UpdateCache() {
	go UpdateStatus()
	go UpdateDestaques()
	go UpdateMercado()
	go UpdatePontuados()
	go UpdatePartidas()
	go UpdateHistoricoPontuacao()
}

func SleepCacheSecond(t int) {
	time.Sleep(time.Duration(t) * time.Second)
}
