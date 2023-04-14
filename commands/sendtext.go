package commands

import (
	"BeBot/utils"
	"fmt"
	"strings"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func TextHandler(evt interface{}, c *whatsmeow.Client) {
	switch v := evt.(type) {
	case *events.Message:
		msg := strings.ToLower(v.Message.GetConversation())
		text := "Hi there. Welcome to BeBot, your own whatsapp chatbot. \nIt will sort out all your regular needs. Handy for both students and professionals. \nClient's(company) customers query according to the product. \nClient(company) can manage their customer queries effectively, automating customer support."
		text1 := "To get started send the following words for auto messages \n1.Astrology\n 2.Jokes\n 3.deploy"
		switch msg {
		case "hello":
			err := utils.SendMessage(text, c, v.Info.Chat)
			if err != nil {
				fmt.Println(err.Error())
			}
			err1 := utils.SendMessage(text1, c, v.Info.Chat)
			if err1 != nil {
				fmt.Println(err1.Error())
			}

		case "hello there":
			err := utils.SendMessage(text, c, v.Info.Chat)
			if err != nil {
				fmt.Println(err.Error())
			}
			err1 := utils.SendMessage(text1, c, v.Info.Chat)
			if err1 != nil {
				fmt.Println(err1.Error())
			}

		case "hi":
			err := utils.SendMessage(text, c, v.Info.Chat)
			if err != nil {
				fmt.Println(err.Error())
			}
			err1 := utils.SendMessage(text1, c, v.Info.Chat)
			if err1 != nil {
				fmt.Println(err1.Error())
			}

		case "hey":
			err := utils.SendMessage(text, c, v.Info.Chat)
			if err != nil {
				fmt.Println(err.Error())
			}
			err1 := utils.SendMessage(text1, c, v.Info.Chat)
			if err1 != nil {
				fmt.Println(err1.Error())
			}

		}
	}
}
