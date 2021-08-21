package notification

import (
	"encoding/json"
	"fmt"
	"github.com/pkritiotis/go-climb/internal/app/notification"
)

// ConsoleService provides a console implementation of the Service
type ConsoleService struct{}

// NewConsoleService constructor for ConsoleService
func NewConsoleService() *ConsoleService {
	return &ConsoleService{}
}

// Notify prints out the notifications in console
func (ConsoleService) Notify(notification notification.Notification) error {
	jsonNotification, err := json.Marshal(notification)
	if err != nil {
		return err
	}
	fmt.Printf("Notification Received: %v", string(jsonNotification))
	return nil
}
