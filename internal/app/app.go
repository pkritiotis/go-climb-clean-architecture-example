package app

import (
	"github.com/pkritiotis/go-clean/internal/app/common"
	"github.com/pkritiotis/go-clean/internal/app/commands"
	"github.com/pkritiotis/go-clean/internal/app/queries"
	"github.com/pkritiotis/go-clean/internal/app/services"
)

//Queries Contains all available query handlers of this app
type Queries struct {
	GetAllCragsHandler queries.GetAllCragsQueryHandler
	GetCragHandler     queries.GetCragQueryHandler
}

//Commands Contains all available command handlers of this app
type Commands struct {
	AddCragHandler    commands.AddCragCommandHandler
	UpdateCragHandler commands.UpdateCragCommandHandler
	DeleteCragHandler commands.DeleteCragCommandHandler
}

//App Contains the grouped queries and commands of the app layer
type App struct {
	Queries  Queries
	Commands Commands
}

// NewApp Bootstraps Application Layer dependencies
func NewApp(up common.UUIDProvider, tp common.TimeProvider, cragRepo services.CragRepository) App {
	return App{
		Queries: Queries{
			GetAllCragsHandler: queries.NewGetAllCragsQueryHandler(cragRepo),
			GetCragHandler:     queries.NewGetCragQueryHandler(cragRepo),
		},
		Commands: Commands{
			AddCragHandler:    commands.NewAddCragCommandHandler(up, tp, cragRepo),
			UpdateCragHandler: commands.NewUpdateCragCommandHandler(cragRepo),
			DeleteCragHandler: commands.NewDeleteCragCommandHandler(cragRepo),
		},
	}
}
