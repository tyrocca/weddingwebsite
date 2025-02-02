// This file contains the implementation of the in-memory repository for the guest entity.

package memrepo

import (
	"log"
	"sync"
	"weddingwebsite/domain"
	"weddingwebsite/repository"
)

type MemoryGuestRepository struct {
	guests   map[int]domain.Guest
	parties  map[int]domain.Party
	contacts map[int]domain.GuestContact

	guestNextID   int32
	partyNextID   int32
	contactNextID int32
	mu            sync.RWMutex
}

func NewMemoryGuestRepository() repository.GuestRepository {
	return &MemoryGuestRepository{
		guests:        make(map[int]domain.Guest),
		parties:       make(map[int]domain.Party),
		contacts:      make(map[int]domain.GuestContact),
		guestNextID:   1,
		partyNextID:   1,
		contactNextID: 1,
		mu:            sync.RWMutex{},
	}
}

// Unique Constraints
const NO_ID = -1

// No two contacts can have the same email or phone number
func checkContactConflict(contacts map[int]domain.GuestContact, email, phone string, allowedID int) error {
	for _, contact := range contacts {
		// if the contact already exists you don't check it for the conflict
		if contact.ID != allowedID {
			if contact.Email == email || contact.Phone == phone {
				log.Println("Contact already exists", contact, "Tried to make", email, phone)
				return domain.ErrContactAlreadyExists
			}
		}
	}
	return nil
}

// No two parties can have the same name
func checkPartyConflict(parties map[int]domain.Party, name string, allowedID int) error {
	for _, party := range parties {
		// if the party already exists you don't check it for the conflict
		if party.ID != allowedID {
			if party.Name == name {
				log.Println("Party already exists", party, "Tried to make", name)
				return domain.ErrPartyAlreadyExists
			}
		}
	}
	return nil
}

// No Two Guests can have the same name in the same party or contact
func checkGuestConflict(guests map[int]domain.Guest, name string, allowedID int, allowedPartyID int, allowedContactID int) error {
	for _, guest := range guests {
		// if the guest already exists you don't check it for the conflict
		if guest.ID != allowedID && guest.Name == name {
			if guest.PartyID == allowedPartyID || guest.ContactID != allowedContactID {
				log.Println("Guest already exists in this scope", guest, "Tried to make", name)
				return domain.ErrGuestAlreadyExists
			}
		}
	}
	return nil
}

// Helpers
func findPartyByName(parties map[int]domain.Party, name string) *domain.Party {
	for _, party := range parties {
		if party.Name == name {
			return &party
		}
	}
	return nil
}

// GuestRepository implementation

// PARTY METHODS

// CreateParty Create a new party.
func (r *MemoryGuestRepository) CreateParty(name string, partySize int) (*domain.Party, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if the name is already taken
	if err := checkPartyConflict(r.parties, name, NO_ID); err != nil {
		return nil, err
	}

	party := domain.Party{
		ID:        int(r.partyNextID),
		Name:      name,
		PartySize: partySize,
	}
	r.parties[party.ID] = party
	r.partyNextID++
	return &party, nil
}

// UpdateParty updates a party.
func (r *MemoryGuestRepository) UpdateParty(partyID int, name string, partySize int) (*domain.Party, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if the name is already taken (and not the current party)
	if err := checkPartyConflict(r.parties, name, partyID); err != nil {
		return nil, err
	}

	// Find Party
	party, ok := r.parties[partyID]
	if !ok {
		return nil, domain.ErrPartyNotFound
	}
	party.Name = name
	party.PartySize = partySize
	r.parties[partyID] = party

	return &party, nil
}

// DeleteParty deletes a party.
func (r *MemoryGuestRepository) DeleteParty(partyID int) error {
	// Note this doesn't delete the guests associated with the party
	r.mu.Lock()
	defer r.mu.Unlock()

	// Find Party
	_, ok := r.parties[partyID]
	if !ok {
		return domain.ErrPartyNotFound
	}

	// Ensure delete is allowed
	for _, guest := range r.guests {
		if guest.PartyID == partyID {
			return domain.ErrPartyNotEmpty
		}
	}

	delete(r.parties, partyID)
	return nil
}

// AddGuestsToParty adds guests to a party.
func (r *MemoryGuestRepository) AddGuestsToParty(partyID int, guestIDs []int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Find Party
	_, ok := r.parties[partyID]
	if !ok {
		return domain.ErrPartyNotFound
	}

	// ensure all guests exist
	guests := make([]*domain.Guest, 0, len(guestIDs))
	for _, guestID := range guestIDs {
		guest, ok := r.guests[guestID]
		if !ok {
			return domain.ErrGuestNotFound
		}
		guests = append(guests, &guest)
	}

	// assign party ID to guests
	for _, guestToAdd := range guests {
		// Adding guest should not break unique constriant
		if err := checkGuestConflict(r.guests, guestToAdd.Name, guestToAdd.ID, partyID, guestToAdd.ContactID); err != nil {
			return err
		}
		guestToAdd.PartyID = partyID
		r.guests[guestToAdd.ID] = *guestToAdd
	}

	return nil
}

// GetPartyByID finds a party by ID.
func (r *MemoryGuestRepository) GetPartyByID(partyID int) (*domain.Party, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Find a party
	party, ok := r.parties[partyID]
	if !ok {
		return nil, domain.ErrPartyNotFound
	}
	return &party, nil
}

// GetPartyByName finds a party by name.
func (r *MemoryGuestRepository) GetPartyByName(partyName string) (*domain.Party, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	party := findPartyByName(r.parties, partyName)
	if party == nil {
		return nil, domain.ErrPartyNotFound
	}
	return party, nil
}

// ListAllParties lists all parties.
func (r *MemoryGuestRepository) ListAllParties() ([]*domain.Party, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	parties := make([]*domain.Party, 0, len(r.parties))
	for _, party := range r.parties {
		parties = append(parties, &party)
	}
	return parties, nil
}

// ListGuestsByParty lists all guests for a party.
func (r *MemoryGuestRepository) ListGuestsByParty(partyID int) ([]*domain.Guest, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, ok := r.parties[partyID]
	if !ok {
		return nil, domain.ErrPartyNotFound
	}

	guests := make([]*domain.Guest, 0)
	for _, guest := range r.guests {
		if guest.PartyID == partyID {
			guests = append(guests, &guest)
		}
	}
	return guests, nil
}

// GUEST CONTACT METHODS

// Guest Helper methods

// findContactByInfo finds a contact by email or phone.
func findContactByInfo(contacts map[int]domain.GuestContact, email, phone string) *domain.GuestContact {
	for _, contact := range contacts {
		if contact.Email == email || contact.Phone == phone {
			return &contact
		}
	}
	return nil
}

// CreateGuestContact creates a new guest contact.
func (r *MemoryGuestRepository) CreateGuestContact(email, phone, phoneCountryCode string) (*domain.GuestContact, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if the contact already exists
	if err := checkContactConflict(r.contacts, email, phone, NO_ID); err != nil {
		return nil, err
	}

	contact := domain.GuestContact{
		ID:               int(r.contactNextID),
		Email:            email,
		Phone:            phone,
		PhoneCountryCode: phoneCountryCode,
	}
	r.contacts[contact.ID] = contact
	r.contactNextID++
	return &contact, nil
}

func (r *MemoryGuestRepository) UpdateGuestContact(contactID int, email, phone, phoneCountryCode string) (*domain.GuestContact, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := checkContactConflict(r.contacts, email, phone, NO_ID); err != nil {
		return nil, err
	}

	// Ensure exists
	contact, ok := r.contacts[contactID]
	if !ok {
		return nil, domain.ErrContactNotFound
	}
	contact.Email = email
	contact.Phone = phone
	contact.PhoneCountryCode = phoneCountryCode
	r.contacts[contactID] = contact

	return &contact, nil
}

// DeleteGuestContact deletes a guest contact.
func (r *MemoryGuestRepository) DeleteGuestContact(contactID int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Ensure exists
	_, ok := r.contacts[contactID]
	if !ok {
		return domain.ErrContactNotFound
	}

	// Cascade delete
	for _, guest := range r.guests {
		if guest.ContactID == contactID {
			return domain.ErrContactNotEmpty
		}
	}
	delete(r.contacts, contactID)
	return nil
}

// FindContactByInfo finds a contact by email or phone.
func (r *MemoryGuestRepository) FindContactByInfo(email string, phone string) (*domain.GuestContact, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	contact := findContactByInfo(r.contacts, email, phone)
	if contact == nil {
		return nil, domain.ErrContactNotFound
	}
	return contact, nil
}

// ListGuestsForContact lists all guests for a given set of contact info
func (r *MemoryGuestRepository) ListGuestsForContact(contactID int) ([]*domain.Guest, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, ok := r.contacts[contactID]
	if !ok {
		return nil, domain.ErrContactNotFound
	}

	guests := make([]*domain.Guest, 0)
	for _, guest := range r.guests {
		if guest.ContactID == contactID {
			guests = append(guests, &guest)
		}
	}
	return guests, nil
}

// GUEST METHODS

// Helper methods

// CreateGuest creates a new guest.
func (r *MemoryGuestRepository) CreateGuest(name, alias string, contactID, partyID int, attending bool) (*domain.Guest, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := checkGuestConflict(r.guests, name, NO_ID, partyID, contactID); err != nil {
		return nil, err
	}

	guest := domain.Guest{
		ID:        int(r.guestNextID),
		Name:      name,
		Alias:     alias,
		ContactID: contactID,
		PartyID:   partyID,
		Attending: attending,
	}
	r.guests[guest.ID] = guest
	r.guestNextID++
	return &guest, nil
}

// UpdateGuest updates a guest.
func (r *MemoryGuestRepository) UpdateGuest(guestID int, name, alias string, contactID, partyID int, attending bool) (*domain.Guest, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := checkGuestConflict(r.guests, name, guestID, partyID, contactID); err != nil {
		return nil, err
	}

	guest, ok := r.guests[guestID]
	if !ok {
		return nil, domain.ErrGuestNotFound
	}
	guest.Name = name
	guest.Alias = alias
	guest.ContactID = contactID
	guest.PartyID = partyID
	guest.Attending = attending
	r.guests[guestID] = guest

	return &guest, nil
}

// DeleteGuest deletes a guest.
func (r *MemoryGuestRepository) DeleteGuest(guestID int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.guests[guestID]
	if !ok {
		return domain.ErrGuestNotFound
	}

	delete(r.guests, guestID)
	return nil
}

// GetGuestByID finds a guest by ID.
func (r *MemoryGuestRepository) GetGuestByID(guestID int) (*domain.Guest, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	guest, ok := r.guests[guestID]
	if !ok {
		return nil, domain.ErrGuestNotFound
	}
	return &guest, nil
}

// ListGuestsByIds finds guests by ID.
func (r *MemoryGuestRepository) ListGuestsByIds(guestIDs []int) ([]*domain.Guest, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	guests := make([]*domain.Guest, 0, len(guestIDs))
	for _, guestID := range guestIDs {
		guest, ok := r.guests[guestID]
		if !ok {
			return nil, domain.ErrGuestNotFound
		}
		guests = append(guests, &guest)
	}
	return guests, nil
}

// ListAllGuests lists all guests.
func (r *MemoryGuestRepository) ListAllGuests() ([]*domain.Guest, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	guests := make([]*domain.Guest, 0, len(r.guests))
	for _, guest := range r.guests {
		guests = append(guests, &guest)
	}
	return guests, nil
}
