package repository

import "weddingwebsite/domain"

var fakeGuest = domain.Guest{}
var fakeParty = domain.Party{}
var fakeGuestContact = domain.GuestContact{}

// GuestRepository provides access to the guest data store.
// Models: Guest, Party, GuestContact
// Use cases:
//   - Create a new guest.
//   - List all guests.
//   - Update a guest.
//   - Delete a guest.
//   - Find a guest by ID.

type GuestRepository interface {
	// Party Operations

	// CreateParty creates a new party.
	CreateParty(name string, partySize int) (*domain.Party, error)

	// UpdateParty updates a party.
	UpdateParty(partyID int, name string, partySize int) (*domain.Party, error)

	// DeleteParty deletes a party.
	DeleteParty(partyID int) error
	// AddGuestsToParty adds guests to a party.
	AddGuestsToParty(partyID int, guestIDs []int) error

	// GetPartyByID finds a party by ID.
	GetPartyByID(partyID int) (*domain.Party, error)
	GetPartyByName(partyName string) (*domain.Party, error)

	// ListAllParties lists all parties
	ListAllParties() ([]*domain.Party, error)

	// ListGuestsByParty lists all guests for a party.
	ListGuestsByParty(partyID int) ([]*domain.Guest, error)

	// Communication Operations

	// CreateGuestContact creates a new guest contact.
	CreateGuestContact(email, phone, phoneCountryCode string) (*domain.GuestContact, error)

	// UpdateGuestContact updates a guest contact.
	UpdateGuestContact(contactID int, email, phone, phoneCountryCode string) (*domain.GuestContact, error)

	// DeleteGuestContact deletes a guest contact.
	DeleteGuestContact(contactID int) error

	// Get Contact by ID
	FindContactByInfo(email string, phone string) (*domain.GuestContact, error)

	// ListGuestsForContact lists all guests for a given set of contact info
	ListGuestsForContact(contactID int) ([]*domain.Guest, error)

	// Guest Management

	// CreateGuest creates a new guest.
	CreateGuest(name, alias string, contactID, partyID int, attending bool) (*domain.Guest, error)

	// UpdateGuest updates a guest.
	UpdateGuest(guestID int, name, alias string, contactID, partyID int, attending bool) (*domain.Guest, error)

	// DeleteGuest deletes a guest.
	DeleteGuest(guestID int) error

	// GetGuestByID finds a guest by ID.
	GetGuestByID(guestID int) (*domain.Guest, error)

	// ListGuestsByIds finds guests by ID.
	ListGuestsByIds(guestIDs []int) ([]*domain.Guest, error)

	// ListAllGuests lists all guests.
	ListAllGuests() ([]*domain.Guest, error)
}
