package main

import (
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/app"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/inputports"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/interfaceadapters"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/pkg/time"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/pkg/uuid"
)

func main() {
	interfaceAdapterServices := interfaceadapters.NewServices()
	tp := time.NewTimeProvider()
	up := uuid.NewUUIDProvider()
	appServices := app.NewServices(interfaceAdapterServices.CragRepository, interfaceAdapterServices.NotificationService, up, tp)
	inputPortsServices := inputports.NewServices(appServices)
	inputPortsServices.Server.ListenAndServe(":8080")
}
