package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dr4g0n7ly/hotel-management-system-golang/db"
	"github.com/dr4g0n7ly/hotel-management-system-golang/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	hotel := types.Hotel{
		Name:     "Golden Hotel",
		Location: "Bucharest, Romania",
	}
	room := types.Room{
		Type:  types.Single,
		Price: 150,
	}
	_ = room

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("seeding the database")
	fmt.Println(insertedHotel)
}
