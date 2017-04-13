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

// VirtualMedia100VirtualMedia This is the schema definition for the Virtual Media Service.
// swagger:model VirtualMedia.1.0.0_VirtualMedia
type VirtualMedia100VirtualMedia struct {

	// at odata context
	// Read Only: true
	AtOdataContext strfmt.URI `json:"@odata.context,omitempty"`

	// at odata id
	// Read Only: true
	AtOdataID strfmt.URI `json:"@odata.id,omitempty"`

	// at odata type
	// Read Only: true
	AtOdataType string `json:"@odata.type,omitempty"`

	// Current virtual media connection methods
	// Read Only: true
	ConnectedVia string `json:"ConnectedVia,omitempty"`

	// Provides a description of this resource and is used for commonality  in the schema definitions.
	// Read Only: true
	Description string `json:"Description,omitempty"`

	// Uniquely identifies the resource within the collection of like resources.
	// Read Only: true
	ID string `json:"Id,omitempty"`

	// A URI providing the location of the selected image
	// Read Only: true
	Image string `json:"Image,omitempty"`

	// The current image name
	// Read Only: true
	ImageName string `json:"ImageName,omitempty"`

	// Indicates if virtual media is inserted in the virtual device.
	// Read Only: true
	Inserted *bool `json:"Inserted,omitempty"`

	// This is the media types supported as virtual media.
	// Read Only: true
	MediaTypes []string `json:"MediaTypes"`

	// The name of the resource or array element.
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	Oem ResourceOem `json:"Oem,omitempty"`

	// Indicates the media is write protected.
	// Read Only: true
	WriteProtected *bool `json:"WriteProtected,omitempty"`
}

// Validate validates this virtual media 1 0 0 virtual media
func (m *VirtualMedia100VirtualMedia) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectedVia(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMediaTypes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var virtualMedia100VirtualMediaTypeConnectedViaPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NotConnected","URI","Applet","Oem"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		virtualMedia100VirtualMediaTypeConnectedViaPropEnum = append(virtualMedia100VirtualMediaTypeConnectedViaPropEnum, v)
	}
}

const (
	// VirtualMedia100VirtualMediaConnectedViaNotConnected captures enum value "NotConnected"
	VirtualMedia100VirtualMediaConnectedViaNotConnected string = "NotConnected"
	// VirtualMedia100VirtualMediaConnectedViaURI captures enum value "URI"
	VirtualMedia100VirtualMediaConnectedViaURI string = "URI"
	// VirtualMedia100VirtualMediaConnectedViaApplet captures enum value "Applet"
	VirtualMedia100VirtualMediaConnectedViaApplet string = "Applet"
	// VirtualMedia100VirtualMediaConnectedViaOem captures enum value "Oem"
	VirtualMedia100VirtualMediaConnectedViaOem string = "Oem"
)

// prop value enum
func (m *VirtualMedia100VirtualMedia) validateConnectedViaEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, virtualMedia100VirtualMediaTypeConnectedViaPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VirtualMedia100VirtualMedia) validateConnectedVia(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectedVia) { // not required
		return nil
	}

	// value enum
	if err := m.validateConnectedViaEnum("ConnectedVia", "body", m.ConnectedVia); err != nil {
		return err
	}

	return nil
}

var virtualMedia100VirtualMediaMediaTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CD","Floppy","USBStick","DVD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		virtualMedia100VirtualMediaMediaTypesItemsEnum = append(virtualMedia100VirtualMediaMediaTypesItemsEnum, v)
	}
}

func (m *VirtualMedia100VirtualMedia) validateMediaTypesItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, virtualMedia100VirtualMediaMediaTypesItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *VirtualMedia100VirtualMedia) validateMediaTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.MediaTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.MediaTypes); i++ {

		// value enum
		if err := m.validateMediaTypesItemsEnum("MediaTypes"+"."+strconv.Itoa(i), "body", m.MediaTypes[i]); err != nil {
			return err
		}

	}

	return nil
}