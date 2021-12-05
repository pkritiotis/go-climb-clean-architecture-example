package commands

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/domain/crag"
)

//UpdateCragRequest Update Model
type UpdateCragRequest struct {
	ID      uuid.UUID
	Name    string
	Desc    string
	Country string
}

//UpdateCragRequestHandler Contains the dependencies of the handler
type UpdateCragRequestHandler interface {
	Handle(command UpdateCragRequest) error
}

type updateCragRequestHandler struct {
	repo crag.Repository
}

//NewUpdateCragRequestHandler Constructor
func NewUpdateCragRequestHandler(repo crag.Repository) UpdateCragRequestHandler {
	return updateCragRequestHandler{repo: repo}
}

//Handle Handles the update request
func (h updateCragRequestHandler) Handle(command UpdateCragRequest) error {
	crag, err := h.repo.GetByID(command.ID)
	if crag == nil {
		return fmt.Errorf("the provided crag id does not exist")
	}
	if err != nil {
		return err
	}

	crag.Name = command.Name
	crag.Desc = command.Desc
	crag.Country = command.Country

	return h.repo.Update(*crag)

}
