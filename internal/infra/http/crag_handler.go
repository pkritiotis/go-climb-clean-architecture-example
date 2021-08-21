package http

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkritiotis/go-climb/internal/app"
	"github.com/pkritiotis/go-climb/internal/app/crag/commands"
	"github.com/pkritiotis/go-climb/internal/app/crag/queries"
	"net/http"
)

//CragHandler Crag http request handler
type CragHandler struct {
	app app.App
}

//NewCragHandler Constructor
func NewCragHandler(app app.App) *CragHandler {
	return &CragHandler{app: app}
}

//GetAllCrags Returns all available crags
func (c CragHandler) GetAllCrags(w http.ResponseWriter, _ *http.Request) {
	crags, err := c.app.Queries.GetAllCragsHandler.Handle()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(crags)
	if err != nil {
		return
	}
}

const getCragIDURLParam = "cragId"

//GetCrag Returns the crag with the provided id
func (c CragHandler) GetCrag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cragID := vars[getCragIDURLParam]
	crag, err := c.app.Queries.GetCragHandler.Handle(queries.GetCragRequest{CragID: uuid.MustParse(cragID)})
	if err == nil && crag == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not Found")
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(crag)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
}

//AddCragRequestModel represents the request model expected for Add request
type AddCragRequestModel struct {
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Country string `json:"country"`
}

//AddCrag Adds the provides crag
func (c CragHandler) AddCrag(w http.ResponseWriter, r *http.Request) {
	var cragToAdd AddCragRequestModel
	decodeErr := json.NewDecoder(r.Body).Decode(&cragToAdd)
	if decodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, decodeErr.Error())
		return
	}
	err := c.app.Commands.AddCragHandler.Handle(commands.AddCragRequest{
		Name:    cragToAdd.Name,
		Desc:    cragToAdd.Desc,
		Country: cragToAdd.Country,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

const updateCragIDURLParam = "cragId"

//UpdateCragRequestModel represents the  request model of Update
type UpdateCragRequestModel struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Desc    string    `json:"desc"`
	Country string    `json:"country"`
}

//UpdateCrag Updates path with the provided data
func (c CragHandler) UpdateCrag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cragID := uuid.MustParse(vars[updateCragIDURLParam])

	var reqCragToUpdate UpdateCragRequestModel
	decodeErr := json.NewDecoder(r.Body).Decode(&reqCragToUpdate)
	if decodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, decodeErr)
		return
	}

	if cragID != reqCragToUpdate.ID {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprint(w, "Inconsistency between route id and body id")
		return
	}
	cragToUpdateCommand := commands.UpdateCragRequest{
		ID:      reqCragToUpdate.ID,
		Name:    reqCragToUpdate.Name,
		Desc:    reqCragToUpdate.Desc,
		Country: reqCragToUpdate.Country,
	}

	err := c.app.Commands.UpdateCragHandler.Handle(cragToUpdateCommand)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}
	w.WriteHeader(http.StatusOK)
}

const deleteCragIDURLParam = "cragId"

//DeleteCrag Deletes the crag with the provided id
func (c CragHandler) DeleteCrag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cragID := vars[deleteCragIDURLParam]
	err := c.app.Commands.DeleteCragHandler.Handle(commands.DeleteCragRequest{CragID: uuid.MustParse(cragID)})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
	}
}
