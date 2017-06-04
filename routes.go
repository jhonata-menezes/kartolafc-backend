package kartolafc

import (
	"github.com/pressly/chi"
)

func BuildRoutes(mux *chi.Mux){
	mux.Get("/", GetHome)
	mux.Get("/mercado/status", GetStatus)
	mux.Get("/times/:q", GetTimes)
	mux.Get("/time/id/:id", GetTime)
	mux.Get("/atletas/mercado", GetMercado)
	mux.Get("/mercado/destaques", GetDestaques)
	mux.Get("/ligas/:q", GetLigas)
	mux.Get("/liga/:id/:page", GetLiga)
	mux.Get("/atletas/pontuados", GetPontuados)
	mux.Get("/partidas/:partida", GetPartida)
	mux.Post("/notificacao/adicionar", AddNotificacao)
	mux.Get("/ranking/melhores", GetMelhoresRanking)
	mux.Get("/ranking/time/id/:id", GetRankingTimeId)
}