package crag

import (
	"github.com/google/uuid"
	"time"
)

// QueryResult is the return model of Crag Query Handlers
type QueryResult struct {
	ID        uuid.UUID
	Name      string
	Desc      string
	Country   string
	CreatedAt time.Time
}
