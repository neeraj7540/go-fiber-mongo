package main

import (
	"log"
	"os"
	"github.com/gofiber/fiber"
	"github.com/kamva/mgm/v3"
    "go.mongodb.org/mongo-driver/mongo/options"
	"go-fiber-mongo/controllers"

)

func init() {
	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	if len(connectionString) == 0 {
		connectionString = "mongodb+srv://neeraj:neeraj@cluster0.dpk92.mongodb.net/test"
	}

	err := mgm.SetDefaultConfig(nil, "users", options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	
}

func main() {
    app := fiber.New()

	app.Get("/api/users", controllers.GetAllUser)
	app.Post("/api/user", controllers.CreateUser)
	app.Get("/api/user/:id", controllers.GetUserByID)
	app.Patch("/api/user/:id", controllers.UpdateUser)
    

    app.Listen(":3000")
}