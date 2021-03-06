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
	"github.com/go-openapi/strfmt"

	"github.com/Meat-Hook/point-bank/internal/modules/user/internal/api/web/generated/models"
)

// NewUpdatePasswordParams creates a new UpdatePasswordParams object
// with the default values initialized.
func NewUpdatePasswordParams() *UpdatePasswordParams {
	var ()
	return &UpdatePasswordParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePasswordParamsWithTimeout creates a new UpdatePasswordParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePasswordParamsWithTimeout(timeout time.Duration) *UpdatePasswordParams {
	var ()
	return &UpdatePasswordParams{

		timeout: timeout,
	}
}

// NewUpdatePasswordParamsWithContext creates a new UpdatePasswordParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePasswordParamsWithContext(ctx context.Context) *UpdatePasswordParams {
	var ()
	return &UpdatePasswordParams{

		Context: ctx,
	}
}

// NewUpdatePasswordParamsWithHTTPClient creates a new UpdatePasswordParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePasswordParamsWithHTTPClient(client *http.Client) *UpdatePasswordParams {
	var ()
	return &UpdatePasswordParams{
		HTTPClient: client,
	}
}

/*UpdatePasswordParams contains all the parameters to send to the API endpoint
for the update password operation typically these are written to a http.Request
*/
type UpdatePasswordParams struct {

	/*Args*/
	Args *models.UpdatePassword

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update password params
func (o *UpdatePasswordParams) WithTimeout(timeout time.Duration) *UpdatePasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update password params
func (o *UpdatePasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update password params
func (o *UpdatePasswordParams) WithContext(ctx context.Context) *UpdatePasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update password params
func (o *UpdatePasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update password params
func (o *UpdatePasswordParams) WithHTTPClient(client *http.Client) *UpdatePasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update password params
func (o *UpdatePasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArgs adds the args to the update password params
func (o *UpdatePasswordParams) WithArgs(args *models.UpdatePassword) *UpdatePasswordParams {
	o.SetArgs(args)
	return o
}

// SetArgs adds the args to the update password params
func (o *UpdatePasswordParams) SetArgs(args *models.UpdatePassword) {
	o.Args = args
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Args != nil {
		if err := r.SetBodyParam(o.Args); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
