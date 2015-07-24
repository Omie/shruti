package notification

import (
	"errors"
	"time"
)

const (
	// Priorities
	PRIO_LOW  = 10
	PRIO_MED  = 20
	PRIO_HIGH = 30

	// Actions
	ACT_POLL = 10
	ACT_PUSH = 20

	// Heard status
	HRD_UNHEARD = 10
	HRD_HEARD   = 20
)

type Notification struct {
	Id        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Url       string    `json:"url, omitempty" db:"url"`
	Key       string    `json:"key" db:"key"`
	Heard     int       `json:"heard" db:"heard"`
	Provider  int       `json:"provider" db:"provider"`
	CreatedOn time.Time `json:"created_on, omitempty" db:"created_on"`
	Priority  int       `json:"priority" db:"priority"`
	Action    int       `json:"action" db:"action"`

	// this exists to enable providers to push without knowing their id
	ProviderName string `json:"provider_name, omitempty" db:"provider_name"`
}

func (self *Notification) ValidInsert() error {
	if self.Id > 0 {
		return errors.New("Notification ID already set")
	}

	if self.Title == "" {
		return errors.New("Notification title cannot be empty")
	}

	if self.Key == "" {
		return errors.New("Notification key cannot be empty")
	}

	if self.ProviderName == "" {
		return errors.New("Notification provider name cannot be empty")
	}

	if self.Priority != PRIO_LOW && self.Priority != PRIO_MED &&
		self.Priority != PRIO_HIGH {
		return errors.New("Invalid priority")
	}

	if self.Action != ACT_POLL && self.Action != ACT_PUSH {
		return errors.New("Invalid default action")
	}

	return nil
}
