package notification

import (
	"errors"
	"time"
)

type Notification struct {
	Id        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Url       string    `json:"url, omitempty" db:"url"`
	Key       string    `json:"key" db:"key"`
	Provider  string    `json:"provider" db:"provider"`
	CreatedOn time.Time `json:"created_on, omitempty" db:"created_on"`
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

	if self.Provider == "" {
		return errors.New("Notification provider cannot be empty")
	}

	return nil
}
