package app

import (
	"github.com/pkritiotis/go-climb/internal/adapters/notification"
	"github.com/pkritiotis/go-climb/internal/app/commands"
	"github.com/pkritiotis/go-climb/internal/app/common"
	"github.com/pkritiotis/go-climb/internal/app/queries"
	services2 "github.com/pkritiotis/go-climb/internal/app/services"
	"github.com/pkritiotis/go-climb/internal/domain/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewApp(t *testing.T) {
	mockRepo := services.MockRepository{}
	UUIDProvider := common.NewUUIDProvider()
	timeProvider := common.NewTimeProvider()
	notificationService := notification.ConsoleNotificationService{}

	type args struct {
		up                  common.UUIDProvider
		tp                  common.TimeProvider
		cragRepo            services.CragRepository
		notificationService services2.NotificationService
	}
	tests := []struct {
		name string
		args args
		want App
	}{
		{
			name: "should initialize application layer",
			args: args{
				cragRepo:            mockRepo,
				notificationService: notificationService,
			},
			want: App{
				Queries: Queries{
					GetAllCragsHandler: queries.NewGetAllCragsQueryHandler(mockRepo),
					GetCragHandler:     queries.NewGetCragQueryHandler(mockRepo),
				},
				Commands: Commands{
					AddCragHandler:    commands.NewAddCragCommandHandler(UUIDProvider, timeProvider, mockRepo, notificationService),
					UpdateCragHandler: commands.NewUpdateCragCommandHandler(mockRepo),
					DeleteCragHandler: commands.NewDeleteCragCommandHandler(mockRepo),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewApp(tt.args.cragRepo, tt.args.notificationService)
			assert.Equal(t, tt.want, got)
		})
	}
}
