package setting

import (
	"errors"
)

type Setting struct {
	Id           int    `json:"id" db:"id"`
	Key          string `json:"key" db:"key"`
	Value        string `json:"value" db:"value"`
	Provider     int    `json:"provider, omitempty" db:"provider"`
	ProviderName string `json:"provider_name" db:"name"`
}

func (self *Setting) ValidInsert() error {
	if self.Id > 0 {
		return errors.New("Setting ID already set")
	}

	if self.Key == "" {
		return errors.New("Setting key cannot be empty")
	}

	if self.Value == "" {
		return errors.New("Setting value cannot be empty")
	}

	if self.ProviderName == "" {
		return errors.New("Setting provider name cannot be empty")
	}

	return nil
}

func (self *Setting) ValidUpdate() error {
	if self.Key == "" {
		return errors.New("Setting key cannot be empty")
	}

	if self.Value == "" {
		return errors.New("Setting value cannot be empty")
	}

	if self.ProviderName == "" {
		return errors.New("Setting provider name cannot be empty")
	}

	return nil
}
