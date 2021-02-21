package model

import (
	"errors"
)


//Country table
type Country struct {
	Name     string    `json:"name"`
	Location    string    `json:"location"`
	Wiki    string  `json:"wiki"`
}

func (c *Country) Validate() error {
	if c.Name == " " {
		return errors.New("should have Name")
	}
	if c.Location == " " {
		return errors.New("should have Location")
	}
	if c.Wiki == " " {
		return errors.New("should link Wiki")
	}
	return nil
}
