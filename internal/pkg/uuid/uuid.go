package uuid

import "github.com/google/uuid"

// Provider abstracts the uuid generation
type Provider interface {
	NewUUID() uuid.UUID
}

// NewUUIDProvider constructor that returns default uuid generation
func NewUUIDProvider() Provider {
	return uuidProvider{}
}

type uuidProvider struct {
}

// NewUUID generates a new UUID
func (u uuidProvider) NewUUID() uuid.UUID {
	return uuid.New()
}
