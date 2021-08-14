package queries

import (
	"github.com/pkritiotis/go-climb/internal/domain/services"
)

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

	res, err := h.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var crags []CragQueryResult
	for _, crag := range res {
		crags = append(crags, CragQueryResult{ID: crag.ID, Name: crag.Name, Desc: crag.Desc, Country: crag.Country, CreatedAt: crag.CreatedAt})
	}
	return crags, nil
}
