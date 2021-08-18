package notification

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConsoleNotificationService_Notify(t *testing.T) {
	type args struct {
		notification Notification
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should not return error",
			args: args{
				notification: Notification{
					Subject: "Test Subject",
					Message: "Test Message",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			co := ConsoleNotificationService{}
			err := co.Notify(tt.args.notification)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
