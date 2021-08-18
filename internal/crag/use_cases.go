package crag

import (
	"github.com/pkritiotis/go-climb/internal/notification"
	"github.com/pkritiotis/go-climb/internal/timeutil"
	"github.com/pkritiotis/go-climb/internal/uuidutil"
)

//Queries Contains all available query handlers of this app
type Queries struct {
	GetAllCragsHandler GetAllCragsQueryHandler
	GetCragHandler     GetCragQueryHandler
}

//Commands Contains all available command handlers of this app
type Commands struct {
	AddCragHandler    AddCragCommandHandler
	UpdateCragHandler UpdateCragCommandHandler
	DeleteCragHandler DeleteCragCommandHandler
}

//UseCases Contains the grouped queries and commands of the app layer
type UseCases struct {
	Queries  Queries
	Commands Commands
}

// NewUseCases Bootstraps Application Layer dependencies
func NewUseCases(cragRepo Repository, ns notification.Service, up uuidutil.UUIDProvider, tp timeutil.TimeProvider) UseCases {
	return UseCases{
		Queries: Queries{
			GetAllCragsHandler: NewGetAllCragsQueryHandler(cragRepo),
			GetCragHandler:     NewGetCragQueryHandler(cragRepo),
		},
		Commands: Commands{
			AddCragHandler:    NewAddCragCommandHandler(up, tp, cragRepo, ns),
			UpdateCragHandler: NewUpdateCragCommandHandler(cragRepo),
			DeleteCragHandler: NewDeleteCragCommandHandler(cragRepo),
		},
	}
}
