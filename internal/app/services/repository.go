package services

import (
	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean/internal/domain"
)

//CragRepository Repository Interface for crags
type CragRepository interface {
	GetCrag(id uuid.UUID) (*domain.Crag, error)
	GetCrags() ([]domain.Crag, error)
	AddCrag(crag domain.Crag) error
	UpdateCrag(crag domain.Crag) error
	DeleteCrag(id uuid.UUID) error
}
