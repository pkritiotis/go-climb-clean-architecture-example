package main

import (
	"github.com/pkritiotis/go-climb/internal/crag"
	"github.com/pkritiotis/go-climb/internal/http"
	"github.com/pkritiotis/go-climb/internal/notification"
	"github.com/pkritiotis/go-climb/internal/storage"
	"github.com/pkritiotis/go-climb/internal/timeutil"
	"github.com/pkritiotis/go-climb/internal/uuidutil"
)

func main() {
	r := storage.NewInMemory()
	ns := notification.ConsoleNotificationService{}
	up := uuidutil.NewUUIDProvider()
	tp := timeutil.NewTimeProvider()
	useCases := crag.NewUseCases(r, ns, up, tp)
	httpServer := http.NewServer(useCases)
	httpServer.ListenAndServe(":8080")
}
