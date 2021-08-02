package repo

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean/internal/app/services"
	"github.com/pkritiotis/go-clean/internal/domain"
)

//inMemoryRepo Implements the Repository Interface to provide an in-memory storage provider
type inMemoryRepo struct {
	crags map[string]domain.Crag
}

//NewInMemory Constructor
func NewInMemory() services.CragRepository {
	crags := make(map[string]domain.Crag)
	return inMemoryRepo{crags}
}

//GetCrag Returns the crag with the provided id
func (m inMemoryRepo) GetCrag(id uuid.UUID) (*domain.Crag, error) {
	crag, ok := m.crags[id.String()]
	if !ok {
		return nil, nil
	}
	return &crag, nil
}

//GetCrags Returns all stored crags
func (m inMemoryRepo) GetCrags() ([]domain.Crag, error) {
	keys := make([]string, 0)

	for key := range m.crags {
		keys = append(keys, key)
	}

	var values []domain.Crag
	for _, value := range m.crags {
		values = append(values, value)
	}
	return values, nil
}

//AddCrag Adds the provided crag
func (m inMemoryRepo) AddCrag(crag domain.Crag) error {
	m.crags[crag.ID.String()] = crag
	return nil
}

//UpdateCrag Updates the provided crag
func (m inMemoryRepo) UpdateCrag(crag domain.Crag) error {
	m.crags[crag.ID.String()] = crag
	return nil
}

//DeleteCrag Deletes the crag with the provided id
func (m inMemoryRepo) DeleteCrag(id uuid.UUID) error {
	_, exists := m.crags[id.String()]
	if !exists {
		return fmt.Errorf("id %v not found", id.String())
	}
	delete(m.crags, id.String())
	return nil
}
