package interfaceadapters

import (
	"github.com/pkritiotis/go-climb/internal/app/notification"
	"github.com/pkritiotis/go-climb/internal/domain/crag"
	"github.com/pkritiotis/go-climb/internal/interfaceadapters/notification/console"
	"github.com/pkritiotis/go-climb/internal/interfaceadapters/storage/memory"
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
