package container

import (
	"weddingwebsite/repository/memrepo"
	"weddingwebsite/services"
)

func NewInMemoryContainer() *Container {
	// Make all the in-memory repositories
	guestRepo := memrepo.NewMemoryGuestRepository()

	return &Container{
		GuestSvc: services.NewGuestService(guestRepo),
		AdminSvc: services.NewAdminService(guestRepo),
	}
}
