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

type Order struct {
	Number   string `json:"contact"`
	OrderDel string `json:"order_details"`
}

func OrderHandler(evt interface{}, c *whatsmeow.Client) {
	switch v := evt.(type) {
	case *events.Message:
		msg := strings.ToLower(v.Message.GetConversation())
		var number = (v.Info.Sender).String()
		num := number[2:12]
		match1, _ := regexp.MatchString("order:", msg)
		if match1 == true {
			text := "Order Placed"
			PlaceOrder(msg[5:], num)
			err := utils.SendMessage(text, c, v.Info.Chat)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

func PlaceOrder(del string, num string) error {
	url := "https://oms-bebot-backend.onrender.com/api/orders/"
	order := Order{Number: num, OrderDel: del}

	jsonPayload, err := json.Marshal(order)
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
	return err
}
