package repo

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-climb/internal/domain/crag"
)

//inMemory Implements the Repository Interface to provide an in-memory storage provider
type inMemory struct {
	crags map[string]crag.Crag
}

//NewInMemory Constructor
func NewInMemory() crag.Repository {
	crags := make(map[string]crag.Crag)
	return inMemory{crags}
}

//GetCrag Returns the crag with the provided id
func (m inMemory) GetByID(id uuid.UUID) (*crag.Crag, error) {
	crag, ok := m.crags[id.String()]
	if !ok {
		return nil, nil
	}
	return &crag, nil
}

//GetCrags Returns all stored crags
func (m inMemory) GetAll() ([]crag.Crag, error) {
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

//AddCrag Adds the provided crag
func (m inMemory) Add(crag crag.Crag) error {
	m.crags[crag.ID.String()] = crag
	return nil
}

//UpdateCrag Updates the provided crag
func (m inMemory) Update(crag crag.Crag) error {
	m.crags[crag.ID.String()] = crag
	return nil
}

//DeleteCrag Deletes the crag with the provided id
func (m inMemory) Delete(id uuid.UUID) error {
	_, exists := m.crags[id.String()]
	if !exists {
		return fmt.Errorf("id %v not found", id.String())
	}
	delete(m.crags, id.String())
	return nil
}
