package time

import (
	"github.com/stretchr/testify/mock"
	"time"
)

// MockProvider mocks Provider
type MockProvider struct {
	mock.Mock
}

// Now returns the mocked time
func (m MockProvider) Now() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}
