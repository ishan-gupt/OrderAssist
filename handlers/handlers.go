package handlers

import (
	"BeBot/commands"
	"go.mau.fi/whatsmeow"
)

func SetHandlers(c *whatsmeow.Client) {
	// Add Cat Handler
	c.AddEventHandler(func(evt interface{}) {
		commands.ImageHandler(evt, c)
	})

	// Add HelloWorld Handler
	c.AddEventHandler(func(evt interface{}) {
		commands.TextHandler(evt, c)
	})

	// Add ShoudIDeployToday Handler
	c.AddEventHandler(func(evt interface{}) {
		commands.ApiHandler(evt, c)
	})
	c.AddEventHandler(func(evt interface{}) {
		commands.UserHandler(evt, c)
	})
	c.AddEventHandler(func(evt interface{}) {
		commands.OrderHandler(evt, c)
	})
}
