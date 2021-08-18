package app

import (
	"github.com/pkritiotis/go-climb/internal/app/commands"
	"github.com/pkritiotis/go-climb/internal/app/common"
	"github.com/pkritiotis/go-climb/internal/app/queries"
	services2 "github.com/pkritiotis/go-climb/internal/app/services"
	"github.com/pkritiotis/go-climb/internal/domain/services"
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
func NewApp(cragRepo services.CragRepository, ns services2.NotificationService) App {
	up := common.NewUUIDProvider()
	tp := common.NewTimeProvider()
	return App{
		Queries: Queries{
			GetAllCragsHandler: queries.NewGetAllCragsQueryHandler(cragRepo),
			GetCragHandler:     queries.NewGetCragQueryHandler(cragRepo),
		},
		Commands: Commands{
			AddCragHandler:    commands.NewAddCragCommandHandler(up, tp, cragRepo, ns),
			UpdateCragHandler: commands.NewUpdateCragCommandHandler(cragRepo),
			DeleteCragHandler: commands.NewDeleteCragCommandHandler(cragRepo),
		},
	}
}
