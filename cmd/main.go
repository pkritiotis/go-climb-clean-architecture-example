package main

import (
	"github.com/pkritiotis/go-clean/internal/adapters/repo"
	"github.com/pkritiotis/go-clean/internal/app"
	"github.com/pkritiotis/go-clean/internal/ports"
)

func main() {
	r := repo.NewInMemory()
	a := app.NewApp(r)
	httpServer := ports.NewHTTPServer(&a)
	httpServer.ListenAndServe(":8080")

}
