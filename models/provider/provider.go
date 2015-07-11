package provider

import (
	"errors"
)

type Provider struct {
	Id          int    `json:"id, omitempty" db:"id"`
	Name        string `json:"name" db:"name"`
	DisplayName string `json:"display_name" db:"display_name"`
	Description string `json:"description, omitempty" db:"description"`
	WebURL      string `json:"web_url, omitempty" db:"web_url"`
	IconURL     string `json:"icon_url, omitempty" db:"icon_url"`
	Active      bool   `json:"active, omitempty" db:"active"`
}

func (self *Provider) ValidInsert() error {
	if self.Id > 0 {
		return errors.New("Provider ID already set")
	}

	if self.Name == "" {
		return errors.New("Provider name cannot be empty")
	}

	if self.DisplayName == "" {
		return errors.New("Provider display name cannot be empty")
	}

	return nil
}

func (self *Provider) ValidUpdate() error {

	if self.DisplayName == "" {
		return errors.New("Provider display name cannot be empty")
	}

	return nil
}
