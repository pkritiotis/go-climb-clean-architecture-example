// package main contains the entry point of the application.
package main

import (
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/app"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/infra"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/pkg/time"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/pkg/uuid"
)

func main() {
	infraProviders := infra.NewInfraProviders()
	tp := time.NewTimeProvider()
	up := uuid.NewUUIDProvider()
	appServices := app.NewServices(infraProviders.CragRepository, infraProviders.NotificationService, up, tp)
	infraHTTPServer := infra.NewHTTPServer(appServices)
	infraHTTPServer.ListenAndServe(":8080")
}
