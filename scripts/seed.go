package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dr4g0n7ly/hotel-management-system-golang/db"
	"github.com/dr4g0n7ly/hotel-management-system-golang/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func seedUser(client *mongo.Client, isAdmin bool, email, firstname, lastname, password string) {
	userStore := db.NewMongoUserStore(client)
	user, err := types.NewUserFromParams(types.CreateUserParams{
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
		Password:  password,
	})
	if err != nil {
		log.Fatal(err)
	}
	user.IsAdmin = isAdmin
	insertedUser, err := userStore.InsertUser(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertedUser)
}

func seedHotel(client *mongo.Client, name string, location string, rating int) {
	hotelStore := db.NewMongoHotelStore(client)
	roomStore := db.NewMongoRoomStore(client, hotelStore)

	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
		Rating:   rating,
	}
	insertedHotel, err := hotelStore.InsertHotel(context.Background(), &hotel)
	if err != nil {
		log.Fatal(err)
	}
	rooms := []types.Room{
		{
			Type:  types.Single,
			Price: 150.50,
		},
		{
			Type:  types.Single,
			Price: 150.50,
		},
		{
			Type:  types.Deluxe,
			Price: 200,
		},
		{
			Type:  types.Suite,
			Price: 1399,
		},
	}
	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		insertedRoom, err := roomStore.InsertRoom(context.Background(), &room)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(insertedRoom)
	}
	fmt.Println(insertedHotel)

}

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	if err := client.Database(db.DBNAME).Drop(context.Background()); err != nil {
		log.Fatal(err)
	}

	seedUser(client, false, "foo.bar@gmail.com", "foo", "bar", "password")
	seedUser(client, false, "jack.baz@gmail.com", "jack", "baz", "password")
	seedUser(client, true, "admin@gmail.com", "admin", "admin", "adminpassword")
	seedHotel(client, "Park Hyatt", "Hyderabad", 6)
	seedHotel(client, "Grand Hotel", "Bucharest", 8)

	fmt.Println("seeding the database")
}
