package queries

import (
	"github.com/google/uuid"
	"time"
)

// CragQueryResult is the return model of Crag Query Handlers
type CragQueryResult struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
}
