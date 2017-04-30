package kartola

import "github.com/pressly/chi"

func BuildRoutes(mux *chi.Mux){
	mux.Get("/", GetHome)
	mux.Get("/status", GetStatus)
	mux.Get("/times/:q", GetTimes)
	mux.Get("/time/id/:id", GetTime)
	mux.Get("/atletas/mercado", GetMercado)
}