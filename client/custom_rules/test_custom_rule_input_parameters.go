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

// NewTestCustomRuleInputParams creates a new TestCustomRuleInputParams object
// with the default values initialized.
func NewTestCustomRuleInputParams() *TestCustomRuleInputParams {
	var ()
	return &TestCustomRuleInputParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTestCustomRuleInputParamsWithTimeout creates a new TestCustomRuleInputParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTestCustomRuleInputParamsWithTimeout(timeout time.Duration) *TestCustomRuleInputParams {
	var ()
	return &TestCustomRuleInputParams{

		timeout: timeout,
	}
}

// NewTestCustomRuleInputParamsWithContext creates a new TestCustomRuleInputParams object
// with the default values initialized, and the ability to set a context for a request
func NewTestCustomRuleInputParamsWithContext(ctx context.Context) *TestCustomRuleInputParams {
	var ()
	return &TestCustomRuleInputParams{

		Context: ctx,
	}
}

// NewTestCustomRuleInputParamsWithHTTPClient creates a new TestCustomRuleInputParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTestCustomRuleInputParamsWithHTTPClient(client *http.Client) *TestCustomRuleInputParams {
	var ()
	return &TestCustomRuleInputParams{
		HTTPClient: client,
	}
}

/*TestCustomRuleInputParams contains all the parameters to send to the API endpoint
for the test custom rule input operation typically these are written to a http.Request
*/
type TestCustomRuleInputParams struct {

	/*ScanID
	  Scan of which we should get the custom rule test input.

	*/
	ScanID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the test custom rule input params
func (o *TestCustomRuleInputParams) WithTimeout(timeout time.Duration) *TestCustomRuleInputParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test custom rule input params
func (o *TestCustomRuleInputParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test custom rule input params
func (o *TestCustomRuleInputParams) WithContext(ctx context.Context) *TestCustomRuleInputParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test custom rule input params
func (o *TestCustomRuleInputParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test custom rule input params
func (o *TestCustomRuleInputParams) WithHTTPClient(client *http.Client) *TestCustomRuleInputParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test custom rule input params
func (o *TestCustomRuleInputParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScanID adds the scanID to the test custom rule input params
func (o *TestCustomRuleInputParams) WithScanID(scanID string) *TestCustomRuleInputParams {
	o.SetScanID(scanID)
	return o
}

// SetScanID adds the scanId to the test custom rule input params
func (o *TestCustomRuleInputParams) SetScanID(scanID string) {
	o.ScanID = scanID
}

// WriteToRequest writes these params to a swagger request
func (o *TestCustomRuleInputParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param scan_id
	qrScanID := o.ScanID
	qScanID := qrScanID
	if qScanID != "" {
		if err := r.SetQueryParam("scan_id", qScanID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
