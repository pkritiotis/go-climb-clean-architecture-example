package timeutil

import (
	"github.com/stretchr/testify/mock"
	"time"
)

// MockTimeProvider mocks TimeProvider
type MockTimeProvider struct {
	mock.Mock
}

// Now returns the mocked time
func (m MockTimeProvider) Now() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}
