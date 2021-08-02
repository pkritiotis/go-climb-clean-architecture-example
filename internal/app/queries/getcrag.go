package queries

import (
	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean/internal/app/services"
	"github.com/pkritiotis/go-clean/internal/domain"
)

//GetCragQuery Model of the Handler
type GetCragQuery struct {
	CragID uuid.UUID
}

//GetCragQueryHandler Contains the dependencies of the Handler
type GetCragQueryHandler struct {
	repo services.CragRepository
}

//NewGetCragQueryHandler Handler Constructor
func NewGetCragQueryHandler(repo services.CragRepository) GetCragQueryHandler {
	return GetCragQueryHandler{repo: repo}
}

//Handle Handlers the GetCragQuery query
func (h GetCragQueryHandler) Handle(query GetCragQuery) (*domain.Crag, error) {
	return h.repo.GetCrag(query.CragID)
}
