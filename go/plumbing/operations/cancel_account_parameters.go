// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCancelAccountParams creates a new CancelAccountParams object
// with the default values initialized.
func NewCancelAccountParams() *CancelAccountParams {
	var ()
	return &CancelAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCancelAccountParamsWithTimeout creates a new CancelAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCancelAccountParamsWithTimeout(timeout time.Duration) *CancelAccountParams {
	var ()
	return &CancelAccountParams{

		timeout: timeout,
	}
}

// NewCancelAccountParamsWithContext creates a new CancelAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewCancelAccountParamsWithContext(ctx context.Context) *CancelAccountParams {
	var ()
	return &CancelAccountParams{

		Context: ctx,
	}
}

// NewCancelAccountParamsWithHTTPClient creates a new CancelAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCancelAccountParamsWithHTTPClient(client *http.Client) *CancelAccountParams {
	var ()
	return &CancelAccountParams{
		HTTPClient: client,
	}
}

/*CancelAccountParams contains all the parameters to send to the API endpoint
for the cancel account operation typically these are written to a http.Request
*/
type CancelAccountParams struct {

	/*AccountID*/
	AccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cancel account params
func (o *CancelAccountParams) WithTimeout(timeout time.Duration) *CancelAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel account params
func (o *CancelAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel account params
func (o *CancelAccountParams) WithContext(ctx context.Context) *CancelAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel account params
func (o *CancelAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel account params
func (o *CancelAccountParams) WithHTTPClient(client *http.Client) *CancelAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel account params
func (o *CancelAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the cancel account params
func (o *CancelAccountParams) WithAccountID(accountID string) *CancelAccountParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the cancel account params
func (o *CancelAccountParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *CancelAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param account_id
	if err := r.SetPathParam("account_id", o.AccountID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
