package interfaceadapters

import (
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/app/notification"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/domain/crag"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/interfaceadapters/notification/console"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/interfaceadapters/storage/memory"
)

type Services struct {
	NotificationService notification.Service
	CragRepository      crag.Repository
}

func NewServices() Services {
	return Services{
		NotificationService: console.NewNotificationService(),
		CragRepository:      memory.NewRepo(),
	}
}