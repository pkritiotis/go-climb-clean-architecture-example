package http

import (
	"github.com/gorilla/mux"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/app"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/inputports/http/crag"
	"log"
	"net/http"
)

//Server Represents the http server running for this service
type Server struct {
	appServices app.Services
	router      *mux.Router
}

//NewServer HTTP Server constructor
func NewServer(appServices app.Services) *Server {
	httpServer := &Server{appServices: appServices}
	httpServer.router = mux.NewRouter()
	httpServer.AddCragHTTPRoutes()
	http.Handle("/", httpServer.router)

	return httpServer
}

// AddCragHTTPRoutes registers crag route handlers
func (httpServer *Server) AddCragHTTPRoutes() {
	const cragsHTTPRoutePath = "/crags"
	//Queries
	httpServer.router.HandleFunc(cragsHTTPRoutePath, crag.NewHandler(httpServer.appServices.CragServices).GetAll).Methods("GET")
	httpServer.router.HandleFunc(cragsHTTPRoutePath+"/{"+crag.GetCragIDURLParam+"}", crag.NewHandler(httpServer.appServices.CragServices).GetByID).Methods("GET")

	//Commands
	httpServer.router.HandleFunc(cragsHTTPRoutePath, crag.NewHandler(httpServer.appServices.CragServices).Create).Methods("POST")
	httpServer.router.HandleFunc(cragsHTTPRoutePath+"/{"+crag.UpdateCragIDURLParam+"}", crag.NewHandler(httpServer.appServices.CragServices).Update).Methods("PUT")
	httpServer.router.HandleFunc(cragsHTTPRoutePath+"/{"+crag.DeleteCragIDURLParam+"}", crag.NewHandler(httpServer.appServices.CragServices).Delete).Methods("DELETE")

}

//ListenAndServe Starts listening for requests
func (httpServer *Server) ListenAndServe(port string) {
	log.Fatal(http.ListenAndServe(port, nil))
}
