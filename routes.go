package kartolafc

import (
	"github.com/go-chi/chi"
)

func BuildRoutes(mux *chi.Mux){
	mux.Get("/", GetHome)
	mux.Get("/mercado/status", GetStatus)
	mux.Get("/times/{q}", GetTimes)
	mux.Get("/time/id/{id}", GetTime)
	mux.Get("/time/id/{id}/{rodada}", GetTimeHistorico)
	mux.Get("/atletas/mercado", GetMercado)
	mux.Get("/mercado/destaques", GetDestaques)
	mux.Get("/ligas/{q}", GetLigas)
	mux.Get("/liga/{id}/{page}", GetLiga)
	mux.Get("/atletas/pontuados", GetPontuados)
	mux.Get("/partidas/{partida}", GetPartida)
	mux.Post("/notificacao/adicionar", AddNotificacao)
	mux.Get("/ranking/melhores", GetMelhoresRanking)
	//mux.Get("/ranking/melhores/pro", GetMelhoresRankingPro)
	//mux.Get("/ranking/time/id/:id", GetRankingTimeId)
	mux.Get("/atletas/historico/{id}", GetPontuacaHistorico)
	mux.Post("/login/cartolafc", PostLogin)
	mux.Get("/time/info", GetMeuTime)
	mux.Post("/time/salvar", PostSalvarTime)
}