package kartolafc

import (
	"github.com/jhonata-menezes/kartolafc-backend/notification"
	"time"
	"github.com/jhonata-menezes/kartolafc-backend/api"
	"log"
	"fmt"
)

var scoutsPreparados = map[string]string{"G":"Gol", "CV": "Cartão Vermelho",
	"CA": "Cartão Amarelo", "GC": "Gol Contra", "DP": "Defesa de Penalti"}

var scoutsPontuacao = map[string]int{"G":8, "CV": -5, "CA": -2, "GC": -6, "DP": 7}

// notificações duplicadas sendos enviadas, mapear as notificações que foram enviadas para nao duplicar
var atletasNotificados = make(map[int]map[string]int, 200000)

func Run(channelMessage chan *notification.MessageNotification) {
	time.Sleep(5 * time.Second)
	CacheLocal := api.Pontuados{}
	CopyStructs(CachePontuados, &CacheLocal)

	// na inicialiazacao verifica os scouts que tem para nao passar scouts passado
	for atletaId, p := range CachePontuados.Atletas {
		for scout, v := range p.Scout {
			atletasNotificados[atletaId] = map[string]int{scout: v}
		}
	}
	for {
		attCache := false
		for atletaId, p := range CachePontuados.Atletas {
			for scout, v := range p.Scout {
				// apenas os scouts permitidos serão validos
				if !(len(scoutsPreparados[scout]) > 0) {
					continue
				}
				// se o valor do scout atual for maior do que esta em cache. entao houve atualizacao
				if CacheLocal.Atletas[atletaId].Scout[scout] >= 0 && CacheLocal.Atletas[atletaId].Scout[scout] < v {
					// scout existe
					log.Println("notificando por scout att", p.Apelido, scout, v, CacheLocal.Atletas[atletaId].Scout[scout])
					dispatchScout(scout, CachePontuados.Atletas[atletaId], channelMessage, atletaId, v)
					attCache = true
				}
			}
		}

		if attCache {
			CacheLocal = api.Pontuados{}
			CopyStructs(CachePontuados, &CacheLocal)
			attCache = false
		}

		for CacheStatus.StatusMercado != 2 {
			time.Sleep(60 * time.Minute)
		}
		time.Sleep(10 * time.Second)
	}
}

func dispatchScout(scout string, atleta api.Pontuado, ch chan *notification.MessageNotification, atletaId, valorScout int) {
	if len(scoutsPreparados[scout]) > 0 && atletasNotificados[atletaId][scout] < valorScout {
		log.Printf("notificando scout %#v %#v", scout, atleta)
		// existe scout na lista preparada
		m := notification.MessageNotification{}
		m.Title = atleta.Apelido + " #" + scoutsPreparados[scout]
		m.Link = "https://kartolafc.com.br"
		m.Icon = atleta.Foto
		m.Body = fmt.Sprintf("Pontuaçao %.2f", CachePontuados.Atletas[atletaId].Pontuacao)
		m.Scout = scout
		atletasNotificados[atletaId] = map[string]int{scout: valorScout}
		ch <- &m
	}
}
