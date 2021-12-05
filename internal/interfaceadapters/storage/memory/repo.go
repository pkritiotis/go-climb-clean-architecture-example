package memory

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/domain/crag"
)

//Repo Implements the Repository Interface to provide an in-memory storage provider
type Repo struct {
	crags map[string]crag.Crag
}

//NewRepo Constructor
func NewRepo() Repo {
	crags := make(map[string]crag.Crag)
	return Repo{crags}
}

//GetByID Returns the crag with the provided id
func (m Repo) GetByID(id uuid.UUID) (*crag.Crag, error) {
	crag, ok := m.crags[id.String()]
	if !ok {
		return nil, nil
	}
	return &crag, nil
}

//GetAll Returns all stored crags
func (m Repo) GetAll() ([]crag.Crag, error) {
	keys := make([]string, 0)

	for key := range m.crags {
		keys = append(keys, key)
	}

	var values []crag.Crag
	for _, value := range m.crags {
		values = append(values, value)
	}
	return values, nil
}

//Add the provided crag
func (m Repo) Add(crag crag.Crag) error {
	m.crags[crag.ID.String()] = crag
	return nil
}

//Update the provided crag
func (m Repo) Update(crag crag.Crag) error {
	m.crags[crag.ID.String()] = crag
	return nil
}

//Delete the crag with the provided id
func (m Repo) Delete(id uuid.UUID) error {
	_, exists := m.crags[id.String()]
	if !exists {
		return fmt.Errorf("id %v not found", id.String())
	}
	delete(m.crags, id.String())
	return nil
}
