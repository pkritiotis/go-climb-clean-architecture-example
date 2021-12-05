package crag

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/app/crag/commands"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/app/crag/queries"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/app"
	"github.com/stretchr/testify/assert"
)

type MockAddCragHandler struct {
	Handler func(command commands.AddCragRequest) error
}

func (m MockAddCragHandler) Handle(command commands.AddCragRequest) error {
	return m.Handler(command)
}

func TestCragHandler_AddCrag(t *testing.T) {
	var tests = []struct {
		name               string
		handler            commands.CreateCragRequestHandler
		reqVars            map[string]interface{}
		Body               interface{}
		ResultBodyContains string
		ResultStatus       int
	}{
		{
			name: "should add crag successfully",
			handler: MockAddCragHandler{Handler: func(command commands.AddCragRequest) error {
				if command.Country != "country" || command.Desc != "desc" || command.Name != "test" {
					return errors.New("objects not matching")
				}
				return nil
			}},
			reqVars: map[string]interface{}{},
			Body: CreateCragRequestModel{
				Name:    "test",
				Desc:    "desc",
				Country: "country",
			},
			ResultBodyContains: "",
			ResultStatus:       http.StatusOK,
		},
		{
			name: "should return error",
			handler: MockAddCragHandler{Handler: func(command commands.AddCragRequest) error {
				if command.Country != "country" || command.Desc != "desc" || command.Name != "test" {
					return errors.New("objects not matching")
				}
				return errors.New("test error")
			}},
			reqVars: map[string]interface{}{},
			Body: CreateCragRequestModel{
				Name:    "test",
				Desc:    "desc",
				Country: "country",
			},
			ResultBodyContains: errors.New("test error").Error(),
			ResultStatus:       http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewHandler(app.CragServices{Commands: app.Commands{CreateCragHandler: tt.handler}})
			buf := new(bytes.Buffer)
			_ = json.NewEncoder(buf).Encode(tt.Body)
			req, _ := http.NewRequest("POST", "", buf)
			rsp := httptest.NewRecorder()
			c.Create(rsp, req)
			assert.Contains(t, tt.ResultBodyContains, rsp.Body.String())
			assert.Equal(t, tt.ResultStatus, rsp.Code)
		})
	}
}

type MockDeleteCragHandler struct {
	Handler func(command commands.DeleteCragRequest) error
}

func (m MockDeleteCragHandler) Handle(command commands.DeleteCragRequest) error {
	return m.Handler(command)
}
func TestCragHandler_DeleteCrag(t *testing.T) {
	var tests = []struct {
		name               string
		handler            commands.DeleteCragRequestHandler
		id                 string
		Body               interface{}
		ResultBodyContains string
		ResultStatus       int
	}{
		{
			name: "should delete crag successfully",
			handler: MockDeleteCragHandler{Handler: func(command commands.DeleteCragRequest) error {
				if command.CragID != uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322") {
					return errors.New("objects not matching")
				}
				return nil
			}},
			id:                 "3e204a57-4449-4c74-8227-77934cf25322",
			Body:               nil,
			ResultBodyContains: "",
			ResultStatus:       http.StatusOK,
		},
		{
			name: "should return error",
			handler: MockDeleteCragHandler{Handler: func(command commands.DeleteCragRequest) error {
				if command.CragID != uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322") {
					return errors.New("objects not matching")
				}
				return errors.New("test error")
			}},
			id:                 "3e204a57-4449-4c74-8227-77934cf25322",
			Body:               nil,
			ResultBodyContains: errors.New("test error").Error(),
			ResultStatus:       http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewHandler(app.CragServices{Commands: app.Commands{DeleteCragHandler: tt.handler}})
			buf := new(bytes.Buffer)
			if tt.Body != nil {
				_ = json.NewEncoder(buf).Encode(tt.Body)
			}
			req, _ := http.NewRequest("DELETE", "/crag/"+tt.id, buf)
			req = mux.SetURLVars(req, map[string]string{"cragId": tt.id})
			rsp := httptest.NewRecorder()
			c.Delete(rsp, req)
			assert.Contains(t, tt.ResultBodyContains, rsp.Body.String())
			assert.Equal(t, tt.ResultStatus, rsp.Code)
		})
	}
}

type MockGetCragsHandler struct {
	Handler func() ([]queries.GetAllCragsResult, error)
}

func (m MockGetCragsHandler) Handle() ([]queries.GetAllCragsResult, error) {
	return m.Handler()
}

func TestCragHandler_GetCrags(t *testing.T) {
	var tests = []struct {
		name               string
		handler            queries.GetAllCragsRequestHandler
		Body               interface{}
		ResultBodyContains string
		ResultStatus       int
	}{
		{
			name: "should get crags successfully",
			handler: MockGetCragsHandler{Handler: func() ([]queries.GetAllCragsResult, error) {
				return []queries.GetAllCragsResult{{ID: uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")}}, nil
			}},
			Body:               "",
			ResultBodyContains: "3e204a57-4449-4c74-8227-77934cf25322",
			ResultStatus:       http.StatusOK,
		},
		{
			name: "should return ok with empty body",
			handler: MockGetCragsHandler{Handler: func() ([]queries.GetAllCragsResult, error) {
				return nil, nil
			}},
			Body:               "",
			ResultBodyContains: "",
			ResultStatus:       http.StatusOK,
		},
		{
			name: "should return error",
			handler: MockGetCragsHandler{Handler: func() ([]queries.GetAllCragsResult, error) {
				return nil, errors.New("error")
			}},
			Body:               "",
			ResultBodyContains: "error",
			ResultStatus:       http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewHandler(app.CragServices{Queries: app.Queries{GetAllCragsHandler: tt.handler}})
			buf := new(bytes.Buffer)
			_ = json.NewEncoder(buf).Encode(tt.Body)
			req, _ := http.NewRequest("GET", "", buf)
			rsp := httptest.NewRecorder()
			c.GetAll(rsp, req)
			assert.Contains(t, rsp.Body.String(), tt.ResultBodyContains)
			assert.Equal(t, tt.ResultStatus, rsp.Code)
		})
	}
}

type MockGetCragHandler struct {
	Handler func(query queries.GetCragRequest) (*queries.GetCragResult, error)
}

func (m MockGetCragHandler) Handle(query queries.GetCragRequest) (*queries.GetCragResult, error) {
	return m.Handler(query)
}

func TestCragHandler_GetCrag(t *testing.T) {
	var tests = []struct {
		name               string
		handler            queries.GetCragRequestHandler
		id                 string
		Body               interface{}
		ResultBodyContains string
		ResultStatus       int
	}{
		{
			name: "should get crag successfully",
			handler: MockGetCragHandler{Handler: func(query queries.GetCragRequest) (*queries.GetCragResult, error) {
				return &queries.GetCragResult{ID: uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")}, nil
			}},
			id:                 "3e204a57-4449-4c74-8227-77934cf25322",
			Body:               "",
			ResultBodyContains: "3e204a57-4449-4c74-8227-77934cf25322",
			ResultStatus:       http.StatusOK,
		},
		{
			name: "should return not found",
			handler: MockGetCragHandler{Handler: func(query queries.GetCragRequest) (*queries.GetCragResult, error) {
				return nil, nil
			}},
			id:                 "3e204a57-4449-4c74-8227-77934cf25322",
			Body:               "",
			ResultBodyContains: "Not Found",
			ResultStatus:       http.StatusNotFound,
		},
		{
			name: "should return error",
			handler: MockGetCragHandler{Handler: func(query queries.GetCragRequest) (*queries.GetCragResult, error) {
				return nil, errors.New("error")
			}},
			id:                 "3e204a57-4449-4c74-8227-77934cf25322",
			Body:               "",
			ResultBodyContains: "error",
			ResultStatus:       http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewHandler(app.CragServices{Queries: app.Queries{GetCragHandler: tt.handler}})
			buf := new(bytes.Buffer)
			_ = json.NewEncoder(buf).Encode(tt.Body)
			req, _ := http.NewRequest("PUT", "", buf)
			req = mux.SetURLVars(req, map[string]string{"cragId": tt.id})
			rsp := httptest.NewRecorder()
			c.GetByID(rsp, req)
			assert.Contains(t, rsp.Body.String(), tt.ResultBodyContains)
			assert.Equal(t, tt.ResultStatus, rsp.Code)
		})
	}
}

type MockUpdateCragHandler struct {
	Handler func(command commands.UpdateCragRequest) error
}

func (m MockUpdateCragHandler) Handle(command commands.UpdateCragRequest) error {
	return m.Handler(command)
}

func TestCragHandler_UpdateCrag(t *testing.T) {
	var tests = []struct {
		name               string
		handler            commands.UpdateCragRequestHandler
		id                 string
		Body               interface{}
		ResultBodyContains string
		ResultStatus       int
	}{
		{
			name: "should update crag successfully",
			handler: MockUpdateCragHandler{Handler: func(command commands.UpdateCragRequest) error {
				if command.Country != "country" || command.Desc != "desc" || command.Name != "test" || command.ID != uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322") {
					return errors.New("objects not matching")
				}
				return nil
			}},
			id: "3e204a57-4449-4c74-8227-77934cf25322",
			Body: UpdateCragRequestModel{

				ID:      uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322"),
				Name:    "test",
				Desc:    "desc",
				Country: "country",
			},
			ResultBodyContains: "",
			ResultStatus:       http.StatusOK,
		},
		{
			name: "inconsistent url - body ids - should return conflict",
			handler: MockUpdateCragHandler{Handler: func(command commands.UpdateCragRequest) error {
				if command.Country != "country" || command.Desc != "desc" || command.Name != "test" || command.ID != uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322") {
					return errors.New("objects not matching")
				}
				return nil
			}},
			id: "3e204a57-4449-4c74-8227-77934cf25322",
			Body: UpdateCragRequestModel{
				ID:      uuid.MustParse("4e204a57-4449-4c74-8227-77934cf25322"),
				Name:    "test",
				Desc:    "desc",
				Country: "country",
			},
			ResultBodyContains: "Inconsistency between route id and body id",
			ResultStatus:       http.StatusConflict,
		},
		{
			name: "should return internal server error",
			handler: MockUpdateCragHandler{Handler: func(command commands.UpdateCragRequest) error {
				if command.Country != "country" || command.Desc != "desc" || command.Name != "test" || command.ID != uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322") {
					return errors.New("objects not matching")
				}
				return errors.New("error")
			}},
			id: "3e204a57-4449-4c74-8227-77934cf25322",
			Body: UpdateCragRequestModel{
				ID:      uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322"),
				Name:    "test",
				Desc:    "desc",
				Country: "country",
			},
			ResultBodyContains: errors.New("error").Error(),
			ResultStatus:       http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewHandler(app.CragServices{Commands: app.Commands{UpdateCragHandler: tt.handler}})
			buf := new(bytes.Buffer)
			_ = json.NewEncoder(buf).Encode(tt.Body)
			req, _ := http.NewRequest("PUT", "", buf)
			req = mux.SetURLVars(req, map[string]string{"cragId": tt.id})
			rsp := httptest.NewRecorder()
			c.Update(rsp, req)
			assert.Contains(t, tt.ResultBodyContains, rsp.Body.String())
			assert.Equal(t, tt.ResultStatus, rsp.Code)
		})
	}
}
