// Code generated by go-swagger; DO NOT EDIT.

package variables

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
)

// NewDeleteInstanceVariableParams creates a new DeleteInstanceVariableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteInstanceVariableParams() *DeleteInstanceVariableParams {
	return &DeleteInstanceVariableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInstanceVariableParamsWithTimeout creates a new DeleteInstanceVariableParams object
// with the ability to set a timeout on a request.
func NewDeleteInstanceVariableParamsWithTimeout(timeout time.Duration) *DeleteInstanceVariableParams {
	return &DeleteInstanceVariableParams{
		timeout: timeout,
	}
}

// NewDeleteInstanceVariableParamsWithContext creates a new DeleteInstanceVariableParams object
// with the ability to set a context for a request.
func NewDeleteInstanceVariableParamsWithContext(ctx context.Context) *DeleteInstanceVariableParams {
	return &DeleteInstanceVariableParams{
		Context: ctx,
	}
}

// NewDeleteInstanceVariableParamsWithHTTPClient creates a new DeleteInstanceVariableParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteInstanceVariableParamsWithHTTPClient(client *http.Client) *DeleteInstanceVariableParams {
	return &DeleteInstanceVariableParams{
		HTTPClient: client,
	}
}

/* DeleteInstanceVariableParams contains all the parameters to send to the API endpoint
   for the delete instance variable operation.

   Typically these are written to a http.Request.
*/
type DeleteInstanceVariableParams struct {

	/* Instance.

	   target instance
	*/
	Instance string

	/* Namespace.

	   target namespace
	*/
	Namespace string

	/* Variable.

	   target variable
	*/
	Variable string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete instance variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteInstanceVariableParams) WithDefaults() *DeleteInstanceVariableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete instance variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteInstanceVariableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete instance variable params
func (o *DeleteInstanceVariableParams) WithTimeout(timeout time.Duration) *DeleteInstanceVariableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete instance variable params
func (o *DeleteInstanceVariableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete instance variable params
func (o *DeleteInstanceVariableParams) WithContext(ctx context.Context) *DeleteInstanceVariableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete instance variable params
func (o *DeleteInstanceVariableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete instance variable params
func (o *DeleteInstanceVariableParams) WithHTTPClient(client *http.Client) *DeleteInstanceVariableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete instance variable params
func (o *DeleteInstanceVariableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstance adds the instance to the delete instance variable params
func (o *DeleteInstanceVariableParams) WithInstance(instance string) *DeleteInstanceVariableParams {
	o.SetInstance(instance)
	return o
}

// SetInstance adds the instance to the delete instance variable params
func (o *DeleteInstanceVariableParams) SetInstance(instance string) {
	o.Instance = instance
}

// WithNamespace adds the namespace to the delete instance variable params
func (o *DeleteInstanceVariableParams) WithNamespace(namespace string) *DeleteInstanceVariableParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete instance variable params
func (o *DeleteInstanceVariableParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithVariable adds the variable to the delete instance variable params
func (o *DeleteInstanceVariableParams) WithVariable(variable string) *DeleteInstanceVariableParams {
	o.SetVariable(variable)
	return o
}

// SetVariable adds the variable to the delete instance variable params
func (o *DeleteInstanceVariableParams) SetVariable(variable string) {
	o.Variable = variable
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInstanceVariableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param instance
	if err := r.SetPathParam("instance", o.Instance); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param variable
	if err := r.SetPathParam("variable", o.Variable); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
