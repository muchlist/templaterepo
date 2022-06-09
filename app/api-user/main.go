package main

import (
	"fmt"
	"log"
	"monorepo/conf"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// ==============================================
	// example framework use fiber
	// ==============================================

	// load config
	cfg := conf.Load()

	// create fiber app
	app := fiber.New(
		fiber.Config{
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				// Status code defaults to 500
				code := fiber.StatusInternalServerError
				if e, ok := err.(*fiber.Error); ok {
					code = e.Code
				}
				// Send custom error page
				ctx.Status(code).JSON(fiber.Map{
					"data":  nil,
					"error": err.Error(),
				})
				return nil
			},
		},
	)

	// gracefully shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		fmt.Println("gracefully shutting down...")
		_ = app.Shutdown()
	}()

	// fulfill dependency and routing
	PrefareRoute(app, *cfg)

	// blocking and listen for fiber app
	port := "8081"
	if cfg.User.ServerPort != "" {
		port = cfg.User.ServerPort
	}
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Panic()
	}
	// cleanup app
	fmt.Println("Running cleanup tasks...")
}
