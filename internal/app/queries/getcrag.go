package queries

import (
	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean/internal/domain"
	"github.com/pkritiotis/go-clean/internal/domain/services"
)

//GetCragQuery Model of the Handler
type GetCragQuery struct {
	CragID uuid.UUID
}

//GetCragQueryHandler Contains the dependencies of the Handler
type GetCragQueryHandler interface {
	Handle(query GetCragQuery) (*domain.Crag, error)
}

type getCragQueryHandler struct {
	repo services.CragRepository
}

//NewGetCragQueryHandler Handler Constructor
func NewGetCragQueryHandler(repo services.CragRepository) GetCragQueryHandler {
	return getCragQueryHandler{repo: repo}
}

//Handle Handlers the GetCragQuery query
func (h getCragQueryHandler) Handle(query GetCragQuery) (*domain.Crag, error) {
	return h.repo.GetByID(query.CragID)
}
