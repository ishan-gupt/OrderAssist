package whatsapp

import (
	"WhatsText/handlers"
	"WhatsText/utils"
	"context"
	"fmt"
	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
	"os"
	"os/signal"
	"syscall"
)

var client *whatsmeow.Client

func Connect() error {
	var name = "WhatsText"
	//fmt.Println("Enter Your Name: ")
	//fmt.Scanln(&name)
	utils.NameChange(name)
	dbLog := waLog.Stdout("Database", "INFO", true)
	container, err := sqlstore.New("sqlite3", "file:wpp_store.db?_foreign_keys=on", dbLog)
	if err != nil {
		return err
	}

	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		return err
	}

	clientLog := waLog.Stdout("Client", "INFO", true)
	c := whatsmeow.NewClient(deviceStore, clientLog)
	handlers.SetHandlers(c)

	if c.Store.ID == nil {
		qrChan, _ := c.GetQRChannel(context.Background())
		err = c.Connect()
		if err != nil {
			return err
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
			return err
		}
	}

	client = c

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch

	return err
}

func Client() *whatsmeow.Client {
	return client
}

func Disconnect() {
	client.Disconnect()
}
