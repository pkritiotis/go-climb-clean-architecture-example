package crag

import (
	"github.com/google/uuid"
)

// Repository Interface for crags
type Repository interface {
	GetByID(id uuid.UUID) (*Crag, error)
	GetAll() ([]Crag, error)
	Add(crag Crag) error
	Update(crag Crag) error
	Delete(id uuid.UUID) error
}
