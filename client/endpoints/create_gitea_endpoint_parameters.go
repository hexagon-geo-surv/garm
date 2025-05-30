// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	garm_params "github.com/cloudbase/garm/params"
)

// NewCreateGiteaEndpointParams creates a new CreateGiteaEndpointParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateGiteaEndpointParams() *CreateGiteaEndpointParams {
	return &CreateGiteaEndpointParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateGiteaEndpointParamsWithTimeout creates a new CreateGiteaEndpointParams object
// with the ability to set a timeout on a request.
func NewCreateGiteaEndpointParamsWithTimeout(timeout time.Duration) *CreateGiteaEndpointParams {
	return &CreateGiteaEndpointParams{
		timeout: timeout,
	}
}

// NewCreateGiteaEndpointParamsWithContext creates a new CreateGiteaEndpointParams object
// with the ability to set a context for a request.
func NewCreateGiteaEndpointParamsWithContext(ctx context.Context) *CreateGiteaEndpointParams {
	return &CreateGiteaEndpointParams{
		Context: ctx,
	}
}

// NewCreateGiteaEndpointParamsWithHTTPClient creates a new CreateGiteaEndpointParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateGiteaEndpointParamsWithHTTPClient(client *http.Client) *CreateGiteaEndpointParams {
	return &CreateGiteaEndpointParams{
		HTTPClient: client,
	}
}

/*
CreateGiteaEndpointParams contains all the parameters to send to the API endpoint

	for the create gitea endpoint operation.

	Typically these are written to a http.Request.
*/
type CreateGiteaEndpointParams struct {

	/* Body.

	   Parameters used when creating a Gitea endpoint.
	*/
	Body garm_params.CreateGiteaEndpointParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create gitea endpoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGiteaEndpointParams) WithDefaults() *CreateGiteaEndpointParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create gitea endpoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGiteaEndpointParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create gitea endpoint params
func (o *CreateGiteaEndpointParams) WithTimeout(timeout time.Duration) *CreateGiteaEndpointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create gitea endpoint params
func (o *CreateGiteaEndpointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create gitea endpoint params
func (o *CreateGiteaEndpointParams) WithContext(ctx context.Context) *CreateGiteaEndpointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create gitea endpoint params
func (o *CreateGiteaEndpointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create gitea endpoint params
func (o *CreateGiteaEndpointParams) WithHTTPClient(client *http.Client) *CreateGiteaEndpointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create gitea endpoint params
func (o *CreateGiteaEndpointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create gitea endpoint params
func (o *CreateGiteaEndpointParams) WithBody(body garm_params.CreateGiteaEndpointParams) *CreateGiteaEndpointParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create gitea endpoint params
func (o *CreateGiteaEndpointParams) SetBody(body garm_params.CreateGiteaEndpointParams) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGiteaEndpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
