package inputports

import (
	"github.com/pkritiotis/go-climb/internal/app"
	"github.com/pkritiotis/go-climb/internal/inputports/http"
)

type Services struct {
	Server *http.Server
}

func NewServices(appServices app.Services) Services {
	return Services{Server: http.NewServer(appServices)}
}
