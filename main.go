package main

import (
	"context"
	"fmt"
	"log"

	api "github.com/dr4g0n7ly/hotel-management-system-golang/api/handler"
	"github.com/dr4g0n7ly/hotel-management-system-golang/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create a new fiber instance with custom config
var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	userStore := db.NewMongoUserStore(client, db.DBNAME)
	userHandler := api.NewUserHandler(userStore)

	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	roomStore := db.NewMongoRoomStore(client, hotelStore)
	hotelHandler := api.NewHotelHandler(hotelStore, roomStore)

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	// user handlers
	apiv1.Get("/users", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	apiv1.Post("user", userHandler.HandlePostUser)
	apiv1.Delete("user/:id", userHandler.HandleDeleteUser)

	// hotel handlers
	apiv1.Get("/hotels", hotelHandler.HandleGetHotels)
	apiv1.Get("/hotels/:id/rooms", hotelHandler.HandleGetRooms)
	app.Listen(":5000")
}
