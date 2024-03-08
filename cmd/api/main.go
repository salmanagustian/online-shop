package main

import (
	"log"
	"online-shop/apps/auth"
	"online-shop/apps/products"
	"online-shop/apps/transaction"
	"online-shop/external/database"
	infrafiber "online-shop/infra/fiber"
	"online-shop/internal/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	filename := "cmd/api/config.yaml"

	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connected")
	}

	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
	})

	router.Use(infrafiber.Trace())

	auth.Init(router, db)
	products.Init(router, db)
	transaction.Init(router, db)

	router.Listen(config.Cfg.App.Port)
}
