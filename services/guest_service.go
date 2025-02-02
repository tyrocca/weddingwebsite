package services

import (
	"weddingwebsite/domain"
	"weddingwebsite/repository"
)

type GuestService interface {
	// Party Operations
	// bulk loading endpoint
	putGuest(partyName, guestName, guestAlias, guestEmail, guestPhone, guestPhoneCountryCode string, partySize int) (*domain.Guest, error)

	createParty(name string, partySize int) (*domain.Party, error)
	updateParty(partyID int, name string, partySize int) (*domain.Party, error)
	deleteParty(partyID int) error

	addGuestsToParty(partyID int, guestIDs []int) error
	listAllParties() ([]*domain.Party, error)
	getPartyByID(partyID int) (*domain.Party, error)
	listGuestsByParty(partyID int) ([]*domain.Guest, error)
}

type guestService struct {
	repo repository.GuestRepository
}

func (g guestService) putGuest(partyName, guestName, guestAlias, guestEmail, guestPhone, guestPhoneCountryCode string, partySize int) (*domain.Guest, error) {
	//TODO implement me
	panic("implement me")
}

func (g guestService) createParty(name string, partySize int) (*domain.Party, error) {
	//TODO implement me
	panic("implement me")
}

func (g guestService) updateParty(partyID int, name string, partySize int) (*domain.Party, error) {
	//TODO implement me
	panic("implement me")
}

func (g guestService) deleteParty(partyID int) error {
	//TODO implement me
	panic("implement me")
}

func (g guestService) addGuestsToParty(partyID int, guestIDs []int) error {
	//TODO implement me
	panic("implement me")
}

func (g guestService) listAllParties() ([]*domain.Party, error) {
	//TODO implement me
	panic("implement me")
}

func (g guestService) getPartyByID(partyID int) (*domain.Party, error) {
	//TODO implement me
	panic("implement me")
}

func (g guestService) listGuestsByParty(partyID int) ([]*domain.Guest, error) {
	//TODO implement me
	panic("implement me")
}

// Ensure that guestService implements the GuestService interface
var _ GuestService = &guestService{}

func NewGuestService(repo repository.GuestRepository) GuestService {
	return &guestService{
		repo: repo,
	}
}
