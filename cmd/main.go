package main

import (
	"github.com/pkritiotis/go-climb/internal/app"
	"github.com/pkritiotis/go-climb/internal/infra/http"
	"github.com/pkritiotis/go-climb/internal/infra/notification"
	"github.com/pkritiotis/go-climb/internal/infra/repo"
)

func main() {
	r := repo.NewInMemory()
	ns := notification.ConsoleNotificationService{}
	a := app.NewApp(r, ns)
	httpServer := http.NewServer(a)
	httpServer.ListenAndServe(":8080")
}
