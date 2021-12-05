package commands

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/domain/crag"
)

//DeleteCragRequest Command Model
type DeleteCragRequest struct {
	CragID uuid.UUID
}

//DeleteCragRequestHandler Handler Struct with Dependencies
type DeleteCragRequestHandler interface {
	Handle(command DeleteCragRequest) error
}

type deleteCragRequestHandler struct {
	repo crag.Repository
}

//NewDeleteCragRequestHandler Handler constructor
func NewDeleteCragRequestHandler(repo crag.Repository) DeleteCragRequestHandler {
	return deleteCragRequestHandler{repo: repo}
}

//Handle Handlers the DeleteCragRequest request
func (h deleteCragRequestHandler) Handle(command DeleteCragRequest) error {
	crag, err := h.repo.GetByID(command.CragID)
	if crag == nil {
		return fmt.Errorf("the provided crag id does not exist")
	}
	if err != nil {
		return err
	}
	return h.repo.Delete(command.CragID)

}
