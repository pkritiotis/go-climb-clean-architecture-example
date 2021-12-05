package crag

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/app"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/app/crag/commands"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/app/crag/queries"
	"net/http"
)

//Handler Crag http request handler
type Handler struct {
	cragServices app.CragServices
}

//NewHandler Constructor
func NewHandler(app app.CragServices) *Handler {
	return &Handler{cragServices: app}
}

//GetAll Returns all available crags
func (c Handler) GetAll(w http.ResponseWriter, _ *http.Request) {
	crags, err := c.cragServices.Queries.GetAllCragsHandler.Handle()
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

// GetCragIDURLParam contains the parameter identifier to be parsed by the handler
const GetCragIDURLParam = "cragId"

//GetByID Returns the crag with the provided id
func (c Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cragID := vars[GetCragIDURLParam]
	crag, err := c.cragServices.Queries.GetCragHandler.Handle(queries.GetCragRequest{CragID: uuid.MustParse(cragID)})
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

//CreateCragRequestModel represents the request model expected for Add request
type CreateCragRequestModel struct {
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Country string `json:"country"`
}

//Create Adds the provides crag
func (c Handler) Create(w http.ResponseWriter, r *http.Request) {
	var cragToAdd CreateCragRequestModel
	decodeErr := json.NewDecoder(r.Body).Decode(&cragToAdd)
	if decodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, decodeErr.Error())
		return
	}
	err := c.cragServices.Commands.CreateCragHandler.Handle(commands.AddCragRequest{
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

// UpdateCragIDURLParam contains the parameter identifier to be parsed by the handler
const UpdateCragIDURLParam = "cragId"

//UpdateCragRequestModel represents the  request model of Update
type UpdateCragRequestModel struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Desc    string    `json:"desc"`
	Country string    `json:"country"`
}

//Update Updates crag with the provided data
func (c Handler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cragID := uuid.MustParse(vars[UpdateCragIDURLParam])

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

	err := c.cragServices.Commands.UpdateCragHandler.Handle(cragToUpdateCommand)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteCragIDURLParam contains the parameter identifier to be parsed by the handler
const DeleteCragIDURLParam = "cragId"

//Delete Deletes the crag with the provided id
func (c Handler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cragID := vars[DeleteCragIDURLParam]
	err := c.cragServices.Commands.DeleteCragHandler.Handle(commands.DeleteCragRequest{CragID: uuid.MustParse(cragID)})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
	}
}
