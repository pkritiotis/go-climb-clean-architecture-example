package main

import (
	"github.com/pkritiotis/go-climb/internal/app"
	"github.com/pkritiotis/go-climb/internal/inputports"
	"github.com/pkritiotis/go-climb/internal/outputadapters"
	"github.com/pkritiotis/go-climb/internal/pkg/time"
	"github.com/pkritiotis/go-climb/internal/pkg/uuid"
)

func main() {
	outputAdapterServices := outputadapters.NewServices()
	tp := time.NewTimeProvider()
	up := uuid.NewUUIDProvider()
	appServices := app.NewServices(outputAdapterServices.CragRepository, outputAdapterServices.NotificationService, up, tp)
	inputPortsServices := inputports.NewServices(appServices)
	inputPortsServices.Server.ListenAndServe(":8080")
}
