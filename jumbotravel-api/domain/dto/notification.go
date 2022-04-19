package dto

import (
	"encoding/json"
	"fmt"
	"time"
)

type NotificationInput struct {
	Scope            *string
	ResourceId       *int
	ResourceUuid     *string
	Title            *string
	Message          *string
	Link             *string
	Extra            *map[string]string
	NotificationType *string
	Popup            *string
	ExpiresAt        *time.Time
}

func (n NotificationInput) Validate() error {

	if n.Scope == nil {
		return fmt.Errorf("scope is required")
	}

	if n.Title == nil {
		return fmt.Errorf("title is required")
	}

	if n.NotificationType == nil {
		return fmt.Errorf("notificationType is required")
	}

	if n.ExpiresAt == nil {
		return fmt.Errorf("expiresAt is required")
	}

	if n.Popup != nil && *n.Popup != "0" && *n.Popup != "1" && *n.Popup != "2" {
		return fmt.Errorf("popup must be nil, 0, 1 or 2")
	}

	return nil
}

func (n NotificationInput) Build() (string, []interface{}, error) {

	partialQuery := "("
	args := make([]interface{}, 0)

	err := n.Validate()
	if err != nil {
		return "", nil, err
	}

	// Scope
	partialQuery = fmt.Sprintf("%s?", partialQuery)
	if n.Scope != nil {
		args = append(args, *n.Scope)
	} else {
		args = append(args, "")
	}

	// ResourceId
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	if n.ResourceId != nil {
		args = append(args, *n.ResourceId)
	} else {
		args = append(args, nil)
	}

	// ResourceUuid
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	if n.ResourceUuid != nil {
		args = append(args, *n.ResourceUuid)
	} else {
		args = append(args, nil)
	}

	// Title
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	if n.Title != nil {
		args = append(args, *n.Title)
	} else {
		args = append(args, nil)
	}

	// Message
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	if n.Message != nil {
		args = append(args, *n.Message)
	} else {
		args = append(args, nil)
	}

	// Link
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	if n.Link != nil {
		args = append(args, *n.Link)
	} else {
		args = append(args, nil)
	}

	// Extra
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	if n.Extra != nil {
		parsedMap, err := json.Marshal(*n.Extra)
		if err != nil {
			return "", nil, err
		}
		args = append(args, string(parsedMap))
	} else {
		args = append(args, nil)
	}

	// NotificationType
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	if n.NotificationType != nil {
		args = append(args, *n.NotificationType)
	} else {
		args = append(args, nil)
	}

	// Popup
	partialQuery = fmt.Sprintf("%s,?", partialQuery)
	var parsedPopup *bool
	parsedPopup = nil
	if n.Popup != nil && *n.Popup != "0" {
		parsedPopup = new(bool)
		if *n.Popup == "1" {
			*parsedPopup = true
		} else {
			*parsedPopup = false
		}
		args = append(args, *parsedPopup)
	} else {
		args = append(args, false)
	}

	// ExpiresAt
	partialQuery = fmt.Sprintf("%s,TIMESTAMP(?)", partialQuery)
	if n.ExpiresAt != nil {
		args = append(args, (*n.ExpiresAt).Format("2006-01-02 15:04:05"))
	} else {
		args = append(args, nil)
	}

	partialQuery = fmt.Sprintf("%s)", partialQuery)

	return partialQuery, args, nil
}
