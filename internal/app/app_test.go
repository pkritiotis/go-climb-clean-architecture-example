package app

import (
	"github.com/pkritiotis/go-climb/internal/app/crag/commands"
	"github.com/pkritiotis/go-climb/internal/app/crag/queries"
	"github.com/pkritiotis/go-climb/internal/app/notification"
	"github.com/pkritiotis/go-climb/internal/domain/crag"
	"github.com/pkritiotis/go-climb/internal/pkg/time"
	"github.com/pkritiotis/go-climb/internal/pkg/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewApp(t *testing.T) {
	mockRepo := crag.MockRepository{}
	UUIDProvider := &uuid.MockProvider{}
	timeProvider := &time.MockProvider{}
	notificationService := notification.MockNotificationService{}

	type args struct {
		up                  uuid.Provider
		tp                  time.Provider
		cragRepo            crag.Repository
		notificationService notification.Service
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
				up:                  UUIDProvider,
				tp:                  timeProvider,
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
			got := NewApp(tt.args.cragRepo, tt.args.notificationService, tt.args.up, tt.args.tp)
			assert.Equal(t, tt.want, got)
		})
	}
}
