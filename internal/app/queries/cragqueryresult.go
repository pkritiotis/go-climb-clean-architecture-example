package queries

import (
	"github.com/google/uuid"
	"time"
)

// CragQueryResult is the return model of Crag Query Handlers
type CragQueryResult struct {
	ID        uuid.UUID
	Name      string
	Desc      string
	Country   string
	CreatedAt time.Time
}
