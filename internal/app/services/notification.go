package services

// Notification provides a struct to send messages via the NotificationService
type Notification struct {
	Subject string `json:"subject"`
	Message string `json:"message"`
}

// NotificationService sends Notification
type NotificationService interface {
	Notify(notification Notification) error
}
