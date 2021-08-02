package services

import (
	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean/internal/domain"
	"github.com/stretchr/testify/mock"
)

//MockRepository mocks Repository for testing purposes
type MockRepository struct {
	mock.Mock
}

// GetCrag mock
func (m MockRepository) GetCrag(id uuid.UUID) (*domain.Crag, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Crag), args.Error(1)
}

//GetCrags mock
func (m MockRepository) GetCrags() ([]domain.Crag, error) {
	args := m.Called()
	return args.Get(0).([]domain.Crag), args.Error(1)
}

// AddCrag mock
func (m MockRepository) AddCrag(crag domain.Crag) error {
	args := m.Called(crag)
	return args.Error(0)
}

// UpdateCrag mock
func (m MockRepository) UpdateCrag(crag domain.Crag) error {
	args := m.Called(crag)
	return args.Error(0)
}

// DeleteCrag mock
func (m MockRepository) DeleteCrag(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}
