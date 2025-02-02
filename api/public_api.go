package api

import (
	"context"
	"github.com/labstack/gommon/log"
	"time"
	"weddingwebsite/container"
	"weddingwebsite/domain"
	"weddingwebsite/openapi/genopenapi"
)

type WeddingHandler struct {
	//TODO add the container here
	c *container.Container
}

func (w *WeddingHandler) ListEvents(ctx context.Context, request genopenapi.ListEventsRequestObject) (genopenapi.ListEventsResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (w *WeddingHandler) ListGuests(ctx context.Context, request genopenapi.ListGuestsRequestObject) (genopenapi.ListGuestsResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (w *WeddingHandler) GetGuest(ctx context.Context, request genopenapi.GetGuestRequestObject) (genopenapi.GetGuestResponseObject, error) {
	//TODO implement me
	log.Print("GetGuest", request)

	g := domain.Guest{
		ID:        0,
		Name:      "Huck Rocca",
		Alias:     "stinky",
		ContactID: 0,
		PartyID:   0,
		Attending: false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Contact:   nil,
		Party:     nil,
	}

	responseGuest := genopenapi.Guest{
		Id:        &g.ID,
		Name:      &g.Name,
		Alias:     &g.Alias,
		ContactId: &g.ContactID,
		PartyId:   &g.PartyID,
		Attending: &g.Attending,
		CreatedAt: &g.CreatedAt,
		UpdatedAt: &g.UpdatedAt,
	}

	return genopenapi.GetGuest200JSONResponse(responseGuest), nil

}

// Ensure that WeddingHandler implements the genopenapi.ServerInterface
var _ genopenapi.StrictServerInterface = (*WeddingHandler)(nil)

func NewWeddingHandler(c *container.Container) genopenapi.StrictServerInterface {
	return &WeddingHandler{
		c: c,
	}
}
