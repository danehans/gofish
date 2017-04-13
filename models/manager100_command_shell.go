package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Manager100CommandShell Used for describing services like Serial Console, Command Shell or Graphical Console
// swagger:model Manager.1.0.0_CommandShell
type Manager100CommandShell struct {

	// This object is used to enumerate the Command Shell connection types allowed by the implementation.
	// Read Only: true
	ConnectTypesSupported []string `json:"ConnectTypesSupported"`

	// Indicates the maximum number of service sessions, regardless of protocol, this manager is able to support.
	// Read Only: true
	// Minimum: 0
	MaxConcurrentSessions float64 `json:"MaxConcurrentSessions,omitempty"`

	// Indicates if the service is enabled for this manager.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`
}

// Validate validates this manager 1 0 0 command shell
func (m *Manager100CommandShell) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectTypesSupported(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMaxConcurrentSessions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var manager100CommandShellConnectTypesSupportedItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SSH","Telnet","IPMI","Oem"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		manager100CommandShellConnectTypesSupportedItemsEnum = append(manager100CommandShellConnectTypesSupportedItemsEnum, v)
	}
}

func (m *Manager100CommandShell) validateConnectTypesSupportedItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, manager100CommandShellConnectTypesSupportedItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *Manager100CommandShell) validateConnectTypesSupported(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectTypesSupported) { // not required
		return nil
	}

	for i := 0; i < len(m.ConnectTypesSupported); i++ {

		// value enum
		if err := m.validateConnectTypesSupportedItemsEnum("ConnectTypesSupported"+"."+strconv.Itoa(i), "body", m.ConnectTypesSupported[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *Manager100CommandShell) validateMaxConcurrentSessions(formats strfmt.Registry) error {

	if swag.IsZero(m.MaxConcurrentSessions) { // not required
		return nil
	}

	if err := validate.Minimum("MaxConcurrentSessions", "body", float64(m.MaxConcurrentSessions), 0, false); err != nil {
		return err
	}

	return nil
}