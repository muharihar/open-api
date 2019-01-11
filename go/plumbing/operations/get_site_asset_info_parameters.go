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

// NewGetSiteAssetInfoParams creates a new GetSiteAssetInfoParams object
// with the default values initialized.
func NewGetSiteAssetInfoParams() *GetSiteAssetInfoParams {
	var ()
	return &GetSiteAssetInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSiteAssetInfoParamsWithTimeout creates a new GetSiteAssetInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSiteAssetInfoParamsWithTimeout(timeout time.Duration) *GetSiteAssetInfoParams {
	var ()
	return &GetSiteAssetInfoParams{

		timeout: timeout,
	}
}

// NewGetSiteAssetInfoParamsWithContext creates a new GetSiteAssetInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSiteAssetInfoParamsWithContext(ctx context.Context) *GetSiteAssetInfoParams {
	var ()
	return &GetSiteAssetInfoParams{

		Context: ctx,
	}
}

// NewGetSiteAssetInfoParamsWithHTTPClient creates a new GetSiteAssetInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSiteAssetInfoParamsWithHTTPClient(client *http.Client) *GetSiteAssetInfoParams {
	var ()
	return &GetSiteAssetInfoParams{
		HTTPClient: client,
	}
}

/*GetSiteAssetInfoParams contains all the parameters to send to the API endpoint
for the get site asset info operation typically these are written to a http.Request
*/
type GetSiteAssetInfoParams struct {

	/*AssetID*/
	AssetID string
	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get site asset info params
func (o *GetSiteAssetInfoParams) WithTimeout(timeout time.Duration) *GetSiteAssetInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get site asset info params
func (o *GetSiteAssetInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get site asset info params
func (o *GetSiteAssetInfoParams) WithContext(ctx context.Context) *GetSiteAssetInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get site asset info params
func (o *GetSiteAssetInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get site asset info params
func (o *GetSiteAssetInfoParams) WithHTTPClient(client *http.Client) *GetSiteAssetInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get site asset info params
func (o *GetSiteAssetInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssetID adds the assetID to the get site asset info params
func (o *GetSiteAssetInfoParams) WithAssetID(assetID string) *GetSiteAssetInfoParams {
	o.SetAssetID(assetID)
	return o
}

// SetAssetID adds the assetId to the get site asset info params
func (o *GetSiteAssetInfoParams) SetAssetID(assetID string) {
	o.AssetID = assetID
}

// WithSiteID adds the siteID to the get site asset info params
func (o *GetSiteAssetInfoParams) WithSiteID(siteID string) *GetSiteAssetInfoParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get site asset info params
func (o *GetSiteAssetInfoParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSiteAssetInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param asset_id
	if err := r.SetPathParam("asset_id", o.AssetID); err != nil {
		return err
	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
