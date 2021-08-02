package domain

import (
	"github.com/google/uuid"
	"time"
)

//Crag Model that represents the Crag
type Crag struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
}
