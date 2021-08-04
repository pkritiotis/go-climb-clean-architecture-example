package ports

import (
	"github.com/gorilla/mux"
	"github.com/pkritiotis/go-clean/internal/app"
	"github.com/pkritiotis/go-clean/internal/ports/crag"
	"log"
	"net/http"
)

//HTTPServer Represents the http server running for this service
type HTTPServer struct {
	app    app.App
	router *mux.Router
}

//NewHTTPServer HTTPServer constructor
func NewHTTPServer(app app.App) *HTTPServer {
	httpServer := &HTTPServer{app: app}
	httpServer.router = mux.NewRouter()
	httpServer.AddCragHTTPRoutes()
	http.Handle("/", httpServer.router)

	return &HTTPServer{app: app}
}

//AddCragHTTPRoutes Adds all available Routes to the router
func (httpServer *HTTPServer) AddCragHTTPRoutes() {
	//Queries
	httpServer.router.HandleFunc(crag.GetAllCragsRoutePath, crag.NewHTTPHandler(httpServer.app).GetAllCrags).Methods("GET")
	httpServer.router.HandleFunc(crag.GetCragRoutePath, crag.NewHTTPHandler(httpServer.app).GetCrag).Methods("GET")

	//Commands
	httpServer.router.HandleFunc(crag.AddCragRoutePath, crag.NewHTTPHandler(httpServer.app).AddCrag).Methods("POST")
	httpServer.router.HandleFunc(crag.UpdateCragRoutePath, crag.NewHTTPHandler(httpServer.app).UpdateCrag).Methods("PUT")
	httpServer.router.HandleFunc(crag.DeleteCragIDRoutePath, crag.NewHTTPHandler(httpServer.app).DeleteCrag).Methods("DELETE")

}

//ListenAndServe Starts listening for requests
func (httpServer *HTTPServer) ListenAndServe(port string) {
	log.Fatal(http.ListenAndServe(port, nil))
}
