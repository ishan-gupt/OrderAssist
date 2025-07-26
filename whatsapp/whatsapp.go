package whatsapp

import (
	"BeBot/handlers"
	"BeBot/utils"
	"context"
	"fmt"
	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
	"os"
)

var client *whatsmeow.Client

func Connect() error {
	dbLog := waLog.Stdout("Database", "INFO", true)
	container, err := sqlstore.New(context.Background(), "sqlite3", "file:wpp_store.db?_foreign_keys=on", dbLog)
	if err != nil {
		return fmt.Errorf("failed to create database container: %v", err)
	}

	deviceStore, err := container.GetFirstDevice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get device store: %v", err)
	}

	clientLog := waLog.Stdout("Client", "INFO", true)
	c := whatsmeow.NewClient(deviceStore, clientLog)
	handlers.SetHandlers(c)

	if c.Store.ID == nil {
		var name = "BeBot"
		utils.NameChange(name)
		qrChan, _ := c.GetQRChannel(context.Background())
		err = c.Connect()
		if err != nil {
			return fmt.Errorf("failed to connect client: %v", err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				fmt.Println("QR code:", evt.Code)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		err = c.Connect()
		if err != nil {
			return fmt.Errorf("failed to connect client: %v", err)
		}
	}

	client = c
	fmt.Println("WhatsApp client connected successfully")

	return nil
}

func Client() *whatsmeow.Client {
	return client
}

func Disconnect() {
	if client != nil {
		client.Disconnect()
		fmt.Println("WhatsApp client disconnected")
	}
}
