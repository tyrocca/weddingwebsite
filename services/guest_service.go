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

func NewGuestService(repo repository.GuestRepository) GuestService {
	return &guestService{
		repo: repo,
	}
}

func (s *guestService) createParty(name string, partySize int) (*domain.Party, error) {

}
