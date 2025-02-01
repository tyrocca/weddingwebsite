package repository

import "weddingwebsite/domain"

// RelationshipRepository provides access to the relationship data store.
// Operates on: RelationshipLink, RelationshipType
// Use cases:
//   - Create a new relationship type.
//   - List all relationship types.
//   - Delete a relationship type.

//   - Create a relationship between two guests.
//   - Search for relationships between guests.

type RelationshipRepository interface {
	// CreateRelationshipType creates a new type of relationship between guests.
	CreateRelationshipType(name string, creator int) (*domain.RelationshipType, error)

	// ListRelationshipTypes lists all relationship types
	ListAllRelationshipTypes() ([]*domain.RelationshipType, error)

	// DeleteRelationshipType deletes a relationship type.
	DeleteRelationshipType(relationshipTypeID int) error

	// CreateRelationship creates a new relationship between two guests.
	CreateRelationship(fromGuestID, toGuestID, relationshipTypeID int, description string) (*domain.RelationshipLink, error)

	// UpdateRelationship you can update the description or type
	UpdateRelationship(relationshipID, relationshipTypeID int, description string) (*domain.RelationshipLink, error)

	// DeleteRelationship deletes a relationship between two guests.
	DeleteRelationship(relationshipID int) error

	// ListRelationships lists all relationships between guests.
	ListAllRelationships(hydrateGuests bool) ([]*domain.RelationshipLink, error)

	// ListRelationshipsForGuest lists all relationships for a guest.
	ListRelationshipsByGuest(guestID int, hydrateGuests bool) ([]*domain.RelationshipLink, error)
}
