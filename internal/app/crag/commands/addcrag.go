package commands

import (
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/app/notification"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/domain/crag"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/pkg/time"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/pkg/uuid"
)

//AddCragRequest Model of CreateCragRequestHandler
type AddCragRequest struct {
	Name    string
	Desc    string
	Country string
}

//CreateCragRequestHandler Struct that allows handling AddCragRequest
type CreateCragRequestHandler interface {
	Handle(command AddCragRequest) error
}

type addCragRequestHandler struct {
	uuidProvider        uuid.Provider
	timeProvider        time.Provider
	repo                crag.Repository
	notificationService notification.Service
}

//NewAddCragRequestHandler Initializes an AddCommandHandler
func NewAddCragRequestHandler(uuidProvider uuid.Provider, timeProvider time.Provider, repo crag.Repository, notificationService notification.Service) CreateCragRequestHandler {
	return addCragRequestHandler{uuidProvider: uuidProvider, timeProvider: timeProvider, repo: repo, notificationService: notificationService}
}

//Handle Handles the AddCragRequest
func (h addCragRequestHandler) Handle(req AddCragRequest) error {
	c := crag.Crag{
		ID:        h.uuidProvider.NewUUID(),
		Name:      req.Name,
		Desc:      req.Desc,
		Country:   req.Country,
		CreatedAt: h.timeProvider.Now(),
	}
	err := h.repo.Add(c)
	if err != nil {
		return err
	}
	n := notification.Notification{
		Subject: "New crag added",
		Message: "A new crag with name '" + c.Name + "' was added in the repository",
	}
	return h.notificationService.Notify(n)
}
