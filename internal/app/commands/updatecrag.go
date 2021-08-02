package commands

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean/internal/app/services"
)

//UpdateCragCommand UpdateCrag Model
type UpdateCragCommand struct {
	ID      uuid.UUID
	Name    string
	Desc    string
	Country string
}

//UpdateCragCommandHandler Contains the dependencies of the handler
type UpdateCragCommandHandler struct {
	repo services.CragRepository
}

//NewUpdateCragCommandHandler Constructor
func NewUpdateCragCommandHandler(repo services.CragRepository) UpdateCragCommandHandler {
	return UpdateCragCommandHandler{repo: repo}
}

//Handle Handles the update command
func (h UpdateCragCommandHandler) Handle(command UpdateCragCommand) error {
	crag, err := h.repo.GetCrag(command.ID)
	if crag == nil {
		return fmt.Errorf("the provided crag id does not exist")
	}
	if err != nil {
		return err
	}

	crag.Name = command.Name
	crag.Desc = command.Desc
	crag.Country = command.Country

	return h.repo.UpdateCrag(*crag)

}
