package commands

import (
	"BeBot/utils"
	"fmt"
	"regexp"
	"strings"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

// Existing structs for API responses
type ShoudDeploy struct {
	Timezone      string `json:"timezone"`
	ShouldIDeploy bool   `json:"shouldideploy"`
	Message       string `json:"message"`
}
type ShouldJoke struct {
	Error    bool   `json:"error"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Joke     string `json:"joke"`
}

// The ApiHandler function to handle WhatsApp message events
func ApiHandler(evt interface{}, c *whatsmeow.Client) {
	switch v := evt.(type) {
	case *events.Message:
		// Ignore messages sent by myself
		if v.Info.Sender.User == c.Store.ID.User {
			return
		}
		msg := strings.ToLower(v.Message.GetConversation())

		// Match for "deploy" keyword
		match1, _ := regexp.MatchString("deploy", msg)
		if match1 {
			err := ShoudIDeployToday(c, v.Info.Chat)
			fmt.Println(v.Info.Sender)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		// Match for "joke" keyword
		match2, _ := regexp.MatchString("joke", msg)
		if match2 {
			err := ShoudIPJokeToday(c, v.Info.Chat)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		// Match for any chat command to call external API
		match3, _ := regexp.MatchString("plate", msg) // Change "chatgpt" to your desired keyword
		if match3 {
			err := CallChatAPI(c, v.Info.Chat, msg) // Call the API with the message content
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

// Function to call the ShouldIDeploy API
func ShoudIDeployToday(client *whatsmeow.Client, receiver types.JID) error {
	sd := new(ShoudDeploy)
	err := utils.GetJson("https://shouldideploy.today/api?tz=America/Sao_Paulo", sd)
	if err != nil {
		return fmt.Errorf("failed to call shouldideploy API: %v", err)
	}
	
	err = utils.SendMessage(
		sd.Message,
		client,
		receiver)
	return err
}

// Function to call the Joke API
func ShoudIPJokeToday(client *whatsmeow.Client, receiver types.JID) error {
	sd := new(ShouldJoke)
	err := utils.GetJson("https://v2.jokeapi.dev/joke/Miscellaneous,Dark,Pun,Spooky,Christmas?type=single", sd)
	if err != nil {
		return fmt.Errorf("failed to call joke API: %v", err)
	}
	
	err = utils.SendMessage(
		sd.Joke,
		client,
		receiver)
	return err
}

// New function to call the external API using callAPI
func CallChatAPI(client *whatsmeow.Client, receiver types.JID, msg string) error {
	// Call the external chat API with the provided message content
	responseContent, err := callAPI(msg)
	if err != nil {
		return err
	}

	// Send the response content back to WhatsApp chat
	err = utils.SendMessage(
		responseContent,
		client,
		receiver)
	return err
}
