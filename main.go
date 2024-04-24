package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/kshwedha/juke-box/src/api"
	"github.com/kshwedha/juke-box/src/common/config"
)

func main() {
	config.Init()

	fmt.Println("Starting server...")

	app := fiber.New()

	api.Route(app)

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	if err := app.Listen(":8080"); err != nil {
		log.Panic(err)
	}

	fmt.Println("Running cleanup tasks...")

}
