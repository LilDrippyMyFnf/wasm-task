package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jgbz/wasm-task/cmd/server/routes"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	runApplication()
}

func runApplication() {

	err := godotenv.Load("/wasm-task/.env")
	if err != nil {
		panic(err.Error())
	}

	r := routes.NewRoutes()

	r.App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	r.RegisterRoutesGet("v1")
	port := os.Getenv("PORT")

	r.App.Static("/", "/wasm-task/web/html")
	r.App.Listen(":" + port)
}
