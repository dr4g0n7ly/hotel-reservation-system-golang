package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name     string               `bson:"name" json:"name"`
	Location string               `bson:"location" json:"location"`
	Rooms    []primitive.ObjectID `bson:"rooms" json:"rooms"`
	Rating   int                  `bson:"rating" json:"rating"`
}

type RoomType int

const (
	_ RoomType = iota
	Single
	Double
	Deluxe
	Suite
)

type Room struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Type       RoomType           `bson:"type" json:"type"`
	Price      float64            `bson:"price" json:"price"`
	HotelID    primitive.ObjectID `bson:"hotelId" json:"hotelId"`
	MaxPersons int                `bson:"maxPersons" json:"maxPersons"`
}
