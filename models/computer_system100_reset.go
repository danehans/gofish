package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// ComputerSystem100Reset computer system 1 0 0 reset
// swagger:model ComputerSystem.1.0.0_Reset
type ComputerSystem100Reset struct {

	// Link to invoke action
	Target strfmt.URI `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// Validate validates this computer system 1 0 0 reset
func (m *ComputerSystem100Reset) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}