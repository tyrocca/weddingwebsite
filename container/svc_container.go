package container

import "weddingwebsite/services"

// All services are in here
type Container struct {
	GuestSvc services.GuestService
	AdminSvc services.AdminService
}
