package main

import(
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/silabig1294/testbank/routes"
	db "github.com/silabig1294/testbank/database"
)


func main()  {
	db.Connect()
	app := fiber.New()

	app.Use(cors.New(
		cors.Config{
			AllowCredentials: true,
		},
	))

	routes.Setapp(app)

	app.Listen(":8081")

}