package commands

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/domain/crag"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteCragCommandHandler_Handle(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")

	type fields struct {
		repo crag.Repository
	}
	type args struct {
		command DeleteCragRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		err    error
	}{
		{
			name: "happy path - no errors - should return nil",
			fields: fields{
				repo: func() crag.MockRepository {
					mp := crag.MockRepository{}
					mp.On("GetByID", mockUUID).Return(&crag.Crag{ID: mockUUID}, nil)
					mp.On("Delete", mockUUID).Return(nil)
					return mp
				}(),
			},
			args: args{
				command: DeleteCragRequest{
					CragID: mockUUID,
				},
			},
			err: nil,
		},
		{
			name: "get crag returns error - should return error",
			fields: fields{
				repo: func() crag.MockRepository {
					mp := crag.MockRepository{}
					mp.On("GetByID", mockUUID).Return(&crag.Crag{ID: mockUUID}, errors.New("get error"))
					return mp
				}(),
			},
			args: args{
				command: DeleteCragRequest{
					CragID: mockUUID,
				},
			},
			err: errors.New("get error"),
		},
		{
			name: "get crag returns nil - should return error",
			fields: fields{
				repo: func() crag.MockRepository {
					mp := crag.MockRepository{}
					mp.On("GetByID", mockUUID).Return((*crag.Crag)(nil), nil)
					return mp
				}(),
			},
			args: args{
				command: DeleteCragRequest{
					CragID: mockUUID,
				},
			},
			err: fmt.Errorf("the provided crag id does not exist"),
		},
		{
			name: "delete crag returns error - should return error",
			fields: fields{
				repo: func() crag.MockRepository {
					mp := crag.MockRepository{}
					mp.On("GetByID", mockUUID).Return(&crag.Crag{ID: mockUUID}, nil)
					mp.On("Delete", mockUUID).Return(errors.New("delete error"))
					return mp
				}(),
			},
			args: args{
				command: DeleteCragRequest{
					CragID: mockUUID,
				},
			},
			err: errors.New("delete error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := deleteCragRequestHandler{
				repo: tt.fields.repo,
			}
			err := h.Handle(tt.args.command)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestNewDeleteCragCommandHandler(t *testing.T) {
	type args struct {
		repo crag.Repository
	}
	tests := []struct {
		name string
		args args
		want DeleteCragRequestHandler
	}{
		{
			name: "should return delete request handler",
			args: args{
				repo: crag.MockRepository{},
			},
			want: deleteCragRequestHandler{
				repo: crag.MockRepository{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDeleteCragRequestHandler(tt.args.repo)
			assert.Equal(t, tt.want, got)
		})
	}
}
