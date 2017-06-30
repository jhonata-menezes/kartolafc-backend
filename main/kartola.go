package main

import (
	"github.com/pressly/chi"
	kartolafc "github.com/jhonata-menezes/kartolafc-backend"
	"net/http"
	"github.com/pressly/chi/middleware"
	"github.com/pressly/chi/render"
	"github.com/jhonata-menezes/kartolafc-backend/cmd"
	"log"
	"github.com/jhonata-menezes/kartolafc-backend/api"
	"github.com/jhonata-menezes/kartolafc-backend/notification"
	"github.com/goware/cors"
	"github.com/jhonata-menezes/kartolafc-backend/bot"
	"gopkg.in/mgo.v2"
	"time"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

func main() {

	go kartolafc.UpdateCache()

	// mongodb
	session, err := mgo.Dial(cmd.Config.MongoDB)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	session.SetSocketTimeout(3 * time.Hour)
	ranking(session.Copy())

	for i:=0; i<2000; i++ {
		kartolafc.ChannelCollectionTime <- session.Copy().DB("kartolafc").C("times")
	}
	//-------------------------------------------------------------------------------------

	// web push
	go notification.AddSubscribe(&notification.ChannelSubscribe)
	channelMessageNotification := make(chan *notification.MessageNotification, 1000)
	go bot.Run(channelMessageNotification)
	go notification.Notify(channelMessageNotification)

	// pontuados
	go getPontuados(session.Copy())
	go kartolafc.Run(channelMessageNotification)

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
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	router.Use(myCors.Handler)

	kartolafc.BuildRoutes(router)
	log.Println("Bora Cumpade.")
	log.Println("listen", cmd.Config.ServerBind)
	err = http.ListenAndServe(cmd.Config.ServerBind, router)
	if err != nil {
		panic(err)
	}
}

func ranking(session *mgo.Session) {
	colllection := session.DB("kartolafc").C("times")
	go kartolafc.LoadInMemory(colllection)
}

func getPontuados(session *mgo.Session) {
	// aguarda alguns segundos para o status estar disponivel
	time.Sleep(3 * time.Second)
	// se o mercado estiver fechado nao e preciso
	if kartolafc.CacheStatus.StatusMercado == 2 {
		return
	}
	// temporario, para criar pontuados mesmo com mercado aberto
	collection := session.DB("kartolafc").C("pontuados")

	var pontuadosResult  interface{}
	err := collection.Find(bson.M{"rodada": (kartolafc.CacheStatus.RodadaAtual - 1)}).One(&pontuadosResult)
	if err != nil {
		log.Println("erro na consulta dos pontuados", err)
		return
	}

	pontuadosByte, err := json.Marshal(pontuadosResult)
	if err != nil {
		log.Println("nao foi possivel mapear os pontuados na interface")
	}

	pontuados := api.Pontuados{}
	json.Unmarshal(pontuadosByte, &pontuados)

	if pontuados.Rodada == 0 {
		log.Println("quantidade de pontuados retornado da collection igual a zero")
		return
	}
	kartolafc.CachePontuados = pontuados
}
