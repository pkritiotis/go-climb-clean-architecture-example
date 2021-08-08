package notification

import (
	"encoding/json"
	"fmt"
	"github.com/pkritiotis/go-clean/internal/app/services"
)

// ConsoleNotificationService provides a console implementation of the NotificationService
type ConsoleNotificationService struct{}

// Notify prints out the notifications in console
func (ConsoleNotificationService) Notify(notification services.Notification) error {
	jsonNotification, err := json.Marshal(notification)
	if err != nil {
		return err
	}
	fmt.Printf("Notification Received: %v", string(jsonNotification))
	return nil
}
