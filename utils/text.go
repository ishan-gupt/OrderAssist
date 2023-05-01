package utils

import (
	ctx "context"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

func SendMessage(message string, client *whatsmeow.Client, receiver types.JID) error {
	_, err := client.SendMessage(ctx.Background(), receiver, &waProto.Message{
		Conversation: proto.String(message),
	})

	return err
}
