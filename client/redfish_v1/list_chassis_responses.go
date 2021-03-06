package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/danehans/gofish/models"
)

// ListChassisReader is a Reader for the ListChassis structure.
type ListChassisReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListChassisReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListChassisOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListChassisBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListChassisUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListChassisForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListChassisInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListChassisOK creates a ListChassisOK with default headers values
func NewListChassisOK() *ListChassisOK {
	return &ListChassisOK{}
}

/*ListChassisOK handles this case with default header values.

Success
*/
type ListChassisOK struct {
	Payload *models.ChassisCollectionChassisCollection
}

func (o *ListChassisOK) Error() string {
	return fmt.Sprintf("[GET /Chassis][%d] listChassisOK  %+v", 200, o.Payload)
}

func (o *ListChassisOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChassisCollectionChassisCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListChassisBadRequest creates a ListChassisBadRequest with default headers values
func NewListChassisBadRequest() *ListChassisBadRequest {
	return &ListChassisBadRequest{}
}

/*ListChassisBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information (such as validation error on an input field, a missing required value, and so on). An extended error shall be returned in the response body, as defined in section Extended Error Handling.

*/
type ListChassisBadRequest struct {
}

func (o *ListChassisBadRequest) Error() string {
	return fmt.Sprintf("[GET /Chassis][%d] listChassisBadRequest ", 400)
}

func (o *ListChassisBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListChassisUnauthorized creates a ListChassisUnauthorized with default headers values
func NewListChassisUnauthorized() *ListChassisUnauthorized {
	return &ListChassisUnauthorized{}
}

/*ListChassisUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type ListChassisUnauthorized struct {
}

func (o *ListChassisUnauthorized) Error() string {
	return fmt.Sprintf("[GET /Chassis][%d] listChassisUnauthorized ", 401)
}

func (o *ListChassisUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListChassisForbidden creates a ListChassisForbidden with default headers values
func NewListChassisForbidden() *ListChassisForbidden {
	return &ListChassisForbidden{}
}

/*ListChassisForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not possess authorization to perform this request.

*/
type ListChassisForbidden struct {
}

func (o *ListChassisForbidden) Error() string {
	return fmt.Sprintf("[GET /Chassis][%d] listChassisForbidden ", 403)
}

func (o *ListChassisForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListChassisInternalServerError creates a ListChassisInternalServerError with default headers values
func NewListChassisInternalServerError() *ListChassisInternalServerError {
	return &ListChassisInternalServerError{}
}

/*ListChassisInternalServerError handles this case with default header values.

Error
*/
type ListChassisInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ListChassisInternalServerError) Error() string {
	return fmt.Sprintf("[GET /Chassis][%d] listChassisInternalServerError  %+v", 500, o.Payload)
}

func (o *ListChassisInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
