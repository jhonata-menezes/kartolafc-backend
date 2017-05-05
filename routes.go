package kartolafc

import "github.com/pressly/chi"

func BuildRoutes(mux *chi.Mux){
	mux.Get("/", GetHome)
	mux.Get("/mercado/status", GetStatus)
	mux.Get("/times/:q", GetTimes)
	mux.Get("/time/id/:id", GetTime)
	mux.Get("/atletas/mercado", GetMercado)
	mux.Get("/mercado/destaques", GetDestaques)
	mux.Get("/ligas/:q", GetLigas)
	mux.Get("/liga/:id", GetLiga)
}