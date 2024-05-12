package fixtures

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dr4g0n7ly/hotel-management-system-golang/db"
	"github.com/dr4g0n7ly/hotel-management-system-golang/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUser(store *db.Store, fn, ln string, admin bool) *types.User {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		Email:     fmt.Sprintf("%s@%s.com", fn, ln),
		FirstName: fn,
		LastName:  ln,
		Password:  fmt.Sprintf("%s_%s", fn, ln),
	})
	if err != nil {
		log.Fatal(err)
	}
	user.IsAdmin = admin
	insertedUser, err := store.User.InsertUser(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	return insertedUser
}

func AddRoom(store *db.Store, roomType types.RoomType, price float64, hotelID primitive.ObjectID, maxP int) *types.Room {
	room := &types.Room{
		Type:       roomType,
		Price:      price,
		HotelID:    hotelID,
		MaxPersons: maxP,
	}
	insertedRoom, err := store.Room.InsertRoom(context.Background(), room)
	if err != nil {
		log.Fatal(err)
	}
	return insertedRoom
}

func AddHotel(store *db.Store, name, loc string, rating int, rooms []primitive.ObjectID) *types.Hotel {
	var roomIDs = rooms
	if rooms == nil {
		roomIDs = []primitive.ObjectID{}
	}
	hotel := types.Hotel{
		Name:     name,
		Location: loc,
		Rooms:    roomIDs,
		Rating:   rating,
	}
	insertedHotel, err := store.Hotel.InsertHotel(context.Background(), &hotel)
	if err != nil {
		log.Fatal(err)
	}
	return insertedHotel
}

func AddBooking(store *db.Store, uid, rid primitive.ObjectID, from, till time.Time) *types.Booking {
	booking := &types.Booking{
		UserID:   uid,
		RoomID:   rid,
		FromDate: from,
		TillDate: till,
	}
	insertedBooking, err := store.Booking.InsertBooking(context.Background(), booking)
	if err != nil {
		log.Fatal(err)
	}
	return insertedBooking
}
