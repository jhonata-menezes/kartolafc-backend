package notification

type MessageNotification struct {
	Title string `json:"title"`
	Body string `json:"body"`
	Icon string `json:"icon"`
	Badge string `json:"badge"`
	Vibrate string `json:"vibrate"`
	Link string `json:"link"`
}

