package commands

import (
	"fmt"
	"strings"

	"BeBot/utils"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

type Image struct {
	Copyright    string `json:"copyright"`
	Date         string `json:"date"`
	Explaination string `json:"explaination"`
	URL          string `json:"url"`
	HDURL        string `json:"hdurl"`
	Title        string `json:"title"`
}

func ImageHandler(evt interface{}, c *whatsmeow.Client) {
	switch v := evt.(type) {
	case *events.Message:
		msg := strings.ToLower(v.Message.GetConversation())

		switch msg {
		case "astrology":
			err := SendImage(c, v.Info.Chat)
			if err != nil {
				fmt.Println(err.Error())
			}

		case "astro":
			err := SendImage(c, v.Info.Chat)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

func SendImage(client *whatsmeow.Client, receiver types.JID) error {
	c := Image{}
	utils.GetJson("https://go-apod.herokuapp.com/apod", &c)
	err := utils.UploadImage(
		c.URL,
		c.Title,
		"image/png",
		client,
		receiver)

	return err
}
