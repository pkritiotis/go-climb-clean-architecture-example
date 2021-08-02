package commands

import (
	"github.com/pkritiotis/go-clean/internal/app/common"
	"github.com/pkritiotis/go-clean/internal/app/services"
	"github.com/pkritiotis/go-clean/internal/domain"
)

//AddCragCommand Model of AddCragCommandHandler
type AddCragCommand struct {
	Name    string
	Desc    string
	Country string
}

//AddCragCommandHandler Struct that allows handling AddCragCommand
type AddCragCommandHandler struct {
	uuidProvider common.UUIDProvider
	timeProvider common.TimeProvider
	repo         services.CragRepository
}

//NewAddCragCommandHandler Initializes an AddCommandHandler
func NewAddCragCommandHandler(uuidProvider common.UUIDProvider, timeProvider common.TimeProvider, repo services.CragRepository) AddCragCommandHandler {
	return AddCragCommandHandler{uuidProvider: uuidProvider, timeProvider: timeProvider, repo: repo}
}

//Handle Handles the AddCragCommand
func (h AddCragCommandHandler) Handle(command AddCragCommand) error {
	crag := domain.Crag{
		ID:        h.uuidProvider.NewUUID(),
		Name:      command.Name,
		Desc:      command.Desc,
		Country:   command.Country,
		CreatedAt: h.timeProvider.Now(),
	}
	return h.repo.AddCrag(crag)

}
