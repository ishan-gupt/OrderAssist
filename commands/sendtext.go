package commands

import (
	"BeBot/utils"
	"fmt"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
	"go.mau.fi/whatsmeow/types"
	"strings"
)

type Search struct {
	Contact string `json:"contact"`
}

// sendWelcomeMessage sends the welcome messages to the user
func sendWelcomeMessage(c *whatsmeow.Client, chat types.JID) error {
	text := "Hi there. Welcome to BeBot, your own whatsapp chatbot. \n As you are using us for the first time plese help us with your details \n"
	text1 := "To get started select the store you want to order from \n 1.Abc stores \n 2.Bkc stores \n 3. Mlp stores \n Send us the below asked deatils. Please replace *\":\"* in the below messages with *\"-\"*\nThank You"
	text2 := "Name:<Your name>\nAddress:<Your address>\nStoreid:<1 or 2 or 3>"

	messages := []string{text, text1, text2}
	
	for _, msg := range messages {
		err := utils.SendMessage(msg, c, chat)
		if err != nil {
			return fmt.Errorf("failed to send welcome message: %v", err)
		}
	}
	
	return nil
}

func TextHandler(evt interface{}, c *whatsmeow.Client) {
	switch v := evt.(type) {
	case *events.Message:
		// Ignore messages sent by myself
		if v.Info.Sender.User == c.Store.ID.User {
			return
		}
		msg := strings.ToLower(v.Message.GetConversation())
		print(msg)

		switch msg {
		case "hello", "hello there", "hi", "hey":
			err := sendWelcomeMessage(c, v.Info.Chat)
			if err != nil {
				fmt.Printf("Error sending welcome message: %v\n", err)
			}
		}
	}

}

// func DoesItExisit(num string) int {
// 	url := "https://oms-bebot-backend.onrender.com/api/user/exist"
// 	user := Search{Contact: num}

// 	jsonPayload, err := json.Marshal(user)
// 	if err != nil {
// 		panic(err)
// 	}

// 	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonPayload))
// 	if err != nil {
// 		panic(err)
// 	}

// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	var result map[string]interface{}
// 	err = json.NewDecoder(resp.Body).Decode(&result)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(result)
// 	return resp.StatusCode

// }
