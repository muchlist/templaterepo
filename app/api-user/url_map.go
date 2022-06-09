package main

import (
	"monorepo/app/api-user/handler"
	"monorepo/conf"
	"monorepo/internal/business/notifserv"
	"monorepo/internal/business/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func PrefareRoute(app *fiber.App, cfg conf.Config) {

	// simple common middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Accept, Authorization",
	}))

	// Dependency Injection
	// fulfill repo
	userRepo := user.NewRepoUser()
	notifier := notifserv.NewNotifServ()

	// fullfill usecase / service / core
	userCore := user.NewUserCore(userRepo, notifier)

	// generate handler
	userHandler := handler.NewUserHandler(userCore)

	// mapping url
	app.Get("/profile", userHandler.Login)

}
