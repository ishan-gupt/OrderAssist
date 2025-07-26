package main

import (
	"BeBot/whatsapp"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start WhatsApp connection
	err := whatsapp.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to WhatsApp: %v", err)
	}
	
	// Wait for shutdown signal
	<-sigChan
	fmt.Println("\nShutting down gracefully...")
	
	// Disconnect from WhatsApp
	whatsapp.Disconnect()
	fmt.Println("Disconnected from WhatsApp")
}
