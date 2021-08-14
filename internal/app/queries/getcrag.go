package queries

import (
	"github.com/google/uuid"
	"github.com/pkritiotis/go-climb/internal/domain/services"
)

//GetCragQuery Model of the Handler
type GetCragQuery struct {
	CragID uuid.UUID
}

//GetCragQueryHandler provides an interfaces to handle a GetCragQuery and return a *CragQueryResult
type GetCragQueryHandler interface {
	Handle(query GetCragQuery) (*CragQueryResult, error)
}

type getCragQueryHandler struct {
	repo services.CragRepository
}

//NewGetCragQueryHandler Handler Constructor
func NewGetCragQueryHandler(repo services.CragRepository) GetCragQueryHandler {
	return getCragQueryHandler{repo: repo}
}

//Handle Handlers the GetCragQuery query
func (h getCragQueryHandler) Handle(query GetCragQuery) (*CragQueryResult, error) {
	crag, err := h.repo.GetByID(query.CragID)
	var result *CragQueryResult
	if crag != nil && err == nil {
		result = &CragQueryResult{ID: crag.ID, Name: crag.Name, Desc: crag.Desc, Country: crag.Country, CreatedAt: crag.CreatedAt}
	}
	return result, err
}
