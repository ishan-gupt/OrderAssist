package commands

import (
	"BeBot/utils"
	"fmt"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
	"strings"
)

type Search struct {
	Contact string `json:"contact"`
}

func TextHandler(evt interface{}, c *whatsmeow.Client) {
	switch v := evt.(type) {
	case *events.Message:
		msg := strings.ToLower(v.Message.GetConversation())
		print(msg)
		var text string
		var text1 string
		var text2 string

		text = "Hi there. Welcome to BeBot, your own whatsapp chatbot. \n As you are using us for the first time plese help us with your details \n"
		text1 = "To get started select the store you want to order from \n 1.Abc stores \n 2.Bkc stores \n 3. Mlp stores \n Send us the below asked deatils. Please replace *\":\"* in the below messages with *\"-\"*\nThank You"
		text2 = "Name:<Your name>\nAddress:<Your address>\nStoreid:<1 or 2 or 3>"

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
			err3 := utils.SendMessage(text2, c, v.Info.Chat)
			if err3 != nil {
				fmt.Println(err3.Error())
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
			err3 := utils.SendMessage(text2, c, v.Info.Chat)
			if err3 != nil {
				fmt.Println(err3.Error())
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
			err3 := utils.SendMessage(text2, c, v.Info.Chat)
			if err3 != nil {
				fmt.Println(err3.Error())
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
			err3 := utils.SendMessage(text2, c, v.Info.Chat)
			if err3 != nil {
				fmt.Println(err3.Error())
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
