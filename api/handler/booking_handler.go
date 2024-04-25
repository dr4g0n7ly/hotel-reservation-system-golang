package api

import (
	"github.com/dr4g0n7ly/hotel-management-system-golang/db"
	"github.com/gofiber/fiber/v2"
)

type BookingHandler struct {
	store *db.Store
}

func NewBookingHandler(store db.Store) *BookingHandler {
	return &BookingHandler{
		store: &store,
	}
}

func (h *BookingHandler) HandleGetBookings(c *fiber.Ctx) error {
	bookings, err := h.store.Booking.GetBookings(c.Context(), nil)
	if err != nil {
		return err
	}
	return c.JSON(bookings)
}

func (h *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.store.Booking.GetBookingByID(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(booking)
}
