// Code generated by go-swagger; DO NOT EDIT.

package c_o_r_s

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

// NewOptionsRulesRuleIDParams creates a new OptionsRulesRuleIDParams object
// with the default values initialized.
func NewOptionsRulesRuleIDParams() *OptionsRulesRuleIDParams {
	var ()
	return &OptionsRulesRuleIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOptionsRulesRuleIDParamsWithTimeout creates a new OptionsRulesRuleIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOptionsRulesRuleIDParamsWithTimeout(timeout time.Duration) *OptionsRulesRuleIDParams {
	var ()
	return &OptionsRulesRuleIDParams{

		timeout: timeout,
	}
}

// NewOptionsRulesRuleIDParamsWithContext creates a new OptionsRulesRuleIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewOptionsRulesRuleIDParamsWithContext(ctx context.Context) *OptionsRulesRuleIDParams {
	var ()
	return &OptionsRulesRuleIDParams{

		Context: ctx,
	}
}

// NewOptionsRulesRuleIDParamsWithHTTPClient creates a new OptionsRulesRuleIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOptionsRulesRuleIDParamsWithHTTPClient(client *http.Client) *OptionsRulesRuleIDParams {
	var ()
	return &OptionsRulesRuleIDParams{
		HTTPClient: client,
	}
}

/*OptionsRulesRuleIDParams contains all the parameters to send to the API endpoint
for the options rules rule ID operation typically these are written to a http.Request
*/
type OptionsRulesRuleIDParams struct {

	/*RuleID
	  ID of the rule

	*/
	RuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the options rules rule ID params
func (o *OptionsRulesRuleIDParams) WithTimeout(timeout time.Duration) *OptionsRulesRuleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the options rules rule ID params
func (o *OptionsRulesRuleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the options rules rule ID params
func (o *OptionsRulesRuleIDParams) WithContext(ctx context.Context) *OptionsRulesRuleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the options rules rule ID params
func (o *OptionsRulesRuleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the options rules rule ID params
func (o *OptionsRulesRuleIDParams) WithHTTPClient(client *http.Client) *OptionsRulesRuleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the options rules rule ID params
func (o *OptionsRulesRuleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRuleID adds the ruleID to the options rules rule ID params
func (o *OptionsRulesRuleIDParams) WithRuleID(ruleID string) *OptionsRulesRuleIDParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the options rules rule ID params
func (o *OptionsRulesRuleIDParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *OptionsRulesRuleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
