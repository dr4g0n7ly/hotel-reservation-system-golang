package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	api "github.com/dr4g0n7ly/hotel-management-system-golang/api/handler"
	"github.com/dr4g0n7ly/hotel-management-system-golang/api/middleware"
	"github.com/dr4g0n7ly/hotel-management-system-golang/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create a new fiber instance with custom config
var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		if apiError, ok := err.(api.Error); ok {
			return c.Status(apiError.Code).JSON(apiError)
		}
		apiError := api.NewError(http.StatusInternalServerError, err.Error())
		return c.Status(apiError.Code).JSON(apiError.Err)
	},
}

func main() {
	mongodb_uri := os.Getenv("MONGO_DB_URL")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	var (
		userStore    = db.NewMongoUserStore(client)
		hotelStore   = db.NewMongoHotelStore(client)
		roomStore    = db.NewMongoRoomStore(client, hotelStore)
		bookingStore = db.NewMongoBookingStore(client)
		store        = db.Store{
			User:    userStore,
			Hotel:   hotelStore,
			Room:    roomStore,
			Booking: bookingStore,
		}
	)

	authHandler := api.NewAuthHandler(store)
	userHandler := api.NewUserHandler(store)
	hotelHandler := api.NewHotelHandler(store)
	roomHandler := api.NewRoomHandler(store)
	bookingHandler := api.NewBookingHandler(store)

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1", middleware.JWTAuthentication(userStore))
	admin := apiv1.Group("/admin", middleware.AdminAuth)

	// auth handlers
	app.Post("auth", authHandler.HandleAuthenticate)

	// user handlers
	apiv1.Get("/users", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	apiv1.Post("user", userHandler.HandlePostUser)
	apiv1.Delete("user/:id", userHandler.HandleDeleteUser)

	// hotel handlers
	apiv1.Get("/hotels", hotelHandler.HandleGetHotels)
	apiv1.Get("/hotels/:id/rooms", hotelHandler.HandleGetRooms)
	apiv1.Get("/hotel/:id", hotelHandler.HandleGetHotel)

	// room / booking handlers
	apiv1.Post("/room/:id/book", roomHandler.HandleBookRoom)
	apiv1.Get("/bookings", bookingHandler.HandleGetBookings)
	apiv1.Get("/booking/:id", bookingHandler.HandleGetBooking)
	apiv1.Get("/booking/:id/cancel", bookingHandler.HandleCancelBooking)

	// admin handlers
	admin.Get("/users", userHandler.HandleGetUsers)
	admin.Get("/bookings", bookingHandler.HandleGetBookings)

	http_listen_addr := os.Getenv("HTTP_LISTEN_ADDRESS")
	app.Listen(http_listen_addr)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
