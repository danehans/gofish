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

// NewGetAccountsParams creates a new GetAccountsParams object
// with the default values initialized.
func NewGetAccountsParams() *GetAccountsParams {

	return &GetAccountsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountsParamsWithTimeout creates a new GetAccountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountsParamsWithTimeout(timeout time.Duration) *GetAccountsParams {

	return &GetAccountsParams{

		timeout: timeout,
	}
}

// NewGetAccountsParamsWithContext creates a new GetAccountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountsParamsWithContext(ctx context.Context) *GetAccountsParams {

	return &GetAccountsParams{

		Context: ctx,
	}
}

/*GetAccountsParams contains all the parameters to send to the API endpoint
for the get accounts operation typically these are written to a http.Request
*/
type GetAccountsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get accounts params
func (o *GetAccountsParams) WithTimeout(timeout time.Duration) *GetAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get accounts params
func (o *GetAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get accounts params
func (o *GetAccountsParams) WithContext(ctx context.Context) *GetAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get accounts params
func (o *GetAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
