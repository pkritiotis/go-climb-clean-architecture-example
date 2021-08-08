package notification

import (
	"encoding/json"
	"fmt"
	"github.com/pkritiotis/go-clean/internal/app/services"
)

// ConsoleManager provides a console implementation of the NotificationManager
type ConsoleManager struct{}

// Notify prints out the notifications in console
func (ConsoleManager) Notify(notification services.Notification) error {
	jsonNotification, err := json.Marshal(notification)
	if err != nil {
		return err
	}
	fmt.Printf("Notification Received: %v", string(jsonNotification))
	return nil
}
