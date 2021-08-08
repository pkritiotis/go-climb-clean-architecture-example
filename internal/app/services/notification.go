package services

// Notification provides a struct to send messages via the NotificationManager
type Notification struct {
	Subject string `json:"subject"`
	Message string `json:"message"`
}

// NotificationManager sends Notification
type NotificationManager interface {
	Notify(notification Notification) error
}
