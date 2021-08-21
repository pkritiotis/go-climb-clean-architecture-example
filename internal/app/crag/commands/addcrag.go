package commands

import (
	"github.com/pkritiotis/go-climb/internal/app/notification"
	"github.com/pkritiotis/go-climb/internal/domain/crag"
	"github.com/pkritiotis/go-climb/internal/pkg/time"
	"github.com/pkritiotis/go-climb/internal/pkg/uuid"
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
	uuidProvider        uuid.Provider
	timeProvider        time.Provider
	repo                crag.Repository
	notificationService notification.Service
}

//NewAddCragCommandHandler Initializes an AddCommandHandler
func NewAddCragCommandHandler(uuidProvider uuid.Provider, timeProvider time.Provider, repo crag.Repository, notificationService notification.Service) AddCragCommandHandler {
	return addCragCommandHandler{uuidProvider: uuidProvider, timeProvider: timeProvider, repo: repo, notificationService: notificationService}
}

//Handle Handles the AddCragCommand
func (h addCragCommandHandler) Handle(command AddCragCommand) error {
	crag := crag.Crag{
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
	n := notification.Notification{
		Subject: "New crag added",
		Message: "A new crag with name '" + crag.Name + "' was added in the repository",
	}
	return h.notificationService.Notify(n)
}
