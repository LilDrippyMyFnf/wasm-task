package main

import (
	"os"

	"github.com/jgbz/wasm-task/cmd/server/routes"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	runApplication()
}

func runApplication() {

	err := godotenv.Load("../.env")
	if err != nil {
		panic(err.Error())
	}

	r := routes.NewRoutes()

	r.App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	r.RegisterRoutesGet("src")
	port := os.Getenv("PORT")

	r.App.Static("/", "../web/html")
	r.App.Listen(":" + port)
}
