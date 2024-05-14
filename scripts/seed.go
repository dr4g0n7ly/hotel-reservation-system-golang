package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	api "github.com/dr4g0n7ly/hotel-management-system-golang/api/handler"
	"github.com/dr4g0n7ly/hotel-management-system-golang/db"
	"github.com/dr4g0n7ly/hotel-management-system-golang/db/fixtures"
	"github.com/dr4g0n7ly/hotel-management-system-golang/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	if err := client.Database(db.DBNAME).Drop(context.Background()); err != nil {
		log.Fatal(err)
	}

	hotelStore := db.NewMongoHotelStore(client)
	store := &db.Store{
		User:    db.NewMongoUserStore(client),
		Booking: db.NewMongoBookingStore(client),
		Room:    db.NewMongoRoomStore(client, hotelStore),
		Hotel:   hotelStore,
	}

	user := fixtures.AddUser(store, "james", "foo", false)
	admin := fixtures.AddUser(store, "admin", "admin", true)
	hotel := fixtures.AddHotel(store, "hotel", "place", 5, nil)
	room := fixtures.AddRoom(store, types.Deluxe, 500, hotel.ID, 2)
	booking := fixtures.AddBooking(store, user.ID, room.ID, time.Now().AddDate(0, 0, 5), time.Now().AddDate(0, 0, 7))

	for i := 0; i < 100; i++ {
		fixtures.AddHotel(store, fmt.Sprintf("hotel %d", i), fmt.Sprintf("location %d", i), rand.Intn(4)+1, nil)
	}

	fmt.Println("james ->", api.CreateTokenFromUser(user))
	fmt.Println("admin ->", api.CreateTokenFromUser(admin))
	fmt.Println("booking ->", booking.ID)
	return
}
