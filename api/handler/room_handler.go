package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dr4g0n7ly/hotel-management-system-golang/db"
	"github.com/dr4g0n7ly/hotel-management-system-golang/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookRoomParams struct {
	FromDate   time.Time `json:"fromDate"`
	TillDate   time.Time `json:"tillDate"`
	NumPersons int       `json:"numPersons"`
}

func (p BookRoomParams) validate() error {
	now := time.Now()
	if now.After(p.FromDate) {
		return fmt.Errorf("invalid booking date")
	}
	if p.FromDate.After(p.TillDate) || p.FromDate == p.TillDate {
		fmt.Println(p.FromDate, p.TillDate)
		return fmt.Errorf("invalid booking dates")
	}
	return nil
}

type RoomHandler struct {
	store *db.Store
}

func NewRoomHandler(store db.Store) *RoomHandler {
	return &RoomHandler{
		store: &store,
	}
}

func (h *RoomHandler) HandleBookRoom(c *fiber.Ctx) error {
	var params BookRoomParams
	if err := c.BodyParser(&params); err != nil {
		fmt.Println("Error parsing req.body")
		return err
	}
	if err := params.validate(); err != nil {
		return err
	}

	roomID := c.Params("id")
	roomOID, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		return err
	}
	user, ok := c.Context().Value("user").(*types.User)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(GenericResponse{
			Type: "error",
			Msg:  "internal server error",
		})
	}

	filter := bson.M{
		"roomId": roomOID,
		"fromDate": bson.M{
			"$gte": params.FromDate,
		},
		"tillDate": bson.M{
			"$lte": params.TillDate,
		},
		"cancelled": false,
	}

	bookings, err := h.store.Booking.GetBookings(c.Context(), filter)
	if err != nil {
		return err
	}
	if len(bookings) > 0 {
		return c.Status(http.StatusBadRequest).JSON(GenericResponse{
			Type: "error",
			Msg:  fmt.Sprintf("room already booked"),
		})
	}

	booking := types.Booking{
		UserID:     user.ID,
		RoomID:     roomOID,
		FromDate:   params.FromDate,
		TillDate:   params.TillDate,
		NumPersons: params.NumPersons,
	}

	insertedBooking, err := h.store.Booking.InsertBooking(c.Context(), &booking)
	if err != nil {
		return err
	}
	fmt.Println("\nInserted Booking:", insertedBooking)

	return nil
}
