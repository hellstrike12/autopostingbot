package telegram

import (
	"github.com/zelenin/go-tdlib/client"
	"strings"
	"unicode/utf16"
)

// The found variable allows us to be sure entities won't be nil
//goland:noinspection GoNilness
func GetCommand(text []uint16, entities []*client.TextEntity) (command, arguments string, found bool) {

	// A bot command for us needs to be in the beginning of the message
	found = entities != nil && len(entities) > 0 &&
		entities[0].Type.TextEntityTypeType() == client.TypeTextEntityTypeBotCommand

	if !found {
		return
	}

	// The start will always be 0 for the "beginning of the message" hypothesis.
	// Start from 1 to avoid having the / in the command.
	commandUtf16 := text[1:entities[0].Length]
	command = strings.ToLower(string(utf16.Decode(commandUtf16)))

	// Commands may not have arguments
	if len(text) > int(entities[0].Length) {
		argumentsUtf16 := text[entities[0].Length:]
		arguments = strings.TrimSpace(string(utf16.Decode(argumentsUtf16)))
	}

	return

}
