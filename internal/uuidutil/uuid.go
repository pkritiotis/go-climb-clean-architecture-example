package uuidutil

import "github.com/google/uuid"

// UUIDProvider abstracts the uuid generation
type UUIDProvider interface {
	NewUUID() uuid.UUID
}

// NewUUIDProvider constructor that returns default uuid generation
func NewUUIDProvider() UUIDProvider {
	return uuidProvider{}
}

type uuidProvider struct {
}

// NewUUID generates a new UUID
func (u uuidProvider) NewUUID() uuid.UUID {
	return uuid.New()
}
