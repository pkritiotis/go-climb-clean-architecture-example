package main

import (
	"github.com/pkritiotis/go-climb/internal/adapters/http"
	"github.com/pkritiotis/go-climb/internal/adapters/notification"
	"github.com/pkritiotis/go-climb/internal/adapters/repo"
	"github.com/pkritiotis/go-climb/internal/app"
)

func main() {
	r := repo.NewInMemory()
	ns := notification.ConsoleNotificationService{}
	a := app.NewApp(r, ns)
	httpServer := http.NewServer(a)
	httpServer.ListenAndServe(":8080")
}
