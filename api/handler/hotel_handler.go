package api

import (
	"github.com/dr4g0n7ly/hotel-management-system-golang/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	hotelStore db.HotelStore
	roomStore  db.RoomStore
}

func NewHotelHandler(store db.Store) *HotelHandler {
	return &HotelHandler{
		hotelStore: store.Hotel,
		roomStore:  store.Room,
	}
}

type HotelQueryParams struct {
	Rooms bool
}

func (h *HotelHandler) HandleGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"hotelId": oid}
	rooms, err := h.roomStore.GetRooms(c.Context(), filter)
	if err != nil {
		return err
	}
	return c.JSON(rooms)
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var queryFilter db.QueryFilter
	if err := c.QueryParser(&queryFilter); err != nil {
		return err
	}

	hotels, err := h.hotelStore.GetHotels(c.Context(), nil, &queryFilter)
	if err != nil {
		return err
	}
	return c.JSON(hotels)
}

func (h *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID()
	}
	filter := bson.M{"_id": oid}
	var queryFilter db.QueryFilter
	if err := c.QueryParser(&queryFilter); err != nil {
		return err
	}
	hotels, err := h.hotelStore.GetHotels(c.Context(), filter, &queryFilter)
	if err != nil {
		return err
	}
	return c.JSON(hotels)
}
