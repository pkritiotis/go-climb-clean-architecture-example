// Package crag contains the Crag model.
package crag

import (
	"time"

	"github.com/google/uuid"
)

// Crag Model that represents the Crag
type Crag struct {
	ID        uuid.UUID
	Name      string
	Desc      string
	Country   string
	CreatedAt time.Time
}
