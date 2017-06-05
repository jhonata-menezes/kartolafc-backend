package notification

import (
	"log"
	webpush "github.com/SherClockHolmes/webpush-go"
	"encoding/json"
	"github.com/jhonata-menezes/kartolafc-backend/cmd"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
)

// Keys are the base64 encoded values from PushSubscription.getKey()
type Keys struct {
	Auth   string `json:"auth"`
	P256dh string `json:"p256dh"`
}

// Subscription represents a PushSubscription object from the Push API
type Subscription struct {
	Endpoint string `json:"endpoint"`
	Keys     Keys   `json:"keys"`
	TimeId 	 int    `json:"time_id"`
}

type requestBroadcast struct {
	M []byte
	S Subscription
}

var vapidPrivateKey = cmd.Config.VapidPrivate

var usersKartolafc = make(map[string]Subscription, 10000)

var ChannelSubscribe = make(chan Subscription, 1000)

var session *mgo.Session

func init() {
	s, err := mgo.Dial(cmd.Config.MongoDB)
	if err != nil {
		panic(err)
	}
	session = s
	collection := getCollection()
	defer collection.Database.Session.Close()

	usuariosCadastrados := []Subscription{}
	collection.Find(bson.M{}).All(&usuariosCadastrados)

	for _, u := range usuariosCadastrados {
		usersKartolafc[u.Endpoint] = u
	}
}

// adiciona os inscritos no array [temporario]
func AddSubscribe(subscription *chan Subscription) {
	collection := getCollection()
	for s := range *subscription {
		if _, v := usersKartolafc[s.Endpoint]; v {
			continue
		}
		usersKartolafc[s.Endpoint] = s
		// gravar na collection
		qtd, err := collection.Find(bson.M{"endpoint": s.Endpoint}).Count()
		if err != nil {
			panic(err)
		}
		if qtd == 0 {
			err = collection.Insert(s)
			if err != nil {
				log.Println("nao foi possivel inserir novo endpoint para push notification")
			} else {
				log.Println("novo usuario cadastrado para notificaçoes")
			}
		}
	}
}


func Notify(channelNotifcation chan *MessageNotification) {
	channelRequest := make(chan requestBroadcast, 100000)
	for i:=0; i<cmd.Config.JobsNotification; i++ {
		go requestEndpoint(channelRequest)
	}

	for m := range channelNotifcation {
		mByte, err := json.Marshal(m)
		if err != nil {
			continue
		}

		// channel de usuarios, as request sao enviadas paralelamente
		for _, b := range usersKartolafc {
			channelRequest <- requestBroadcast{mByte, b}
		}
	}
}

func requestEndpoint(ch chan requestBroadcast) {
	// Send Notification
	for s := range ch {
		adapter := webpush.Subscription{}
		adapter.Endpoint = s.S.Endpoint
		adapter.Keys.Auth = s.S.Keys.Auth
		adapter.Keys.P256dh = s.S.Keys.P256dh
		res, err := webpush.SendNotification(s.M, &adapter, &webpush.Options{
			Subscriber:      "mailto:jhonatamenezes10@gmail.com",
			TTL:             60,
			VAPIDPrivateKey: vapidPrivateKey,
		})

		if err != nil {
			log.Println(err)
		} else {
			s, _ := ioutil.ReadAll(res.Body)
			log.Printf("vapid: %#v", string(s))
			res.Body.Close()
		}
	}
}

func getCollection() *mgo.Collection {
	return session.Copy().DB("kartolafc").C("webpush")
}
