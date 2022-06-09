package main

import (
	"fmt"
	"log"
	"monorepo/conf"
	dbpg "monorepo/pkg/db-pg"
	"net/http"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

const version = "1.0.0"

func main() {
	// ==============================================
	// example framework use fiber
	// ==============================================

	// load config
	cfg := conf.Load()

	// init database
	database, err := dbpg.OpenDB(dbpg.Config{
		DSN:          cfg.User.DbDsn,
		MaxOpenConns: int32(cfg.User.DbMaxOpenConn),
		MinOpenConns: int32(cfg.User.DbMinOpenConn),
	})
	if err != nil {
		log.Fatal("connection to database", err)
		panic(err.Error())
	}
	defer database.Close()

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

	// start debug server
	debugMux := debugMux(database)
	go func(mux *http.ServeMux) {
		debugPort := "4000"
		if cfg.User.DebugPort != "" {
			debugPort = cfg.User.DebugPort
		}
		if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", debugPort), mux); err != nil {
			log.Print("serve debug api", err)
		}
	}(debugMux)

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
