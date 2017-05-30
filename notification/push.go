package notification

import (
	"log"
	webpush "github.com/SherClockHolmes/webpush-go"
	"encoding/json"
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
}

const (
	vapidPrivateKey = "W1HAZf9KEso6B9PgCM0xqF_d4KBFe88qKGu-KtAkuvA"
)

var usersKartolafc map[string]Subscription

var ChannelSubscribe = make(chan Subscription, 1000)

// adiciona os inscritos no array [temporario]
func AddSubscribe(subscription *chan Subscription) {
	usersKartolafc = make(map[string]Subscription, 10000)
	for s := range *subscription {
		if _, v := usersKartolafc[s.Endpoint]; v {
			continue
		}
		usersKartolafc[s.Endpoint] = s
	}
}


func Notify(channelNotifcation *chan MessageNotification) {

	for m := range *channelNotifcation {
		// Send Notification
		mByte, err := json.Marshal(m)
		if err != nil {
			continue
		}

		for _, s := range usersKartolafc {
			adapter := webpush.Subscription{}
			adapter.Endpoint = s.Endpoint
			adapter.Keys.Auth = s.Keys.Auth
			adapter.Keys.P256dh = s.Keys.P256dh
			_, err := webpush.SendNotification(mByte, &adapter, &webpush.Options{
				Subscriber:      "mailto:jhonatamenezes10@gmail.com",
				TTL:             60,
				VAPIDPrivateKey: vapidPrivateKey,
			})

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
