package api

import "github.com/zelenin/go-tdlib/client"

// SendAnimation shares an animation to a certain chat.
// If replyToMessageID is not 0, the animation will be in reply to that message id.
// caption and entities can be used to attach a message with markdown.
func SendAnimation(chatID, replyToMessageID int64, remoteFileID, caption string, entities []*client.TextEntity) (*client.Message, error) {

	request := client.SendMessageRequest{
		ChatId:           chatID,
		ReplyToMessageId: replyToMessageID,
		InputMessageContent: &client.InputMessageAnimation{
			Animation: &client.InputFileRemote{
				Id: remoteFileID,
			},
			Caption: &client.FormattedText{
				Text:     caption,
				Entities: entities,
			},
		},
	}

	return tdlibClient.SendMessage(&request)

}

// GetAnimationFileInfoFromMessage returns the Animation structure
// of a given client.Message.
func GetAnimationFileInfoFromMessage(message *client.Message) *client.File {
	return message.Content.(*client.MessageAnimation).Animation.Animation
}
