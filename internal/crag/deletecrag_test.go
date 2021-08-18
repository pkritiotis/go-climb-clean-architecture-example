package crag

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteCragCommandHandler_Handle(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")

	type fields struct {
		repo Repository
	}
	type args struct {
		command DeleteCragCommand
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
				repo: func() MockRepository {
					mp := MockRepository{}
					mp.On("GetByID", mockUUID).Return(&Crag{ID: mockUUID}, nil)
					mp.On("Delete", mockUUID).Return(nil)
					return mp
				}(),
			},
			args: args{
				command: DeleteCragCommand{
					CragID: mockUUID,
				},
			},
			err: nil,
		},
		{
			name: "get crag returns error - should return error",
			fields: fields{
				repo: func() MockRepository {
					mp := MockRepository{}
					mp.On("GetByID", mockUUID).Return(&Crag{ID: mockUUID}, errors.New("get error"))
					return mp
				}(),
			},
			args: args{
				command: DeleteCragCommand{
					CragID: mockUUID,
				},
			},
			err: errors.New("get error"),
		},
		{
			name: "get crag returns nil - should return error",
			fields: fields{
				repo: func() MockRepository {
					mp := MockRepository{}
					mp.On("GetByID", mockUUID).Return((*Crag)(nil), nil)
					return mp
				}(),
			},
			args: args{
				command: DeleteCragCommand{
					CragID: mockUUID,
				},
			},
			err: fmt.Errorf("the provided crag id does not exist"),
		},
		{
			name: "delete crag returns error - should return error",
			fields: fields{
				repo: func() MockRepository {
					mp := MockRepository{}
					mp.On("GetByID", mockUUID).Return(&Crag{ID: mockUUID}, nil)
					mp.On("Delete", mockUUID).Return(errors.New("delete error"))
					return mp
				}(),
			},
			args: args{
				command: DeleteCragCommand{
					CragID: mockUUID,
				},
			},
			err: errors.New("delete error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := deleteCragCommandHandler{
				repo: tt.fields.repo,
			}
			err := h.Handle(tt.args.command)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestNewDeleteCragCommandHandler(t *testing.T) {
	type args struct {
		repo Repository
	}
	tests := []struct {
		name string
		args args
		want DeleteCragCommandHandler
	}{
		{
			name: "should return delete command handler",
			args: args{
				repo: MockRepository{},
			},
			want: deleteCragCommandHandler{
				repo: MockRepository{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDeleteCragCommandHandler(tt.args.repo)
			assert.Equal(t, tt.want, got)
		})
	}
}
