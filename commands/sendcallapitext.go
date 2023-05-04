package commands

import (
	"BeBot/utils"
	"fmt"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"regexp"
	"strings"
)

type ShoudDeploy struct {
	Timezone      string `json:"timezone"`
	ShouldIDeploy bool   `json:"shouldideploy"`
	Message       string `json:"message"`
}
type ShouldJoke struct {
	error    bool   `json:"error"`
	category string `json:"category"`
	Type     string `json:"type"`
	Joke     string `json:"joke"`
}

func ApiHandler(evt interface{}, c *whatsmeow.Client) {
	switch v := evt.(type) {
	case *events.Message:
		msg := strings.ToLower(v.Message.GetConversation())
		match1, _ := regexp.MatchString("deploy", msg)
		if match1 == true {
			err := ShoudIDeployToday(c, v.Info.Chat)
			fmt.Println(v.Info.Sender)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		match2, _ := regexp.MatchString("joke", msg)
		if match2 == true {
			err := ShoudIPJokeToday(c, v.Info.Chat)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

func ShoudIDeployToday(client *whatsmeow.Client, receiver types.JID) error {
	sd := new(ShoudDeploy)
	utils.GetJson("https://shouldideploy.today/api?tz=America/Sao_Paulo", sd)
	err := utils.SendMessage(
		sd.Message,
		client,
		receiver)

	return err
}

func ShoudIPJokeToday(client *whatsmeow.Client, receiver types.JID) error {
	sd := new(ShouldJoke)
	utils.GetJson("https://v2.jokeapi.dev/joke/Dark?blacklistFlags=nsfw&type=single", sd)
	err := utils.SendMessage(
		sd.Joke,
		client,
		receiver)

	return err
}
