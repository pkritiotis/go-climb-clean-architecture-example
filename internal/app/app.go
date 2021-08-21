package app

import (
	"github.com/pkritiotis/go-climb/internal/app/crag/commands"
	"github.com/pkritiotis/go-climb/internal/app/crag/queries"
	"github.com/pkritiotis/go-climb/internal/app/notification"
	"github.com/pkritiotis/go-climb/internal/domain/crag"
	"github.com/pkritiotis/go-climb/internal/pkg/time"
	"github.com/pkritiotis/go-climb/internal/pkg/uuid"
)

//Queries Contains all available query handlers of this app
type Queries struct {
	GetAllCragsHandler queries.GetAllCragsRequestHandler
	GetCragHandler     queries.GetCragRequestHandler
}

//Commands Contains all available command handlers of this app
type Commands struct {
	AddCragHandler    commands.AddCragRequestHandler
	UpdateCragHandler commands.UpdateCragRequestHandler
	DeleteCragHandler commands.DeleteCragRequestHandler
}

//App Contains the grouped queries and commands of the app layer
type App struct {
	Queries  Queries
	Commands Commands
}

// NewApp Bootstraps Application Layer dependencies
func NewApp(cragRepo crag.Repository, ns notification.Service, up uuid.Provider, tp time.Provider) App {
	return App{
		Queries: Queries{
			GetAllCragsHandler: queries.NewGetAllCragsRequestHandler(cragRepo),
			GetCragHandler:     queries.NewGetCragRequestHandler(cragRepo),
		},
		Commands: Commands{
			AddCragHandler:    commands.NewAddCragRequestHandler(up, tp, cragRepo, ns),
			UpdateCragHandler: commands.NewUpdateCragRequestHandler(cragRepo),
			DeleteCragHandler: commands.NewDeleteCragRequestHandler(cragRepo),
		},
	}
}
