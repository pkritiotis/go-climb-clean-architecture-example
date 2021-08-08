package notification

import (
	"github.com/pkritiotis/go-clean/internal/app/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConsoleManager_Notify(t *testing.T) {
	type args struct {
		notification services.Notification
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should not return error",
			args: args{
				notification: services.Notification{
					Subject: "Test Subject",
					Message: "Test Message",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			co := ConsoleManager{}
			err := co.Notify(tt.args.notification)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
