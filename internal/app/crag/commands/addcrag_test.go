package commands

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/app/notification"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/domain/crag"
	timeUtil "github.com/pkritiotis/go-climb-clean-architecture-example/internal/pkg/time"
	uuidUtil "github.com/pkritiotis/go-climb-clean-architecture-example/internal/pkg/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAddCragCommandHandler_Handle(t *testing.T) {
	mockTime, _ := time.Parse("yyyy-MM-02", "2021-07-30")
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")
	type fields struct {
		uuidProvider        uuidUtil.Provider
		timeProvider        timeUtil.Provider
		repo                crag.Repository
		notificationService notification.Service
	}
	type args struct {
		request AddCragRequest
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
				uuidProvider: func() uuidUtil.MockProvider {
					id := mockUUID
					mp := uuidUtil.MockProvider{}
					mp.On("NewUUID").Return(id)
					return mp
				}(),
				timeProvider: func() timeUtil.Provider {
					mp := timeUtil.MockProvider{}
					mp.On("Now").Return(mockTime)
					return mp
				}(),
				repo: func() crag.MockRepository {
					acc := crag.Crag{
						ID:        mockUUID,
						Name:      "test",
						Desc:      "test",
						Country:   "test",
						CreatedAt: mockTime,
					}
					mp := crag.MockRepository{}
					mp.On("Add", acc).Return(nil)
					return mp
				}(),
				notificationService: func() notification.MockNotificationService {
					mock := notification.MockNotificationService{}
					n := notification.Notification{
						Subject: "New crag added",
						Message: "A new crag with name 'test' was added in the repository",
					}
					mock.On("Notify", n).Return(nil)
					return mock
				}(),
			},
			args: args{
				request: AddCragRequest{
					Name:    "test",
					Desc:    "test",
					Country: "test",
				},
			},
			err: nil,
		},
		{
			name: "memory error - should return error",
			fields: fields{
				uuidProvider: func() uuidUtil.MockProvider {
					id := mockUUID
					mp := uuidUtil.MockProvider{}
					mp.On("NewUUID").Return(id)
					return mp
				}(),
				timeProvider: func() timeUtil.Provider {
					mp := timeUtil.MockProvider{}
					mp.On("Now").Return(mockTime)
					return mp
				}(),
				repo: func() crag.MockRepository {
					acc := crag.Crag{
						ID:        mockUUID,
						Name:      "test",
						Desc:      "test",
						Country:   "test",
						CreatedAt: mockTime,
					}
					mp := crag.MockRepository{}
					mp.On("Add", acc).Return(errors.New("test"))
					return mp
				}(),
				notificationService: func() notification.MockNotificationService {
					mock := notification.MockNotificationService{}
					n := notification.Notification{
						Subject: "New crag added",
						Message: "A new crag with name 'test' was added in the repository",
					}
					mock.On("Notify", n).Return(nil)
					return mock
				}(),
			},

			args: args{
				request: AddCragRequest{
					Name:    "test",
					Desc:    "test",
					Country: "test",
				},
			},
			err: errors.New("test"),
		},
		{
			name: "happy path - should not return error",
			fields: fields{
				uuidProvider: func() uuidUtil.MockProvider {
					id := mockUUID
					mp := uuidUtil.MockProvider{}
					mp.On("NewUUID").Return(id)
					return mp
				}(),
				timeProvider: func() timeUtil.Provider {
					mp := timeUtil.MockProvider{}
					mp.On("Now").Return(mockTime)
					return mp
				}(),
				repo: func() crag.MockRepository {
					acc := crag.Crag{
						ID:        mockUUID,
						Name:      "test",
						Desc:      "test",
						Country:   "test",
						CreatedAt: mockTime,
					}
					mp := crag.MockRepository{}
					mp.On("Add", acc).Return(nil)
					return mp
				}(),
				notificationService: func() notification.MockNotificationService {
					mock := notification.MockNotificationService{}
					n := notification.Notification{
						Subject: "New crag added",
						Message: "A new crag with name 'test' was added in the repository",
					}
					mock.On("Notify", n).Return(errors.New("notification error"))
					return mock
				}(),
			},
			args: args{
				request: AddCragRequest{
					Name:    "test",
					Desc:    "test",
					Country: "test",
				},
			},
			err: errors.New("notification error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := addCragRequestHandler{
				uuidProvider:        tt.fields.uuidProvider,
				timeProvider:        tt.fields.timeProvider,
				repo:                tt.fields.repo,
				notificationService: tt.fields.notificationService,
			}

			err := h.Handle(tt.args.request)
			assert.Equal(t, err, tt.err)

		})
	}
}

func TestNewAddCragCommandHandler(t *testing.T) {
	type args struct {
		uuidProvider        uuidUtil.Provider
		timeProvider        timeUtil.Provider
		repo                crag.Repository
		notificationService notification.Service
	}
	tests := []struct {
		name string
		args args
		want CreateCragRequestHandler
	}{
		{
			name: "should create a request handler",
			args: args{
				uuidProvider:        uuidUtil.MockProvider{},
				timeProvider:        timeUtil.MockProvider{},
				repo:                crag.MockRepository{},
				notificationService: notification.MockNotificationService{},
			},
			want: addCragRequestHandler{
				uuidProvider:        uuidUtil.MockProvider{},
				timeProvider:        timeUtil.MockProvider{},
				repo:                crag.MockRepository{},
				notificationService: notification.MockNotificationService{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAddCragRequestHandler(tt.args.uuidProvider, tt.args.timeProvider, tt.args.repo, tt.args.notificationService)
			assert.Equal(t, got, tt.want)
		})
	}
}
