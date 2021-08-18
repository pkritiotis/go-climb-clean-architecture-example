package notification

import (
	"encoding/json"
	"fmt"
)

// ConsoleNotificationService provides a console implementation of the Service
type ConsoleNotificationService struct{}

// Notify prints out the notifications in console
func (ConsoleNotificationService) Notify(notification Notification) error {
	jsonNotification, err := json.Marshal(notification)
	if err != nil {
		return err
	}
	fmt.Printf("Notification Received: %v", string(jsonNotification))
	return nil
}
