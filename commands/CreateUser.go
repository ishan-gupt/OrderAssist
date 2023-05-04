package commands

import (
	"BeBot/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
	"net/http"
	"regexp"
	"strings"
)

type User struct {
	Number  string `json:"contact"`
	Address string `json:"address"`
	Name    string `json:"name"`
	StoreId string `json:"storeid"`
}

var name string = "null"
var address string = "null"
var storeid string = "null"

func UserHandler(evt interface{}, c *whatsmeow.Client) {
	var code int
	switch v := evt.(type) {
	case *events.Message:
		var number = (v.Info.Sender).String()
		num := number[2:12]
		msg := strings.ToLower(v.Message.GetConversation())
		match1, _ := regexp.MatchString("name-", msg)
		if match1 == true {
			if address == "null" {
				err := utils.SendMessage("Where Do you live?", c, v.Info.Chat)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
			name = msg[5:]
			fmt.Println(name)
		}
		match2, _ := regexp.MatchString("address-", msg)
		if match2 == true {
			if storeid == "null" {
				err := utils.SendMessage("Where would you like to order from ", c, v.Info.Chat)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
			address = msg[8:]
			fmt.Println(address)
		}
		match3, _ := regexp.MatchString("storeid-", msg)
		if match3 == true {
			if name == "null" {
				err := utils.SendMessage("What should we call you?", c, v.Info.Chat)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
			if address == "null" {
				err := utils.SendMessage("Where Do you live?", c, v.Info.Chat)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
			storeid = msg[8:]
			fmt.Println(storeid)
		}
		if name != "null" && storeid != "null" && address != "null" {
			code = CreateUser(name, num, address, storeid)
		}
		if code == 201 {
			err := utils.SendMessage("Profile Created", c, v.Info.Chat)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

	}
}

func CreateUser(name string, num string, address string, storeid string) int {
	url := "https://oms-bebot-backend.onrender.com/api/user/signup"
	user := User{Number: num, Name: name, Address: address, StoreId: storeid}

	jsonPayload, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	return resp.StatusCode
}
