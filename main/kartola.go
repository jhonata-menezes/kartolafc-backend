package main

import (
	"github.com/pressly/chi"
	kartolafc "github.com/jhonata-menezes/kartolafc-backend"
	"net/http"
	"github.com/pressly/chi/middleware"
	"github.com/pressly/chi/render"
	"github.com/jhonata-menezes/kartolafc-backend/cmd"
	"log"
	"io/ioutil"
	"encoding/json"
	"github.com/jhonata-menezes/kartolafc-backend/api"
	"github.com/jhonata-menezes/kartolafc-backend/notification"
	"time"
	"github.com/goware/cors"
)

func main() {

	go kartolafc.UpdateCache()

	// temporario, para criar pontuados mesmo com mercado aberto
	fileByte, err := ioutil.ReadFile("./pontuados.json")
	if err == nil {
		p := api.Pontuados{}
		json.Unmarshal(fileByte, &p)
		kartolafc.CachePontuados = p
	}
	//-------------------------------------------------------------------------------------

	// temporario para teste com web push
	go notification.AddSubscribe(&notification.ChannelSubscribe)
	channelMessageNotification := make(chan notification.MessageNotification, 1000)
	go notification.Notify(&channelMessageNotification)
	go func (){
		for {
			m := notification.MessageNotification{}
			m.Body = "teste 001"
			m.Badge = "http://images.terra.com/2015/05/20/corinthians.png"
			m.Icon = "http://images.terra.com/2015/05/20/corinthians.png"
			m.Link = "http://kartolafc.com.br/#/ligas"
			m.Title = "Title Teste"
			m.Vibrate = []int{200, 100, 200}
			channelMessageNotification <- m

			log.Println("Enviando notificacao")
			time.Sleep(30 * time.Second)
		}
	}()
	// ----------------------------------

	router := chi.NewRouter()
	router.Use(middleware.DefaultCompress)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)

	myCors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	router.Use(myCors.Handler)

	kartolafc.BuildRoutes(router)
	log.Println("Bora Cumpade.")
	log.Println("listen", cmd.ServerBind)
	err = http.ListenAndServe(cmd.ServerBind, router)
	if err != nil {
		panic(err)
	}
}
