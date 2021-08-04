package commands

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean/internal/app/common"
	"github.com/pkritiotis/go-clean/internal/app/services"
	"github.com/pkritiotis/go-clean/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAddCragCommandHandler_Handle(t *testing.T) {
	mockTime, _ := time.Parse("yyyy-MM-02", "2021-07-30")
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")
	type fields struct {
		uuidProvider common.UUIDProvider
		timeProvider common.TimeProvider
		repo         services.CragRepository
	}
	type args struct {
		command AddCragCommand
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		err    error
	}{
		{
			name: "happy path - should not return error",
			fields: fields{
				uuidProvider: func() common.MockUUIDProvider {
					id := mockUUID
					mp := common.MockUUIDProvider{}
					mp.On("NewUUID").Return(id)
					return mp
				}(),
				timeProvider: func() common.TimeProvider {
					mp := common.MockTimeProvider{}
					mp.On("Now").Return(mockTime)
					return mp
				}(),
				repo: func() services.MockRepository {
					acc := domain.Crag{
						ID:        mockUUID,
						Name:      "test",
						Desc:      "test",
						Country:   "test",
						CreatedAt: mockTime,
					}
					mp := services.MockRepository{}
					mp.On("AddCrag", acc).Return(nil)
					return mp
				}(),
			},
			args: args{
				command: AddCragCommand{
					Name:    "test",
					Desc:    "test",
					Country: "test",
				},
			},
			err: nil,
		},
		{
			name: "repo error - should return error",
			fields: fields{
				uuidProvider: func() common.MockUUIDProvider {
					id := mockUUID
					mp := common.MockUUIDProvider{}
					mp.On("NewUUID").Return(id)
					return mp
				}(),
				timeProvider: func() common.TimeProvider {
					mp := common.MockTimeProvider{}
					mp.On("Now").Return(mockTime)
					return mp
				}(),
				repo: func() services.MockRepository {
					acc := domain.Crag{
						ID:        mockUUID,
						Name:      "test",
						Desc:      "test",
						Country:   "test",
						CreatedAt: mockTime,
					}
					mp := services.MockRepository{}
					mp.On("AddCrag", acc).Return(errors.New("test"))
					return mp
				}(),
			},
			args: args{
				command: AddCragCommand{
					Name:    "test",
					Desc:    "test",
					Country: "test",
				},
			},
			err: errors.New("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := addCragCommandHandler{
				uuidProvider: tt.fields.uuidProvider,
				timeProvider: tt.fields.timeProvider,
				repo:         tt.fields.repo,
			}

			err := h.Handle(tt.args.command)
			assert.Equal(t, err, tt.err)

		})
	}
}

func TestNewAddCragCommandHandler(t *testing.T) {
	type args struct {
		uuidProvider common.UUIDProvider
		timeProvider common.TimeProvider
		repo         services.CragRepository
	}
	tests := []struct {
		name string
		args args
		want AddCragCommandHandler
	}{
		{
			name: "should create a command handler",
			args: args{
				uuidProvider: common.MockUUIDProvider{},
				timeProvider: common.MockTimeProvider{},
				repo:         services.MockRepository{},
			},
			want: addCragCommandHandler{
				uuidProvider: common.MockUUIDProvider{},
				timeProvider: common.MockTimeProvider{},
				repo:         services.MockRepository{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAddCragCommandHandler(tt.args.uuidProvider, tt.args.timeProvider, tt.args.repo)
			assert.Equal(t, got, tt.want)
		})
	}
}
