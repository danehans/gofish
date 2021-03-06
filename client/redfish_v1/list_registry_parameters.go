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

// NewListRegistryParams creates a new ListRegistryParams object
// with the default values initialized.
func NewListRegistryParams() *ListRegistryParams {

	return &ListRegistryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListRegistryParamsWithTimeout creates a new ListRegistryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListRegistryParamsWithTimeout(timeout time.Duration) *ListRegistryParams {

	return &ListRegistryParams{

		timeout: timeout,
	}
}

// NewListRegistryParamsWithContext creates a new ListRegistryParams object
// with the default values initialized, and the ability to set a context for a request
func NewListRegistryParamsWithContext(ctx context.Context) *ListRegistryParams {

	return &ListRegistryParams{

		Context: ctx,
	}
}

/*ListRegistryParams contains all the parameters to send to the API endpoint
for the list registry operation typically these are written to a http.Request
*/
type ListRegistryParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list registry params
func (o *ListRegistryParams) WithTimeout(timeout time.Duration) *ListRegistryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list registry params
func (o *ListRegistryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list registry params
func (o *ListRegistryParams) WithContext(ctx context.Context) *ListRegistryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list registry params
func (o *ListRegistryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *ListRegistryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
