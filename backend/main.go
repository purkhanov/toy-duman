package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"toy-duman/database"
	"toy-duman/web"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	server := web.NewServer()
	err = server.Start()
	if err != nil {
		log.Fatalf("Error starting web server: %v", err)
		return
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGTERM)

	sig := <-sigCh
	log.Printf("Received signal: %s, stopping web server...", sig)

	err = server.Stop()
	if err != nil {
		log.Printf("Error stopping web server: %v", err)
	}
}
