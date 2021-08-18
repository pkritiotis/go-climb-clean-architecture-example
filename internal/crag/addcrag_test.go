package crag

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pkritiotis/go-climb/internal/notification"
	"github.com/pkritiotis/go-climb/internal/timeutil"
	"github.com/pkritiotis/go-climb/internal/uuidutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAddCragCommandHandler_Handle(t *testing.T) {
	mockTime, _ := time.Parse("yyyy-MM-02", "2021-07-30")
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")
	type fields struct {
		uuidProvider        uuidutil.UUIDProvider
		timeProvider        timeutil.TimeProvider
		repo                Repository
		notificationService notification.Service
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
				uuidProvider: func() uuidutil.MockUUIDProvider {
					id := mockUUID
					mp := uuidutil.MockUUIDProvider{}
					mp.On("NewUUID").Return(id)
					return mp
				}(),
				timeProvider: func() timeutil.TimeProvider {
					mp := timeutil.MockTimeProvider{}
					mp.On("Now").Return(mockTime)
					return mp
				}(),
				repo: func() MockRepository {
					acc := Crag{
						ID:        mockUUID,
						Name:      "test",
						Desc:      "test",
						Country:   "test",
						CreatedAt: mockTime,
					}
					mp := MockRepository{}
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
				uuidProvider: func() uuidutil.MockUUIDProvider {
					id := mockUUID
					mp := uuidutil.MockUUIDProvider{}
					mp.On("NewUUID").Return(id)
					return mp
				}(),
				timeProvider: func() timeutil.TimeProvider {
					mp := timeutil.MockTimeProvider{}
					mp.On("Now").Return(mockTime)
					return mp
				}(),
				repo: func() MockRepository {
					acc := Crag{
						ID:        mockUUID,
						Name:      "test",
						Desc:      "test",
						Country:   "test",
						CreatedAt: mockTime,
					}
					mp := MockRepository{}
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
				command: AddCragCommand{
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
				uuidProvider: func() uuidutil.MockUUIDProvider {
					id := mockUUID
					mp := uuidutil.MockUUIDProvider{}
					mp.On("NewUUID").Return(id)
					return mp
				}(),
				timeProvider: func() timeutil.TimeProvider {
					mp := timeutil.MockTimeProvider{}
					mp.On("Now").Return(mockTime)
					return mp
				}(),
				repo: func() MockRepository {
					acc := Crag{
						ID:        mockUUID,
						Name:      "test",
						Desc:      "test",
						Country:   "test",
						CreatedAt: mockTime,
					}
					mp := MockRepository{}
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
				command: AddCragCommand{
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
			h := addCragCommandHandler{
				uuidProvider:        tt.fields.uuidProvider,
				timeProvider:        tt.fields.timeProvider,
				repo:                tt.fields.repo,
				notificationService: tt.fields.notificationService,
			}

			err := h.Handle(tt.args.command)
			assert.Equal(t, err, tt.err)

		})
	}
}

func TestNewAddCragCommandHandler(t *testing.T) {
	type args struct {
		uuidProvider        uuidutil.UUIDProvider
		timeProvider        timeutil.TimeProvider
		repo                Repository
		notificationService notification.Service
	}
	tests := []struct {
		name string
		args args
		want AddCragCommandHandler
	}{
		{
			name: "should create a command handler",
			args: args{
				uuidProvider:        uuidutil.MockUUIDProvider{},
				timeProvider:        timeutil.MockTimeProvider{},
				repo:                MockRepository{},
				notificationService: notification.MockNotificationService{},
			},
			want: addCragCommandHandler{
				uuidProvider:        uuidutil.MockUUIDProvider{},
				timeProvider:        timeutil.MockTimeProvider{},
				repo:                MockRepository{},
				notificationService: notification.MockNotificationService{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAddCragCommandHandler(tt.args.uuidProvider, tt.args.timeProvider, tt.args.repo, tt.args.notificationService)
			assert.Equal(t, got, tt.want)
		})
	}
}
