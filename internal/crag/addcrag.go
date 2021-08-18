package crag

import (
	appService "github.com/pkritiotis/go-climb/internal/notification"
	"github.com/pkritiotis/go-climb/internal/timeutil"
	"github.com/pkritiotis/go-climb/internal/uuidutil"
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
	uuidProvider        uuidutil.UUIDProvider
	timeProvider        timeutil.TimeProvider
	repo                Repository
	notificationService appService.Service
}

//NewAddCragCommandHandler Initializes an AddCommandHandler
func NewAddCragCommandHandler(uuidProvider uuidutil.UUIDProvider, timeProvider timeutil.TimeProvider, repo Repository, notificationService appService.Service) AddCragCommandHandler {
	return addCragCommandHandler{uuidProvider: uuidProvider, timeProvider: timeProvider, repo: repo, notificationService: notificationService}
}

//Handle Handles the AddCragCommand
func (h addCragCommandHandler) Handle(command AddCragCommand) error {
	crag := Crag{
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
