package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ManagerNetworkProtocol100Protocol manager network protocol 1 0 0 protocol
// swagger:model ManagerNetworkProtocol.1.0.0_Protocol
type ManagerNetworkProtocol100Protocol struct {

	// Indicates the protocol port.
	// Minimum: 0
	Port *float64 `json:"Port,omitempty"`

	// Indicates if the protocol is enabled or disabled
	ProtocolEnabled bool `json:"ProtocolEnabled,omitempty"`
}

// Validate validates this manager network protocol 1 0 0 protocol
func (m *ManagerNetworkProtocol100Protocol) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePort(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagerNetworkProtocol100Protocol) validatePort(formats strfmt.Registry) error {

	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.Minimum("Port", "body", float64(*m.Port), 0, false); err != nil {
		return err
	}

	return nil
}