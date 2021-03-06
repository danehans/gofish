package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPowerParams creates a new GetPowerParams object
// with the default values initialized.
func NewGetPowerParams() *GetPowerParams {
	var ()
	return &GetPowerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPowerParamsWithTimeout creates a new GetPowerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPowerParamsWithTimeout(timeout time.Duration) *GetPowerParams {
	var ()
	return &GetPowerParams{

		timeout: timeout,
	}
}

// NewGetPowerParamsWithContext creates a new GetPowerParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPowerParamsWithContext(ctx context.Context) *GetPowerParams {
	var ()
	return &GetPowerParams{

		Context: ctx,
	}
}

/*GetPowerParams contains all the parameters to send to the API endpoint
for the get power operation typically these are written to a http.Request
*/
type GetPowerParams struct {

	/*Identifier*/
	Identifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get power params
func (o *GetPowerParams) WithTimeout(timeout time.Duration) *GetPowerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get power params
func (o *GetPowerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get power params
func (o *GetPowerParams) WithContext(ctx context.Context) *GetPowerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get power params
func (o *GetPowerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithIdentifier adds the identifier to the get power params
func (o *GetPowerParams) WithIdentifier(identifier string) *GetPowerParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the get power params
func (o *GetPowerParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WriteToRequest writes these params to a swagger request
func (o *GetPowerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
