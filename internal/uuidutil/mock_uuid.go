package uuidutil

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// MockUUIDProvider mocks uuid provider
type MockUUIDProvider struct {
	mock.Mock
}

// NewUUID returns the mocked uuid
func (m MockUUIDProvider) NewUUID() uuid.UUID {
	args := m.Called()
	return args.Get(0).(uuid.UUID)
}
