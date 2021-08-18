package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/pkritiotis/go-climb/internal/crag"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type MockAddCragHandler struct {
	Handler func(command crag.AddCragCommand) error
}

func (m MockAddCragHandler) Handle(command crag.AddCragCommand) error {
	return m.Handler(command)
}

func TestCragHandler_AddCrag(t *testing.T) {
	var tests = []struct {
		name               string
		handler            crag.AddCragCommandHandler
		reqVars            map[string]interface{}
		Body               interface{}
		ResultBodyContains string
		ResultStatus       int
	}{
		{
			name: "should add crag successfully",
			handler: MockAddCragHandler{Handler: func(command crag.AddCragCommand) error {
				if command.Country != "country" || command.Desc != "desc" || command.Name != "test" {
					return errors.New("objects not matching")
				}
				return nil
			}},
			reqVars: map[string]interface{}{},
			Body: AddCragRequestModel{
				Name:    "test",
				Desc:    "desc",
				Country: "country",
			},
			ResultBodyContains: "",
			ResultStatus:       http.StatusOK,
		},
		{
			name: "should return error",
			handler: MockAddCragHandler{Handler: func(command crag.AddCragCommand) error {
				if command.Country != "country" || command.Desc != "desc" || command.Name != "test" {
					return errors.New("objects not matching")
				}
				return errors.New("test error")
			}},
			reqVars: map[string]interface{}{},
			Body: AddCragRequestModel{
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
			c := NewCragHandler(crag.UseCases{Commands: crag.Commands{AddCragHandler: tt.handler}})
			buf := new(bytes.Buffer)
			_ = json.NewEncoder(buf).Encode(tt.Body)
			req, _ := http.NewRequest("POST", "", buf)
			rsp := httptest.NewRecorder()
			c.AddCrag(rsp, req)
			assert.Contains(t, tt.ResultBodyContains, rsp.Body.String())
			assert.Equal(t, tt.ResultStatus, rsp.Code)
		})
	}
}

type MockDeleteCragHandler struct {
	Handler func(command crag.DeleteCragCommand) error
}

func (m MockDeleteCragHandler) Handle(command crag.DeleteCragCommand) error {
	return m.Handler(command)
}
func TestCragHandler_DeleteCrag(t *testing.T) {
	var tests = []struct {
		name               string
		handler            crag.DeleteCragCommandHandler
		id                 string
		Body               interface{}
		ResultBodyContains string
		ResultStatus       int
	}{
		{
			name: "should delete crag successfully",
			handler: MockDeleteCragHandler{Handler: func(command crag.DeleteCragCommand) error {
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
			handler: MockDeleteCragHandler{Handler: func(command crag.DeleteCragCommand) error {
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
			c := NewCragHandler(crag.UseCases{Commands: crag.Commands{DeleteCragHandler: tt.handler}})
			buf := new(bytes.Buffer)
			if tt.Body != nil {
				_ = json.NewEncoder(buf).Encode(tt.Body)
			}
			req, _ := http.NewRequest("DELETE", "/crag/"+tt.id, buf)
			req = mux.SetURLVars(req, map[string]string{"cragId": tt.id})
			rsp := httptest.NewRecorder()
			c.DeleteCrag(rsp, req)
			assert.Contains(t, tt.ResultBodyContains, rsp.Body.String())
			assert.Equal(t, tt.ResultStatus, rsp.Code)
		})
	}
}

type MockGetCragsHandler struct {
	Handler func() ([]crag.QueryResult, error)
}

func (m MockGetCragsHandler) Handle() ([]crag.QueryResult, error) {
	return m.Handler()
}

func TestCragHandler_GetCrags(t *testing.T) {
	var tests = []struct {
		name               string
		handler            crag.GetAllCragsQueryHandler
		Body               interface{}
		ResultBodyContains string
		ResultStatus       int
	}{
		{
			name: "should get crags successfully",
			handler: MockGetCragsHandler{Handler: func() ([]crag.QueryResult, error) {
				return []crag.QueryResult{{ID: uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")}}, nil
			}},
			Body:               "",
			ResultBodyContains: "3e204a57-4449-4c74-8227-77934cf25322",
			ResultStatus:       http.StatusOK,
		},
		{
			name: "should return ok with empty body",
			handler: MockGetCragsHandler{Handler: func() ([]crag.QueryResult, error) {
				return nil, nil
			}},
			Body:               "",
			ResultBodyContains: "",
			ResultStatus:       http.StatusOK,
		},
		{
			name: "should return error",
			handler: MockGetCragsHandler{Handler: func() ([]crag.QueryResult, error) {
				return nil, errors.New("error")
			}},
			Body:               "",
			ResultBodyContains: "error",
			ResultStatus:       http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCragHandler(crag.UseCases{Queries: crag.Queries{GetAllCragsHandler: tt.handler}})
			buf := new(bytes.Buffer)
			_ = json.NewEncoder(buf).Encode(tt.Body)
			req, _ := http.NewRequest("GET", "", buf)
			rsp := httptest.NewRecorder()
			c.GetAllCrags(rsp, req)
			assert.Contains(t, rsp.Body.String(), tt.ResultBodyContains)
			assert.Equal(t, tt.ResultStatus, rsp.Code)
		})
	}
}

type MockGetCragHandler struct {
	Handler func(query crag.GetCragQuery) (*crag.QueryResult, error)
}

func (m MockGetCragHandler) Handle(query crag.GetCragQuery) (*crag.QueryResult, error) {
	return m.Handler(query)
}

func TestCragHandler_GetCrag(t *testing.T) {
	var tests = []struct {
		name               string
		handler            crag.GetCragQueryHandler
		id                 string
		Body               interface{}
		ResultBodyContains string
		ResultStatus       int
	}{
		{
			name: "should get crag successfully",
			handler: MockGetCragHandler{Handler: func(query crag.GetCragQuery) (*crag.QueryResult, error) {
				return &crag.QueryResult{ID: uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")}, nil
			}},
			id:                 "3e204a57-4449-4c74-8227-77934cf25322",
			Body:               "",
			ResultBodyContains: "3e204a57-4449-4c74-8227-77934cf25322",
			ResultStatus:       http.StatusOK,
		},
		{
			name: "should return not found",
			handler: MockGetCragHandler{Handler: func(query crag.GetCragQuery) (*crag.QueryResult, error) {
				return nil, nil
			}},
			id:                 "3e204a57-4449-4c74-8227-77934cf25322",
			Body:               "",
			ResultBodyContains: "Not Found",
			ResultStatus:       http.StatusNotFound,
		},
		{
			name: "should return error",
			handler: MockGetCragHandler{Handler: func(query crag.GetCragQuery) (*crag.QueryResult, error) {
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
			c := NewCragHandler(crag.UseCases{Queries: crag.Queries{GetCragHandler: tt.handler}})
			buf := new(bytes.Buffer)
			_ = json.NewEncoder(buf).Encode(tt.Body)
			req, _ := http.NewRequest("PUT", "", buf)
			req = mux.SetURLVars(req, map[string]string{"cragId": tt.id})
			rsp := httptest.NewRecorder()
			c.GetCrag(rsp, req)
			assert.Contains(t, rsp.Body.String(), tt.ResultBodyContains)
			assert.Equal(t, tt.ResultStatus, rsp.Code)
		})
	}
}

type MockUpdateCragHandler struct {
	Handler func(command crag.UpdateCragCommand) error
}

func (m MockUpdateCragHandler) Handle(command crag.UpdateCragCommand) error {
	return m.Handler(command)
}

func TestCragHandler_UpdateCrag(t *testing.T) {
	var tests = []struct {
		name               string
		handler            crag.UpdateCragCommandHandler
		id                 string
		Body               interface{}
		ResultBodyContains string
		ResultStatus       int
	}{
		{
			name: "should update crag successfully",
			handler: MockUpdateCragHandler{Handler: func(command crag.UpdateCragCommand) error {
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
			handler: MockUpdateCragHandler{Handler: func(command crag.UpdateCragCommand) error {
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
			handler: MockUpdateCragHandler{Handler: func(command crag.UpdateCragCommand) error {
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
			c := NewCragHandler(crag.UseCases{Commands: crag.Commands{UpdateCragHandler: tt.handler}})
			buf := new(bytes.Buffer)
			_ = json.NewEncoder(buf).Encode(tt.Body)
			req, _ := http.NewRequest("PUT", "", buf)
			req = mux.SetURLVars(req, map[string]string{"cragId": tt.id})
			rsp := httptest.NewRecorder()
			c.UpdateCrag(rsp, req)
			assert.Contains(t, tt.ResultBodyContains, rsp.Body.String())
			assert.Equal(t, tt.ResultStatus, rsp.Code)
		})
	}
}
