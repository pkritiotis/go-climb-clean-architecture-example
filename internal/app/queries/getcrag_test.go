package queries

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean/internal/app/services"
	"github.com/pkritiotis/go-clean/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetCragQueryHandler_Handle(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")
	mockCrag := &domain.Crag{
		ID:        mockUUID,
		Name:      "test",
		Desc:      "test",
		Country:   "test",
		CreatedAt: time.Time{},
	}
	type fields struct {
		repo services.CragRepository
	}
	type args struct {
		query GetCragQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *domain.Crag
		err    error
	}{
		{
			name: "happy path - no errors - return crag",
			fields: fields{
				repo: func() services.MockRepository {
					mp := services.MockRepository{}
					mp.On("GetCrag", mockUUID).Return(mockCrag, nil)

					return mp
				}(),
			},
			args: args{
				query: GetCragQuery{
					CragID: mockUUID,
				},
			},
			want: mockCrag,
			err:  nil,
		},
		{
			name: "no crag - no errors - return nil",
			fields: fields{
				repo: func() services.MockRepository {
					mp := services.MockRepository{}
					mp.On("GetCrag", mockUUID).Return((*domain.Crag)(nil), nil)

					return mp
				}(),
			},
			args: args{
				query: GetCragQuery{
					CragID: mockUUID,
				},
			},
			want: (*domain.Crag)(nil),
			err:  nil,
		},
		{
			name: "get crag error - return nil",
			fields: fields{
				repo: func() services.MockRepository {
					mp := services.MockRepository{}
					mp.On("GetCrag", mockUUID).Return((*domain.Crag)(nil), errors.New("get error"))

					return mp
				}(),
			},
			args: args{
				query: GetCragQuery{
					CragID: mockUUID,
				},
			},
			want: (*domain.Crag)(nil),
			err:  errors.New("get error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := GetCragQueryHandler{
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
		repo services.CragRepository
	}
	tests := []struct {
		name string
		args args
		want GetCragQueryHandler
	}{
		{
			name: "construct handler",
			args: args{
				repo: services.MockRepository{},
			},
			want: GetCragQueryHandler{
				repo: services.MockRepository{},
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