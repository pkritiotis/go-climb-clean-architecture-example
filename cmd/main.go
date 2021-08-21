package main

import (
	"github.com/pkritiotis/go-climb/internal/app"
	"github.com/pkritiotis/go-climb/internal/infra/http"
	"github.com/pkritiotis/go-climb/internal/infra/notification"
	"github.com/pkritiotis/go-climb/internal/infra/repo"
	"github.com/pkritiotis/go-climb/internal/pkg/time"
	"github.com/pkritiotis/go-climb/internal/pkg/uuid"
)

func main() {
	r := repo.NewInMemory()
	ns := notification.NewConsoleService()
	tp := time.NewTimeProvider()
	up := uuid.NewUUIDProvider()
	a := app.NewApp(r, ns, up, tp)
	httpServer := http.NewServer(a)
	httpServer.ListenAndServe(":8080")
}
