package queries

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-climb/internal/domain/crag"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetCragQueryHandler_Handle(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")
	mockCrag := &crag.Crag{
		ID:        mockUUID,
		Name:      "test",
		Desc:      "test",
		Country:   "test",
		CreatedAt: time.Time{},
	}

	cragQueryResult := &CragQueryResult{
		ID:        mockUUID,
		Name:      mockCrag.Name,
		Desc:      mockCrag.Desc,
		Country:   mockCrag.Country,
		CreatedAt: mockCrag.CreatedAt,
	}
	type fields struct {
		repo crag.Repository
	}
	type args struct {
		query GetCragQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CragQueryResult
		err    error
	}{
		{
			name: "happy path - no errors - return crag",
			fields: fields{
				repo: func() crag.MockRepository {
					mp := crag.MockRepository{}
					mp.On("GetByID", mockUUID).Return(mockCrag, nil)

					return mp
				}(),
			},
			args: args{
				query: GetCragQuery{
					CragID: mockUUID,
				},
			},
			want: cragQueryResult,
			err:  nil,
		},
		{
			name: "no crag - no errors - return nil",
			fields: fields{
				repo: func() crag.MockRepository {
					mp := crag.MockRepository{}
					mp.On("GetByID", mockUUID).Return((*crag.Crag)(nil), nil)

					return mp
				}(),
			},
			args: args{
				query: GetCragQuery{
					CragID: mockUUID,
				},
			},
			want: (*CragQueryResult)(nil),
			err:  nil,
		},
		{
			name: "get crag error - return nil",
			fields: fields{
				repo: func() crag.MockRepository {
					mp := crag.MockRepository{}
					mp.On("GetByID", mockUUID).Return((*crag.Crag)(nil), errors.New("get error"))

					return mp
				}(),
			},
			args: args{
				query: GetCragQuery{
					CragID: mockUUID,
				},
			},
			want: (*CragQueryResult)(nil),
			err:  errors.New("get error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := getCragQueryHandler{
				repo: tt.fields.repo,
			}
			got, err := h.Handle(tt.args.query)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestNewGetCragQueryHandler(t *testing.T) {
	type args struct {
		repo crag.Repository
	}
	tests := []struct {
		name string
		args args
		want GetCragQueryHandler
	}{
		{
			name: "construct handler",
			args: args{
				repo: crag.MockRepository{},
			},
			want: getCragQueryHandler{
				repo: crag.MockRepository{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGetCragQueryHandler(tt.args.repo)
			assert.Equal(t, tt.want, got)
		})
	}
}
