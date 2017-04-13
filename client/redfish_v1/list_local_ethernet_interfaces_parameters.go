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

// NewListLocalEthernetInterfacesParams creates a new ListLocalEthernetInterfacesParams object
// with the default values initialized.
func NewListLocalEthernetInterfacesParams() *ListLocalEthernetInterfacesParams {

	return &ListLocalEthernetInterfacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListLocalEthernetInterfacesParamsWithTimeout creates a new ListLocalEthernetInterfacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListLocalEthernetInterfacesParamsWithTimeout(timeout time.Duration) *ListLocalEthernetInterfacesParams {

	return &ListLocalEthernetInterfacesParams{

		timeout: timeout,
	}
}

// NewListLocalEthernetInterfacesParamsWithContext creates a new ListLocalEthernetInterfacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListLocalEthernetInterfacesParamsWithContext(ctx context.Context) *ListLocalEthernetInterfacesParams {

	return &ListLocalEthernetInterfacesParams{

		Context: ctx,
	}
}

/*ListLocalEthernetInterfacesParams contains all the parameters to send to the API endpoint
for the list local ethernet interfaces operation typically these are written to a http.Request
*/
type ListLocalEthernetInterfacesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list local ethernet interfaces params
func (o *ListLocalEthernetInterfacesParams) WithTimeout(timeout time.Duration) *ListLocalEthernetInterfacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list local ethernet interfaces params
func (o *ListLocalEthernetInterfacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list local ethernet interfaces params
func (o *ListLocalEthernetInterfacesParams) WithContext(ctx context.Context) *ListLocalEthernetInterfacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list local ethernet interfaces params
func (o *ListLocalEthernetInterfacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *ListLocalEthernetInterfacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}