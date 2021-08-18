package crag

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

//MockRepository mocks Repository for testing purposes
type MockRepository struct {
	mock.Mock
}

// GetByID mock
func (m MockRepository) GetByID(id uuid.UUID) (*Crag, error) {
	args := m.Called(id)
	return args.Get(0).(*Crag), args.Error(1)
}

// GetAll mock
func (m MockRepository) GetAll() ([]Crag, error) {
	args := m.Called()
	return args.Get(0).([]Crag), args.Error(1)
}

// Add mock
func (m MockRepository) Add(crag Crag) error {
	args := m.Called(crag)
	return args.Error(0)
}

// Update mock
func (m MockRepository) Update(crag Crag) error {
	args := m.Called(crag)
	return args.Error(0)
}

// Delete mock
func (m MockRepository) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}
