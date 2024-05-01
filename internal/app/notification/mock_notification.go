// Package notification contains the mock implementation of the NotificationService interface.
package notification

import "github.com/stretchr/testify/mock"

// MockNotificationService sends mock Notifications
type MockNotificationService struct {
	mock.Mock
}

// Notify sends mock Notifications
func (m MockNotificationService) Notify(notification Notification) error {
	args := m.Called(notification)
	return args.Error(0)
}
