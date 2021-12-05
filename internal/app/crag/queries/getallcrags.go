package queries

import (
	"github.com/google/uuid"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/domain/crag"
	"time"
)

// GetAllCragsResult is the result of the GetAllCragsRequest Query
type GetAllCragsResult struct {
	ID        uuid.UUID
	Name      string
	Desc      string
	Country   string
	CreatedAt time.Time
}

//GetAllCragsRequestHandler Contains the dependencies of the Handler
type GetAllCragsRequestHandler interface {
	Handle() ([]GetAllCragsResult, error)
}

type getAllCragsRequestHandler struct {
	repo crag.Repository
}

//NewGetAllCragsRequestHandler Handler constructor
func NewGetAllCragsRequestHandler(repo crag.Repository) GetAllCragsRequestHandler {
	return getAllCragsRequestHandler{repo: repo}
}

//Handle Handles the query
func (h getAllCragsRequestHandler) Handle() ([]GetAllCragsResult, error) {

	res, err := h.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var result []GetAllCragsResult
	for _, crag := range res {
		result = append(result, GetAllCragsResult{ID: crag.ID, Name: crag.Name, Desc: crag.Desc, Country: crag.Country, CreatedAt: crag.CreatedAt})
	}
	return result, nil
}
