// Code generated by go-swagger; DO NOT EDIT.

package custom_rules

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

// NewGetCustomRuleParams creates a new GetCustomRuleParams object
// with the default values initialized.
func NewGetCustomRuleParams() *GetCustomRuleParams {
	var ()
	return &GetCustomRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomRuleParamsWithTimeout creates a new GetCustomRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomRuleParamsWithTimeout(timeout time.Duration) *GetCustomRuleParams {
	var ()
	return &GetCustomRuleParams{

		timeout: timeout,
	}
}

// NewGetCustomRuleParamsWithContext creates a new GetCustomRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomRuleParamsWithContext(ctx context.Context) *GetCustomRuleParams {
	var ()
	return &GetCustomRuleParams{

		Context: ctx,
	}
}

// NewGetCustomRuleParamsWithHTTPClient creates a new GetCustomRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomRuleParamsWithHTTPClient(client *http.Client) *GetCustomRuleParams {
	var ()
	return &GetCustomRuleParams{
		HTTPClient: client,
	}
}

/*GetCustomRuleParams contains all the parameters to send to the API endpoint
for the get custom rule operation typically these are written to a http.Request
*/
type GetCustomRuleParams struct {

	/*RuleID
	  The ID of the Rule to get.

	*/
	RuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get custom rule params
func (o *GetCustomRuleParams) WithTimeout(timeout time.Duration) *GetCustomRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get custom rule params
func (o *GetCustomRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get custom rule params
func (o *GetCustomRuleParams) WithContext(ctx context.Context) *GetCustomRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get custom rule params
func (o *GetCustomRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get custom rule params
func (o *GetCustomRuleParams) WithHTTPClient(client *http.Client) *GetCustomRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get custom rule params
func (o *GetCustomRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRuleID adds the ruleID to the get custom rule params
func (o *GetCustomRuleParams) WithRuleID(ruleID string) *GetCustomRuleParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the get custom rule params
func (o *GetCustomRuleParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param rule_id
	if err := r.SetPathParam("rule_id", o.RuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
