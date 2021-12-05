package queries

import (
	"github.com/google/uuid"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/domain/crag"
	"time"
)

//GetCragRequest Model of the Handler
type GetCragRequest struct {
	CragID uuid.UUID
}

// GetCragResult is the return model of Crag Query Handlers
type GetCragResult struct {
	ID        uuid.UUID
	Name      string
	Desc      string
	Country   string
	CreatedAt time.Time
}

//GetCragRequestHandler provides an interfaces to handle a GetCragRequest and return a *GetCragResult
type GetCragRequestHandler interface {
	Handle(query GetCragRequest) (*GetCragResult, error)
}

type getCragRequestHandler struct {
	repo crag.Repository
}

//NewGetCragRequestHandler Handler Constructor
func NewGetCragRequestHandler(repo crag.Repository) GetCragRequestHandler {
	return getCragRequestHandler{repo: repo}
}

//Handle Handlers the GetCragRequest query
func (h getCragRequestHandler) Handle(query GetCragRequest) (*GetCragResult, error) {
	crag, err := h.repo.GetByID(query.CragID)
	var result *GetCragResult
	if crag != nil && err == nil {
		result = &GetCragResult{ID: crag.ID, Name: crag.Name, Desc: crag.Desc, Country: crag.Country, CreatedAt: crag.CreatedAt}
	}
	return result, err
}
