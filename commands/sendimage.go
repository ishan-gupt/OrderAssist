package commands

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"BeBot/utils"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

type Image struct {
	Prompt string `json:"prompt"`
}
type Response struct {
	Image string `json:"image"`
}

func ImageHandler(evt interface{}, c *whatsmeow.Client) {
	switch v := evt.(type) {
	case *events.Message:
		// Ignore messages sent by myself
		if v.Info.Sender.User == c.Store.ID.User {
			return
		}
		msg := strings.ToLower(v.Message.GetConversation())
		match1, _ := regexp.MatchString("Image of", msg)
		if match1 == true {
			err := SendImage(c, v.Info.Chat, msg)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

func SendImage(client *whatsmeow.Client, receiver types.JID, message string) error {
	c := Response{}
	myJson := Image{message}
	reqData, err := json.Marshal(myJson)
	if err != nil {
		return fmt.Errorf("failed to marshal image request: %v", err)
	}
	
	err = utils.PostJson("https://openai-image-7la8.onrender.com/image/generate", &c, reqData)
	if err != nil {
		return fmt.Errorf("failed to call image generation API: %v", err)
	}
	
	err = utils.UploadImage(
		c.Image,
		client,
		receiver)
	if err != nil {
		return fmt.Errorf("failed to upload image: %v", err)
	}

	return nil
}
