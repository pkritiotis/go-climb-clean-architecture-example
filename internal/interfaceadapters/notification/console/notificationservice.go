package console

import (
	"encoding/json"
	"fmt"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/app/notification"
)

// NotificationService provides a console implementation of the Service
type NotificationService struct{}

// NewNotificationService constructor for NotificationService
func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

// Notify prints out the notifications in console
func (NotificationService) Notify(notification notification.Notification) error {
	jsonNotification, err := json.Marshal(notification)
	if err != nil {
		return err
	}
	fmt.Printf("Notification Received: %v", string(jsonNotification))
	return nil
}
