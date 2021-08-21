package uuid

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// MockProvider mocks uuid provider
type MockProvider struct {
	mock.Mock
}

// NewUUID returns the mocked uuid
func (m MockProvider) NewUUID() uuid.UUID {
	args := m.Called()
	return args.Get(0).(uuid.UUID)
}
