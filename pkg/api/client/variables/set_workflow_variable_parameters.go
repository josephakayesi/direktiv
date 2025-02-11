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

// NewSetWorkflowVariableParams creates a new SetWorkflowVariableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetWorkflowVariableParams() *SetWorkflowVariableParams {
	return &SetWorkflowVariableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetWorkflowVariableParamsWithTimeout creates a new SetWorkflowVariableParams object
// with the ability to set a timeout on a request.
func NewSetWorkflowVariableParamsWithTimeout(timeout time.Duration) *SetWorkflowVariableParams {
	return &SetWorkflowVariableParams{
		timeout: timeout,
	}
}

// NewSetWorkflowVariableParamsWithContext creates a new SetWorkflowVariableParams object
// with the ability to set a context for a request.
func NewSetWorkflowVariableParamsWithContext(ctx context.Context) *SetWorkflowVariableParams {
	return &SetWorkflowVariableParams{
		Context: ctx,
	}
}

// NewSetWorkflowVariableParamsWithHTTPClient creates a new SetWorkflowVariableParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetWorkflowVariableParamsWithHTTPClient(client *http.Client) *SetWorkflowVariableParams {
	return &SetWorkflowVariableParams{
		HTTPClient: client,
	}
}

/* SetWorkflowVariableParams contains all the parameters to send to the API endpoint
   for the set workflow variable operation.

   Typically these are written to a http.Request.
*/
type SetWorkflowVariableParams struct {

	/* Data.

	   Payload that contains variable data.
	*/
	Data string

	/* Namespace.

	   target namespace
	*/
	Namespace string

	/* Var.

	   target variable
	*/
	Var string

	/* Workflow.

	   path to target workflow
	*/
	Workflow string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set workflow variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetWorkflowVariableParams) WithDefaults() *SetWorkflowVariableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set workflow variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetWorkflowVariableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set workflow variable params
func (o *SetWorkflowVariableParams) WithTimeout(timeout time.Duration) *SetWorkflowVariableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set workflow variable params
func (o *SetWorkflowVariableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set workflow variable params
func (o *SetWorkflowVariableParams) WithContext(ctx context.Context) *SetWorkflowVariableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set workflow variable params
func (o *SetWorkflowVariableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set workflow variable params
func (o *SetWorkflowVariableParams) WithHTTPClient(client *http.Client) *SetWorkflowVariableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set workflow variable params
func (o *SetWorkflowVariableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the set workflow variable params
func (o *SetWorkflowVariableParams) WithData(data string) *SetWorkflowVariableParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the set workflow variable params
func (o *SetWorkflowVariableParams) SetData(data string) {
	o.Data = data
}

// WithNamespace adds the namespace to the set workflow variable params
func (o *SetWorkflowVariableParams) WithNamespace(namespace string) *SetWorkflowVariableParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the set workflow variable params
func (o *SetWorkflowVariableParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithVar adds the varVar to the set workflow variable params
func (o *SetWorkflowVariableParams) WithVar(varVar string) *SetWorkflowVariableParams {
	o.SetVar(varVar)
	return o
}

// SetVar adds the var to the set workflow variable params
func (o *SetWorkflowVariableParams) SetVar(varVar string) {
	o.Var = varVar
}

// WithWorkflow adds the workflow to the set workflow variable params
func (o *SetWorkflowVariableParams) WithWorkflow(workflow string) *SetWorkflowVariableParams {
	o.SetWorkflow(workflow)
	return o
}

// SetWorkflow adds the workflow to the set workflow variable params
func (o *SetWorkflowVariableParams) SetWorkflow(workflow string) {
	o.Workflow = workflow
}

// WriteToRequest writes these params to a swagger request
func (o *SetWorkflowVariableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// query param var
	qrVar := o.Var
	qVar := qrVar
	if qVar != "" {

		if err := r.SetQueryParam("var", qVar); err != nil {
			return err
		}
	}

	// path param workflow
	if err := r.SetPathParam("workflow", o.Workflow); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
