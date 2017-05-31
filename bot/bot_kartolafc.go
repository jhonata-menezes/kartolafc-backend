package bot

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"github.com/jhonata-menezes/kartolafc-backend/cmd"
	"regexp"
	"github.com/jhonata-menezes/kartolafc-backend/notification"
	"errors"
)

func Run(channelNotification chan *notification.MessageNotification) {
	// se nao tiver configurado nenhuma chave, bot nao e iniciado.
	if cmd.Config.BotKey == "" {
		log.Println("bot desabilitado")
		return
	}
	bot, err := tgbotapi.NewBotAPI(cmd.Config.BotKey)
	if err != nil {
		log.Panic(err)
	}
	//bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.From.ID != cmd.Config.BotIdClient {
			continue
		}
		log.Printf("[ %s %s ] %s \n", update.Message.From.UserName, update.Message.From.ID, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "comando não encontrado")
		msg.ReplyToMessageID = update.Message.MessageID
		command := TelegramGetCommand(update.Message.Text)

		switch command {
		case "/n":
			notify, err := getNotificacao(update.Message.Text)
			if err == nil {
				channelNotification <- &notify
				msg.Text = "Notificaçao enviada"
			} else {
				msg.Text = "padrao incorreto"
			}
			break
		}

		bot.Send(msg)
	}
}


func TelegramGetCommand(s string) string {
	regComando, _ := regexp.Compile("(/[\\w-_]+) ?")
	if regComando.MatchString(s) {
		match := regComando.FindStringSubmatch(s)
		return match[1]
	}
	return ""
}

func getNotificacao(s string) (notification.MessageNotification, error) {
	regComando, _ := regexp.Compile("(/[\\w-_]+) ?\n(.*?)\n(.*?)\n(.*?)\n(.*?)$")
	if regComando.MatchString(s) {
		match := regComando.FindStringSubmatch(s)
		return notification.MessageNotification{match[2], match[3], match[4], "", nil, match[5]}, nil
	}
	log.Println("errou")
	return notification.MessageNotification{}, errors.New("nao foi possivel parsear")
}

func TelegramCommandGetMessage(s string) string {
	regMensagem, _ := regexp.Compile("(/[\\w-_]+) (.*?)$")
	if regMensagem.MatchString(s) {
		match := regMensagem.FindStringSubmatch(s)
		return match[2]
	}
	return ""
}