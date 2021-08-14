package commands

import (
	"github.com/pkritiotis/go-climb/internal/app/common"
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
	uuidProvider common.UUIDProvider
	timeProvider common.TimeProvider
	repo         services.CragRepository
}

//NewAddCragCommandHandler Initializes an AddCommandHandler
func NewAddCragCommandHandler(uuidProvider common.UUIDProvider, timeProvider common.TimeProvider, repo services.CragRepository) AddCragCommandHandler {
	return addCragCommandHandler{uuidProvider: uuidProvider, timeProvider: timeProvider, repo: repo}
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
	return h.repo.Add(crag)

}
