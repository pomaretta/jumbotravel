package entity

import "time"

type Notification struct {
	NotificationId *int       `json:"notification_id"`
	Scope          *string    `json:"scope"`
	ResourceId     *int       `json:"resource_id"`
	Title          *string    `json:"title"`
	Message        *string    `json:"message"`
	Link           *string    `json:"link"`
	Extra          *string    `json:"extra"`
	Type           *string    `json:"type"`
	Popup          *bool      `json:"popup"`
	ExpiresAt      *time.Time `json:"expires_at"`
	CreatedAt      *time.Time `json:"created_at"`
	Seen           *bool      `json:"seen"`
	Active         *bool      `json:"active"`
	Signature      *string    `json:"signature"`
} // @name Notification

func (v *Notification) GetDestFields() []interface{} {
	return []interface{}{
		&v.NotificationId,
		&v.Scope,
		&v.ResourceId,
		&v.Title,
		&v.Message,
		&v.Link,
		&v.Extra,
		&v.Type,
		&v.Popup,
		&v.ExpiresAt,
		&v.CreatedAt,
		&v.Seen,
		&v.Active,
	}
}

func (v *Notification) New() {
	*v = Notification{}
}

func (v *Notification) Val() interface{} {
	return *v
}
