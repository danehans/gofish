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

// GetAccountReader is a Reader for the GetAccount structure.
type GetAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountOK creates a GetAccountOK with default headers values
func NewGetAccountOK() *GetAccountOK {
	return &GetAccountOK{}
}

/*GetAccountOK handles this case with default header values.

Success
*/
type GetAccountOK struct {
	Payload *models.ManagerAccount100ManagerAccount
}

func (o *GetAccountOK) Error() string {
	return fmt.Sprintf("[GET /AccountService/Accounts/{name}][%d] getAccountOK  %+v", 200, o.Payload)
}

func (o *GetAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagerAccount100ManagerAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountUnauthorized creates a GetAccountUnauthorized with default headers values
func NewGetAccountUnauthorized() *GetAccountUnauthorized {
	return &GetAccountUnauthorized{}
}

/*GetAccountUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type GetAccountUnauthorized struct {
}

func (o *GetAccountUnauthorized) Error() string {
	return fmt.Sprintf("[GET /AccountService/Accounts/{name}][%d] getAccountUnauthorized ", 401)
}

func (o *GetAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccountForbidden creates a GetAccountForbidden with default headers values
func NewGetAccountForbidden() *GetAccountForbidden {
	return &GetAccountForbidden{}
}

/*GetAccountForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not possess authorization to perform this request.

*/
type GetAccountForbidden struct {
}

func (o *GetAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /AccountService/Accounts/{name}][%d] getAccountForbidden ", 403)
}

func (o *GetAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccountInternalServerError creates a GetAccountInternalServerError with default headers values
func NewGetAccountInternalServerError() *GetAccountInternalServerError {
	return &GetAccountInternalServerError{}
}

/*GetAccountInternalServerError handles this case with default header values.

Error
*/
type GetAccountInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetAccountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /AccountService/Accounts/{name}][%d] getAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}