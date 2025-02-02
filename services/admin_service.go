package services

/*
AdminService provides access to the admin operations.

Models: Guest, Party, GuestContact, Location, Event, RelationshipLink, RelationshipType

This is the only service that can
- Create a new party
- Update a party

// I am going to need to make a form to quickly add guests to a party
// if no party make a party

Use cases:
- Create a new party
	- Bulk load guests into the system


*/

import (
	"log"
	"weddingwebsite/domain"
	"weddingwebsite/repository"
)

type createPartyOp struct {
	name string `json:"name"`
}

type createGuestOp struct {
	Name      string `json:"name"`
	Alias     string `json:"alias"`
	Attending bool   `json:"attending"`
}

type createContactOp struct {
	email            string
	phone            string
	phoneCountryCode string
}

type createCompoundGuestOp struct {
	contact *createContactOp
	guests  []*createGuestOp
}

type compoundCreateOperation struct {
	party        *createPartyOp
	contactables []*createCompoundGuestOp
}

type AdminService interface {
	// Loading Operations
	createParty(op *createPartyOp, num int) (*domain.Party, error)
	createContact(op *createContactOp) (*domain.GuestContact, error)
	createGuest(op *createGuestOp, partyID int, contactID int) (*domain.Guest, error)

	// bulk loading endpoint
	compoundCreateGuests(op *compoundCreateOperation) ([]*domain.Guest, error)
}

/////////////////////////////
// SERVICE IMPLEMENTATION
/////////////////////////////

type adminService struct {
	guestRepo repository.GuestRepository
}

// createParty creates a new party.
func (s *adminService) createParty(op *createPartyOp, num int) (*domain.Party, error) {
	// Does this fail if the party already exists?
	party, err := s.guestRepo.GetPartyByName(op.name)
	if err != nil {
		return nil, err
	}
	if party != nil {
		log.Println("Party already exists", party, "Tried to make", op)
		return party, nil
	}
	return s.guestRepo.CreateParty(op.name, num)
}

// createContact creates a new guest contact.
func (s *adminService) createContact(op *createContactOp) (*domain.GuestContact, error) {
	// What do we do if there is multiple of same value?
	contact, err := s.guestRepo.FindContactByInfo(op.email, op.phone)
	if err != nil {
		return nil, err
	}
	if contact != nil {
		log.Println("Contact already exists", contact, "Tried to make", op)
		return contact, nil
	}
	return s.guestRepo.CreateGuestContact(op.email, op.phone, op.phoneCountryCode)
}

func (s *adminService) createGuest(op *createGuestOp, partyID int, contactID int) (*domain.Guest, error) {
	// How do we want to handle unique contraints?
	// Do we want to check if the guest already exists?
	return s.guestRepo.CreateGuest(op.Name, op.Alias, contactID, partyID, op.Attending)
}

func (s *adminService) compoundCreateGuests(op *compoundCreateOperation) ([]*domain.Guest, error) {
	// Create the party
	totalGuests := 0
	for _, guestOp := range op.contactables {
		totalGuests += len(guestOp.guests)
	}

	party, err := s.createParty(op.party, totalGuests)
	if err != nil {
		return nil, err
	}

	// Create the guests
	guests := make([]*domain.Guest, 0)
	for _, contactOp := range op.contactables {
		// Create the contact
		contact, err := s.createContact(contactOp.contact)
		if err != nil {
			return nil, err
		}
		for _, guest := range contactOp.guests {
			guest, err := s.createGuest(guest, party.ID, contact.ID)
			if err != nil {
				return nil, err
			}
			guest.Party = party
			guest.Contact = contact
			guests = append(guests, guest)
		}
	}
	return guests, nil

}

/* Externally accessible functions */

var _ AdminService = &adminService{}

func NewAdminService(guestRepo repository.GuestRepository) AdminService {
	return &adminService{
		guestRepo: guestRepo,
	}
}
