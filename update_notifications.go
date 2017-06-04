package kartolafc

import (
	"github.com/jhonata-menezes/kartolafc-backend/notification"
	"time"
	"github.com/jhonata-menezes/kartolafc-backend/api"
	"log"
)

var scoutsPreparados = map[string]string{"G":"Gol", "CV": "Cartão Vermelho",
	"CA": "Cartão Amarelo", "GC": "Gol Contra", "DP": "Defesa de Penalti"}

// notificações duplicadas sendos enviadas, mapear as notificações que foram enviadas para nao duplicar
var atletasNotificados = make(map[int]map[string]int, 1000)

func Run(channelMessage chan *notification.MessageNotification) {
	time.Sleep(5 * time.Second)
	CacheLocal := CachePontuados
	for {
		Pontuados := CachePontuados
		for atletaId, p := range Pontuados.Atletas {
			for scout, v := range p.Scout {
				// apenas os scouts permitidos serão validos
				if !(len(scoutsPreparados[scout]) > 0) {
					continue
				}

				if CacheLocal.Atletas[atletaId].Scout[scout] >= 0 {
					// scout existe
					if CacheLocal.Atletas[atletaId].Scout[scout] != v {
						log.Println("notificando por scout att", p.Apelido, scout, v)
						dispatchScout(scout, Pontuados.Atletas[atletaId], channelMessage)
						CacheLocal = Pontuados
					}
				} else {
					log.Println("notificando por scout novo", p.Apelido, scout, v)
					dispatchScout(scout, Pontuados.Atletas[atletaId], channelMessage)
					CacheLocal = Pontuados
				}
			}
		}

		for CacheStatus.StatusMercado != 2 {
			time.Sleep(60 * time.Minute)
		}
		time.Sleep(10 * time.Second)
	}
}

func dispatchScout(scout string, atleta api.Pontuado, ch chan *notification.MessageNotification) {
	if len(scoutsPreparados[scout]) > 0 {
		log.Printf("notificando scout %#v %#v", scout, atleta)
		// existe scout na lista preparada
		m := notification.MessageNotification{}
		m.Title = scoutsPreparados[scout]
		m.Link = "https://kartolafc.com.br"
		m.Icon = atleta.Foto
		m.Body = atleta.Apelido
		ch <- &m
	}
}
