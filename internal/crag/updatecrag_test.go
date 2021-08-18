package crag

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewUpdateCragCommandHandler(t *testing.T) {
	type args struct {
		repo Repository
	}
	tests := []struct {
		name string
		args args
		want UpdateCragCommandHandler
	}{
		{
			name: "should construct handler",
			args: args{
				repo: MockRepository{},
			},
			want: updateCragCommandHandler{
				repo: MockRepository{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUpdateCragCommandHandler(tt.args.repo)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUpdateCragCommandHandler_Handle(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")

	type fields struct {
		repo Repository
	}
	type args struct {
		command UpdateCragCommand
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
					returnedCrag := Crag{
						ID:        mockUUID,
						Name:      "initial",
						Desc:      "initial",
						Country:   "initial",
						CreatedAt: time.Time{},
					}
					updatedCrag := Crag{
						ID:        mockUUID,
						Name:      "updated",
						Desc:      "updated",
						Country:   "updated",
						CreatedAt: time.Time{},
					}
					mp.On("GetByID", mockUUID).Return(&returnedCrag, nil)
					mp.On("Update", updatedCrag).Return(nil)

					return mp
				}(),
			},
			args: args{
				command: UpdateCragCommand{
					ID:      mockUUID,
					Name:    "updated",
					Desc:    "updated",
					Country: "updated",
				},
			},
			err: nil,
		},
		{
			name: "get error should return error",
			fields: fields{
				repo: func() MockRepository {
					mp := MockRepository{}
					mp.On("GetByID", mockUUID).Return(&Crag{ID: mockUUID}, errors.New("get error"))

					return mp
				}(),
			},
			args: args{
				command: UpdateCragCommand{
					ID:      mockUUID,
					Name:    "updated",
					Desc:    "updated",
					Country: "updated",
				},
			},
			err: errors.New("get error"),
		},
		{
			name: "get returns nil, should return error",
			fields: fields{
				repo: func() MockRepository {
					mp := MockRepository{}
					mp.On("GetByID", mockUUID).Return((*Crag)(nil), nil)
					return mp
				}(),
			},
			args: args{
				command: UpdateCragCommand{
					ID:      mockUUID,
					Name:    "updated",
					Desc:    "updated",
					Country: "updated",
				},
			},
			err: fmt.Errorf("the provided crag id does not exist"),
		},
		{
			name: "update error - should return error",
			fields: fields{
				repo: func() MockRepository {
					mp := MockRepository{}
					returnedCrag := Crag{
						ID:        mockUUID,
						Name:      "initial",
						Desc:      "initial",
						Country:   "initial",
						CreatedAt: time.Time{},
					}
					updatedCrag := Crag{
						ID:        mockUUID,
						Name:      "updated",
						Desc:      "updated",
						Country:   "updated",
						CreatedAt: time.Time{},
					}
					mp.On("GetByID", mockUUID).Return(&returnedCrag, nil)
					mp.On("Update", updatedCrag).Return(errors.New("update error"))

					return mp
				}(),
			},
			args: args{
				command: UpdateCragCommand{
					ID:      mockUUID,
					Name:    "updated",
					Desc:    "updated",
					Country: "updated",
				},
			},
			err: errors.New("update error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := updateCragCommandHandler{
				repo: tt.fields.repo,
			}
			err := h.Handle(tt.args.command)
			assert.Equal(t, tt.err, err)
		})
	}
}