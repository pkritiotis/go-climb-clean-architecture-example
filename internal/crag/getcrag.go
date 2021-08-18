package crag

import (
	"github.com/google/uuid"
)

//GetCragQuery Model of the Handler
type GetCragQuery struct {
	CragID uuid.UUID
}

//GetCragQueryHandler provides an interfaces to handle a GetCragQuery and return a *QueryResult
type GetCragQueryHandler interface {
	Handle(query GetCragQuery) (*QueryResult, error)
}

type getCragQueryHandler struct {
	repo Repository
}

//NewGetCragQueryHandler Handler Constructor
func NewGetCragQueryHandler(repo Repository) GetCragQueryHandler {
	return getCragQueryHandler{repo: repo}
}

//Handle Handlers the GetCragQuery query
func (h getCragQueryHandler) Handle(query GetCragQuery) (*QueryResult, error) {
	crag, err := h.repo.GetByID(query.CragID)
	var result *QueryResult
	if crag != nil && err == nil {
		result = &QueryResult{ID: crag.ID, Name: crag.Name, Desc: crag.Desc, Country: crag.Country, CreatedAt: crag.CreatedAt}
	}
	return result, err
}
