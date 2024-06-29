package utils

import (
	ctx "context"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

func SendMessage(message string, client *whatsmeow.Client, receiver types.JID) error {
	_, err := client.SendMessage(ctx.Background(), receiver, &waE2E.Message{
		Conversation: proto.String(message),
	})

	return err
}
