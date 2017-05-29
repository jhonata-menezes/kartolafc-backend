package notification

import (
	"log"
	webpush "github.com/SherClockHolmes/webpush-go"
	"encoding/json"
)

const (
	vapidPrivateKey = "W1HAZf9KEso6B9PgCM0xqF_d4KBFe88qKGu-KtAkuvA"
)

var usersKartolafc []webpush.Subscription

var ChannelSubscribe = make(chan webpush.Subscription, 1000)

// adiciona os inscritos no array [temporario]
func addSubscribe(subscription *chan webpush.Subscription) {
	usersKartolafc = make([]webpush.Subscription, 10000)
	for s := range *subscription {
		usersKartolafc = append(usersKartolafc, s)
	}
}


func notify(channelNotifcation *chan MessageNotification,) {

	for m := range *channelNotifcation {
		// Send Notification
		mByte, err := json.Marshal(m)
		if err != nil {
			continue
		}

		for _, s := range usersKartolafc {
			_, err := webpush.SendNotification(mByte, &s, &webpush.Options{
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
