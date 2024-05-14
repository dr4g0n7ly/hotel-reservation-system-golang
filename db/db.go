package db

const (
	DBNAME     = "hotel-reservation"
	TestDBNAME = "test-hotel-reservation"
	DBURI      = "mongodb://localhost:27017/"
)

type QueryFilter struct {
	Limit  int64
	Page   int64
	Rating int
}

type Store struct {
	User    UserStore
	Hotel   HotelStore
	Room    RoomStore
	Booking BookingStore
}
