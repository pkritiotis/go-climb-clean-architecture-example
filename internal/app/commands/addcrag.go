package commands

import (
	"github.com/pkritiotis/go-climb/internal/app/common"
	appService "github.com/pkritiotis/go-climb/internal/app/services"
	"github.com/pkritiotis/go-climb/internal/domain"
	"github.com/pkritiotis/go-climb/internal/domain/services"
)

//AddCragCommand Model of AddCragCommandHandler
type AddCragCommand struct {
	Name    string
	Desc    string
	Country string
}

//AddCragCommandHandler Struct that allows handling AddCragCommand
type AddCragCommandHandler interface {
	Handle(command AddCragCommand) error
}

type addCragCommandHandler struct {
	uuidProvider        common.UUIDProvider
	timeProvider        common.TimeProvider
	repo                services.CragRepository
	notificationService appService.NotificationService
}

//NewAddCragCommandHandler Initializes an AddCommandHandler
func NewAddCragCommandHandler(uuidProvider common.UUIDProvider, timeProvider common.TimeProvider, repo services.CragRepository, notificationService appService.NotificationService) AddCragCommandHandler {
	return addCragCommandHandler{uuidProvider: uuidProvider, timeProvider: timeProvider, repo: repo, notificationService: notificationService}
}

//Handle Handles the AddCragCommand
func (h addCragCommandHandler) Handle(command AddCragCommand) error {
	crag := domain.Crag{
		ID:        h.uuidProvider.NewUUID(),
		Name:      command.Name,
		Desc:      command.Desc,
		Country:   command.Country,
		CreatedAt: h.timeProvider.Now(),
	}
	err := h.repo.Add(crag)
	if err != nil {
		return err
	}
	n := appService.Notification{
		Subject: "New crag added",
		Message: "A new crag with name '" + crag.Name + "' was added in the repository",
	}
	return h.notificationService.Notify(n)
}
