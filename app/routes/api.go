package routes

import (
	"github.com/f100x/go-whatsapp-proxy/app/controllers"
	"github.com/f100x/go-whatsapp-proxy/app/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setup(app *fiber.App, controller *controllers.Controller) {
	app.Use(cors.New())
	app.Use("/api", middlewares.Auth)

	app.Get("/api/user/login", controller.Login)
	app.Get("/api/user/logout", controller.Logout)
	app.Get("/api/user/execute", controller.ExecuteScript)

	app.Post("/api/message/send", controller.SendMessage)

	app.Get("/api/message/last", controller.LastMessage)

	app.Get("/api/tool/check-number/:number", controller.NumberInfo)
}
