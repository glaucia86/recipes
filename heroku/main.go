// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

// Install and configure heroku: https://devcenter.heroku.com/articles/getting-started-with-go#set-up
// You need to read the PORT env from heroku and you need to define the Procfile

// Deploy the app: https://devcenter.heroku.com/articles/getting-started-with-go#deploy-the-app

package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber"
)

func main() {

	// Create new Fiber instance
	app := fiber.New()

	// Create new GET route
	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("Hello Heroku")
	})

	// Get the PORT from heroku env
	port := os.Getenv("PORT")

	// Verify if heroku provided the port or not
	if port == "" {
		port = "3000"
		log.Print("$PORT == 3000")
	}

	// Start server on http://${heroku-url}:${port}
	app.Listen(port)
}
