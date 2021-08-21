package time

import (
	"time"
)

// Provider provides an interface for abstracting time
type Provider interface {
	Now() time.Time
}

type timeProvider struct {
}

// NewTimeProvider Provider constructor that returns the default time provider
func NewTimeProvider() Provider {
	return timeProvider{}
}

// Now returns the current time
func (t timeProvider) Now() time.Time {
	return time.Now()
}
