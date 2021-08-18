package timeutil

import (
	"time"
)

// TimeProvider provides an interface for abstracting time
type TimeProvider interface {
	Now() time.Time
}

type timeProvider struct {
}

// NewTimeProvider TimeProvider constructor that returns the default time provider
func NewTimeProvider() TimeProvider {
	return timeProvider{}
}

// Now returns the current time
func (t timeProvider) Now() time.Time {
	return time.Now()
}
