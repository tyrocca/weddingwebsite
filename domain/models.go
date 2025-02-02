package domain

import (
	"errors"
	"time"
)

// Guest represents an attendee at the wedding.
type Guest struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Alias     string    `json:"alias"`
	ContactID int       `json:"contactId"`
	PartyID   int       `json:"partyId"`
	Attending bool      `json:"attending"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// Links
	Contact *GuestContact `json:"contact,omitempty"`
	Party   *Party        `json:"party,omitempty"`
}

func (g *Guest) GetParty() *Party {
	// should this be hydrated?
	return g.Party
}

func (g *Guest) GetContact() *GuestContact {
	return g.Contact
}

// Unique together? Name + PartyID?

// GuestContact holds contact details for a guest.
type GuestContact struct {
	ID               int       `json:"id"`
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
	PhoneCountryCode string    `json:"phoneCountryCode"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

// Party groups multiple guests together.
type Party struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	PartySize int       `json:"partySize"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type LocationTypeEnum string

// Types of locations.
const (
	Lodging      LocationTypeEnum = "lodging"
	Airport      LocationTypeEnum = "airport"
	WeddingEvent LocationTypeEnum = "wedding_event"
)

// Location represents a place associated with the wedding.
type Location struct {
	ID        int                    `json:"id"`
	Address   string                 `json:"address"`
	Alias     string                 `json:"alias"`
	Type      LocationTypeEnum       `json:"type"` // lodging | airport | wedding event
	Metadata  map[string]interface{} `json:"metadata"`
	CreatedAt time.Time              `json:"createdAt"`
	UpdatedAt time.Time              `json:"updatedAt"`
}

// Event represents a wedding-related event.
type Event struct {
	ID          int                    `json:"id"`
	LocationID  int                    `json:"locationId"`
	When        time.Time              `json:"when"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Metadata    map[string]interface{} `json:"metadata"`
	CreatedAt   time.Time              `json:"createdAt"`
	UpdatedAt   time.Time              `json:"updatedAt"`
	// Links
	Location *Location `json:"location,omitempty"`
}

// RSVPLink represents an RSVP for a guest.
type RSVPLink struct {
	ID        int                    `json:"id"`
	GuestID   int                    `json:"guestId"`
	Attending bool                   `json:"attending"`
	Metadata  map[string]interface{} `json:"metadata"`
	CreatedAt time.Time              `json:"createdAt"`
	UpdatedAt time.Time              `json:"updatedAt"`
	// Links
	Guest *Guest `json:"guest,omitempty"`
}

// RelationshipType defines a type of relationship between guests.
type RelationshipType struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Creator   int       `json:"creator"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// RelationshipLink represents a relationship between two guests.
type RelationshipLink struct {
	ID                 int       `json:"id"`
	FromGuestID        int       `json:"fromGuestId"`
	ToGuestID          int       `json:"toGuestId"`
	RelationshipTypeID int       `json:"relationshipTypeId"`
	Description        string    `json:"description"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	// Links
	FromGuest        *Guest            `json:"fromGuest,omitempty"`
	ToGuest          *Guest            `json:"toGuest,omitempty"`
	RelationshipType *RelationshipType `json:"relationshipType,omitempty"`
}

// LodgingLink represents a guest's lodging details.
type LodgingLink struct {
	ID         int       `json:"id"`
	LocationID int       `json:"locationId"`
	GuestID    int       `json:"guestId"`
	Name       string    `json:"name"`
	Public     bool      `json:"public"`
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	// Links
	Location *Location `json:"location,omitempty"`
	Guest    *Guest    `json:"guest,omitempty"`
}

// FlightLink associates guests with flights.
type FlightLink struct {
	ID        int       `json:"id"`
	GuestID   int       `json:"guestId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// Links
	Guest  *Guest  `json:"guest,omitempty"`
	Flight *Flight `json:"flight,omitempty"`
}

// Flight represents flight details.
type Flight struct {
	Time          time.Time `json:"time"`
	Airline       string    `json:"airline"`
	FlightNumber  string    `json:"flightNumber"`
	LodgingID     int       `json:"lodgingId"`
	FromAirportID int       `json:"fromAirportId"`
	ToAirportID   int       `json:"toAirportId"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	// Links
	Lodging     *LodgingLink `json:"lodging,omitempty"`
	FromAirport *Location    `json:"fromAirport,omitempty"`
	ToAirport   *Location    `json:"toAirport,omitempty"`
}

// Not Found Errors
var (
	ErrPartyNotFound            = errors.New("party not found")
	ErrGuestNotFound            = errors.New("guest not found")
	ErrContactNotFound          = errors.New("contact not found")
	ErrLocationNotFound         = errors.New("location not found")
	ErrEventNotFound            = errors.New("event not found")
	ErrRSVPNotFound             = errors.New("rsvp not found")
	ErrRelationshipTypeNotFound = errors.New("relationship type not found")
	ErrRelationshipLinkNotFound = errors.New("relationship link not found")
	ErrLodgingLinkNotFound      = errors.New("lodging link not found")
	ErrFlightLinkNotFound       = errors.New("flight link not found")
	ErrFlightNotFound           = errors.New("flight not found")
)

// Integrity Errors
var (
	// Creation / Update Errors
	ErrPartyAlreadyExists            = errors.New("party already exists")
	ErrGuestAlreadyExists            = errors.New("guest already exists")
	ErrContactAlreadyExists          = errors.New("contact already exists")
	ErrLocationAlreadyExists         = errors.New("location already exists")
	ErrEventAlreadyExists            = errors.New("event already exists")
	ErrRSVPAlreadyExists             = errors.New("rsvp already exists")
	ErrRelationshipTypeAlreadyExists = errors.New("relationship type already exists")
	ErrRelationshipLinkAlreadyExists = errors.New("relationship link already exists")
	ErrLodgingLinkAlreadyExists      = errors.New("lodging link already exists")
	ErrFlightLinkAlreadyExists       = errors.New("flight link already exists")
	ErrFlightAlreadyExists           = errors.New("flight already exists")

	// Delete Errors
	ErrPartyNotEmpty   = errors.New("party not empty")
	ErrContactNotEmpty = errors.New("contact not empty")
)

//
