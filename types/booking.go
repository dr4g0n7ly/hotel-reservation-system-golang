package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"userId" bson:"userId"`
	RoomID     primitive.ObjectID `json:"roomId" bson:"roomId"`
	FromDate   time.Time          `json:"fromDate" bson:"fromDate"`
	TillDate   time.Time          `json:"tillDate" bson:"tillDate"`
	NumPersons int                `json:"numPersons" bson:"numPersons"`
}
