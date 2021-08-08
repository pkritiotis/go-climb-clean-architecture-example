package http

import (
	"github.com/gorilla/mux"
	"github.com/pkritiotis/go-clean/internal/app"
	"log"
	"net/http"
)

//Server Represents the http server running for this service
type Server struct {
	app    app.App
	router *mux.Router
}

//NewServer HTTP Server constructor
func NewServer(app app.App) *Server {
	httpServer := &Server{app: app}
	httpServer.router = mux.NewRouter()
	httpServer.AddCragHTTPRoutes()
	http.Handle("/", httpServer.router)

	return &Server{app: app}
}

//AddCragHTTPRoutes Adds all available Routes to the router
func (httpServer *Server) AddCragHTTPRoutes() {
	//Queries
	httpServer.router.HandleFunc(GetAllCragsRoutePath, NewCragHandler(httpServer.app).GetAllCrags).Methods("GET")
	httpServer.router.HandleFunc(GetCragRoutePath, NewCragHandler(httpServer.app).GetCrag).Methods("GET")

	//Commands
	httpServer.router.HandleFunc(AddCragRoutePath, NewCragHandler(httpServer.app).AddCrag).Methods("POST")
	httpServer.router.HandleFunc(UpdateCragRoutePath, NewCragHandler(httpServer.app).UpdateCrag).Methods("PUT")
	httpServer.router.HandleFunc(DeleteCragIDRoutePath, NewCragHandler(httpServer.app).DeleteCrag).Methods("DELETE")

}

//ListenAndServe Starts listening for requests
func (httpServer *Server) ListenAndServe(port string) {
	log.Fatal(http.ListenAndServe(port, nil))
}
