package queries

import (
	"github.com/pkritiotis/go-clean/internal/app/services"
	"github.com/pkritiotis/go-clean/internal/domain"
)

//GetAllCragsQueryHandler Contains the dependencies of the Handler
type GetAllCragsQueryHandler struct {
	repo services.CragRepository
}

//NewGetAllCragsQueryHandler Handler constructor
func NewGetAllCragsQueryHandler(repo services.CragRepository) GetAllCragsQueryHandler {
	return GetAllCragsQueryHandler{repo: repo}
}

//Handle Handles the query
func (h GetAllCragsQueryHandler) Handle() ([]domain.Crag, error) {

	return h.repo.GetCrags()
}
