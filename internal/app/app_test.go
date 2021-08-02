package app

import (
	"github.com/pkritiotis/go-clean/internal/app/common"
	"github.com/pkritiotis/go-clean/internal/app/commands"
	"github.com/pkritiotis/go-clean/internal/app/queries"
	"github.com/pkritiotis/go-clean/internal/app/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewApp(t *testing.T) {
	mockRepo := services.MockRepository{}
	mockUUIDProvider := common.MockUUIDProvider{}
	mockTimeProvider := common.MockTimeProvider{}

	type args struct {
		up       common.UUIDProvider
		tp       common.TimeProvider
		cragRepo services.CragRepository
	}
	tests := []struct {
		name string
		args args
		want App
	}{
		{
			name: "should initialize application layer",
			args: args{
				up:       mockUUIDProvider,
				tp:       mockTimeProvider,
				cragRepo: mockRepo,
			},
			want: App{
				Queries: Queries{
					GetAllCragsHandler: queries.NewGetAllCragsQueryHandler(mockRepo),
					GetCragHandler:     queries.NewGetCragQueryHandler(mockRepo),
				},
				Commands: Commands{
					AddCragHandler:    commands.NewAddCragCommandHandler(mockUUIDProvider, mockTimeProvider, mockRepo),
					UpdateCragHandler: commands.NewUpdateCragCommandHandler(mockRepo),
					DeleteCragHandler: commands.NewDeleteCragCommandHandler(mockRepo),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewApp(tt.args.up, tt.args.tp, tt.args.cragRepo)
			assert.Equal(t, tt.want, got)
		})
	}
}
