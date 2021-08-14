package services

import (
	"github.com/google/uuid"
	"github.com/pkritiotis/go-climb/internal/domain"
)

//CragRepository Repository Interface for crags
type CragRepository interface {
	GetByID(id uuid.UUID) (*domain.Crag, error)
	GetAll() ([]domain.Crag, error)
	Add(crag domain.Crag) error
	Update(crag domain.Crag) error
	Delete(id uuid.UUID) error
}
