package queries

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean/internal/app/services"
)

type CragQueryResult struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
}

//GetAllCragsQueryHandler Contains the dependencies of the Handler
type GetAllCragsQueryHandler interface {
	Handle() ([]CragQueryResult, error)
}

type getAllCragsQueryHandler struct {
	repo services.CragRepository
}

//NewGetAllCragsQueryHandler Handler constructor
func NewGetAllCragsQueryHandler(repo services.CragRepository) GetAllCragsQueryHandler {
	return getAllCragsQueryHandler{repo: repo}
}

//Handle Handles the query
func (h getAllCragsQueryHandler) Handle() ([]CragQueryResult, error) {

	res, err := h.repo.GetCrags()
	if err != nil {
		return nil, err
	}
	var crags []CragQueryResult
	for _, crag := range res {
		crags = append(crags, CragQueryResult{ID: crag.ID, Name: crag.Name, Desc: crag.Desc, Country: crag.Country, CreatedAt: crag.CreatedAt})
	}
	return crags, nil
}
