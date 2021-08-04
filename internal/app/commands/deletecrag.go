package commands

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean/internal/app/services"
)

//DeleteCragCommand Command Model
type DeleteCragCommand struct {
	CragID uuid.UUID
}

//DeleteCragCommandHandler Handler Struct with Dependencies
type DeleteCragCommandHandler interface {
	Handle(command DeleteCragCommand) error
}

type deleteCragCommandHandler struct {
	repo services.CragRepository
}

//NewDeleteCragCommandHandler Handler constructor
func NewDeleteCragCommandHandler(repo services.CragRepository) DeleteCragCommandHandler {
	return deleteCragCommandHandler{repo: repo}
}

//Handle Handlers the DeleteCragCommand request
func (h deleteCragCommandHandler) Handle(command DeleteCragCommand) error {
	crag, err := h.repo.GetCrag(command.CragID)
	if crag == nil {
		return fmt.Errorf("the provided crag id does not exist")
	}
	if err != nil {
		return err
	}
	return h.repo.DeleteCrag(command.CragID)

}
