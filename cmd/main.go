package main

import (
	"github.com/pkritiotis/go-climb/internal/adapters/http"
	"github.com/pkritiotis/go-climb/internal/adapters/repo"
	"github.com/pkritiotis/go-climb/internal/app"
)

func main() {
	r := repo.NewInMemory()
	a := app.NewApp(r)
	httpServer := http.NewServer(a)
	httpServer.ListenAndServe(":8080")

}
