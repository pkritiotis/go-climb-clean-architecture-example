package notification

// Notification provides a struct to send messages via the Service
type Notification struct {
	Subject string `json:"subject"`
	Message string `json:"message"`
}

// Service sends Notification
type Service interface {
	Notify(notification Notification) error
}
