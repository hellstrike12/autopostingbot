package updates

import (
	log "github.com/sirupsen/logrus"
	"github.com/zelenin/go-tdlib/client"
	"gitlab.com/shitposting/autoposting-bot/commands"
	"gitlab.com/shitposting/autoposting-bot/repository"
)

var(
	handlers = map[string]commands.Handler {
		"status" : commands.StatusCommandHandler{},
	}
)

func HandleUpdates(listener *client.Listener) {

	for update := range listener.Updates {

		if update.GetClass() == client.ClassUpdate {
			switch update.GetType() {
			case client.TypeUpdateNewMessage:
				handleNewMessage(update.(*client.UpdateNewMessage).Message)
			}

			log.Printf("Type: %s, Value: %#v", update.GetType(), update)

		}
	}

}

func handleNewMessage(message *client.Message) {

	if message.SenderUserId == repository.Me.Id {
		log.Print("MESSAGE FROM SELF")
		return
	}


	switch message.Content.MessageContentType() {
	case client.TypeMessageText:
		handleText(message)
	case client.TypeMessageAnimation:

	case client.TypeMessagePhoto:

	case client.TypeMessageVideo:
		
		
	}

}
