package commands

import (
	"github.com/pkritiotis/go-climb/internal/app/notification"
	"github.com/pkritiotis/go-climb/internal/domain/crag"
	"github.com/pkritiotis/go-climb/internal/pkg/time"
	"github.com/pkritiotis/go-climb/internal/pkg/uuid"
)

//AddCragRequest Model of AddCragRequestHandler
type AddCragRequest struct {
	Name    string
	Desc    string
	Country string
}

//AddCragRequestHandler Struct that allows handling AddCragRequest
type AddCragRequestHandler interface {
	Handle(command AddCragRequest) error
}

type addCragRequestHandler struct {
	uuidProvider        uuid.Provider
	timeProvider        time.Provider
	repo                crag.Repository
	notificationService notification.Service
}

//NewAddCragRequestHandler Initializes an AddCommandHandler
func NewAddCragRequestHandler(uuidProvider uuid.Provider, timeProvider time.Provider, repo crag.Repository, notificationService notification.Service) AddCragRequestHandler {
	return addCragRequestHandler{uuidProvider: uuidProvider, timeProvider: timeProvider, repo: repo, notificationService: notificationService}
}

//Handle Handles the AddCragRequest
func (h addCragRequestHandler) Handle(command AddCragRequest) error {
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