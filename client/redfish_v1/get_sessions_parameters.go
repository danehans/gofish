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

// NewGetSessionsParams creates a new GetSessionsParams object
// with the default values initialized.
func NewGetSessionsParams() *GetSessionsParams {

	return &GetSessionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSessionsParamsWithTimeout creates a new GetSessionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSessionsParamsWithTimeout(timeout time.Duration) *GetSessionsParams {

	return &GetSessionsParams{

		timeout: timeout,
	}
}

// NewGetSessionsParamsWithContext creates a new GetSessionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSessionsParamsWithContext(ctx context.Context) *GetSessionsParams {

	return &GetSessionsParams{

		Context: ctx,
	}
}

/*GetSessionsParams contains all the parameters to send to the API endpoint
for the get sessions operation typically these are written to a http.Request
*/
type GetSessionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sessions params
func (o *GetSessionsParams) WithTimeout(timeout time.Duration) *GetSessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sessions params
func (o *GetSessionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sessions params
func (o *GetSessionsParams) WithContext(ctx context.Context) *GetSessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sessions params
func (o *GetSessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *GetSessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
