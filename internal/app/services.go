package app

import (
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/app/crag/commands"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/app/crag/queries"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/app/notification"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/domain/crag"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/pkg/time"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/pkg/uuid"
)

//Queries Contains all available query handlers of this app
type Queries struct {
	GetAllCragsHandler queries.GetAllCragsRequestHandler
	GetCragHandler     queries.GetCragRequestHandler
}

//Commands Contains all available command handlers of this app
type Commands struct {
	CreateCragHandler commands.CreateCragRequestHandler
	UpdateCragHandler commands.UpdateCragRequestHandler
	DeleteCragHandler commands.DeleteCragRequestHandler
}

//CragServices Contains the grouped queries and commands of the app layer
type CragServices struct {
	Queries  Queries
	Commands Commands
}

//Services contains all exposed services of the application layer
type Services struct {
	CragServices CragServices
}

// NewServices Bootstraps Application Layer dependencies
func NewServices(cragRepo crag.Repository, ns notification.Service, up uuid.Provider, tp time.Provider) Services {
	return Services{
		CragServices: CragServices{
			Queries: Queries{
				GetAllCragsHandler: queries.NewGetAllCragsRequestHandler(cragRepo),
				GetCragHandler:     queries.NewGetCragRequestHandler(cragRepo),
			},
			Commands: Commands{
				CreateCragHandler: commands.NewAddCragRequestHandler(up, tp, cragRepo, ns),
				UpdateCragHandler: commands.NewUpdateCragRequestHandler(cragRepo),
				DeleteCragHandler: commands.NewDeleteCragRequestHandler(cragRepo),
			},
		}}
}
