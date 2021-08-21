package crag

import (
	"github.com/google/uuid"
	"time"
)

//Crag Model that represents the Crag
type Crag struct {
	ID        uuid.UUID
	Name      string
	Desc      string
	Country   string
	CreatedAt time.Time
}
