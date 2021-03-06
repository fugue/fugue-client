// Code generated by go-swagger; DO NOT EDIT.

package rule_waivers

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
	"github.com/go-openapi/swag"
)

// NewListRuleWaiversParams creates a new ListRuleWaiversParams object
// with the default values initialized.
func NewListRuleWaiversParams() *ListRuleWaiversParams {
	var (
		maxItemsDefault       = int64(100)
		offsetDefault         = int64(0)
		orderByDefault        = string("name")
		orderDirectionDefault = string("asc")
	)
	return &ListRuleWaiversParams{
		MaxItems:       &maxItemsDefault,
		Offset:         &offsetDefault,
		OrderBy:        &orderByDefault,
		OrderDirection: &orderDirectionDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListRuleWaiversParamsWithTimeout creates a new ListRuleWaiversParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListRuleWaiversParamsWithTimeout(timeout time.Duration) *ListRuleWaiversParams {
	var (
		maxItemsDefault       = int64(100)
		offsetDefault         = int64(0)
		orderByDefault        = string("name")
		orderDirectionDefault = string("asc")
	)
	return &ListRuleWaiversParams{
		MaxItems:       &maxItemsDefault,
		Offset:         &offsetDefault,
		OrderBy:        &orderByDefault,
		OrderDirection: &orderDirectionDefault,

		timeout: timeout,
	}
}

// NewListRuleWaiversParamsWithContext creates a new ListRuleWaiversParams object
// with the default values initialized, and the ability to set a context for a request
func NewListRuleWaiversParamsWithContext(ctx context.Context) *ListRuleWaiversParams {
	var (
		maxItemsDefault       = int64(100)
		offsetDefault         = int64(0)
		orderByDefault        = string("name")
		orderDirectionDefault = string("asc")
	)
	return &ListRuleWaiversParams{
		MaxItems:       &maxItemsDefault,
		Offset:         &offsetDefault,
		OrderBy:        &orderByDefault,
		OrderDirection: &orderDirectionDefault,

		Context: ctx,
	}
}

// NewListRuleWaiversParamsWithHTTPClient creates a new ListRuleWaiversParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListRuleWaiversParamsWithHTTPClient(client *http.Client) *ListRuleWaiversParams {
	var (
		maxItemsDefault       = int64(100)
		offsetDefault         = int64(0)
		orderByDefault        = string("name")
		orderDirectionDefault = string("asc")
	)
	return &ListRuleWaiversParams{
		MaxItems:       &maxItemsDefault,
		Offset:         &offsetDefault,
		OrderBy:        &orderByDefault,
		OrderDirection: &orderDirectionDefault,
		HTTPClient:     client,
	}
}

/*ListRuleWaiversParams contains all the parameters to send to the API endpoint
for the list rule waivers operation typically these are written to a http.Request
*/
type ListRuleWaiversParams struct {

	/*MaxItems
	  Maximum number of items to return.

	*/
	MaxItems *int64
	/*Offset
	  Number of items to skip before returning. This parameter is used when the number of items spans multiple pages.

	*/
	Offset *int64
	/*OrderBy
	  Field to sort the items by.

	*/
	OrderBy *string
	/*OrderDirection
	  Direction to sort the items in.

	*/
	OrderDirection *string
	/*Query
	  A stringified JSON array of search parameters.

	*/
	Query *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list rule waivers params
func (o *ListRuleWaiversParams) WithTimeout(timeout time.Duration) *ListRuleWaiversParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list rule waivers params
func (o *ListRuleWaiversParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list rule waivers params
func (o *ListRuleWaiversParams) WithContext(ctx context.Context) *ListRuleWaiversParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list rule waivers params
func (o *ListRuleWaiversParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list rule waivers params
func (o *ListRuleWaiversParams) WithHTTPClient(client *http.Client) *ListRuleWaiversParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list rule waivers params
func (o *ListRuleWaiversParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMaxItems adds the maxItems to the list rule waivers params
func (o *ListRuleWaiversParams) WithMaxItems(maxItems *int64) *ListRuleWaiversParams {
	o.SetMaxItems(maxItems)
	return o
}

// SetMaxItems adds the maxItems to the list rule waivers params
func (o *ListRuleWaiversParams) SetMaxItems(maxItems *int64) {
	o.MaxItems = maxItems
}

// WithOffset adds the offset to the list rule waivers params
func (o *ListRuleWaiversParams) WithOffset(offset *int64) *ListRuleWaiversParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list rule waivers params
func (o *ListRuleWaiversParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrderBy adds the orderBy to the list rule waivers params
func (o *ListRuleWaiversParams) WithOrderBy(orderBy *string) *ListRuleWaiversParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the list rule waivers params
func (o *ListRuleWaiversParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithOrderDirection adds the orderDirection to the list rule waivers params
func (o *ListRuleWaiversParams) WithOrderDirection(orderDirection *string) *ListRuleWaiversParams {
	o.SetOrderDirection(orderDirection)
	return o
}

// SetOrderDirection adds the orderDirection to the list rule waivers params
func (o *ListRuleWaiversParams) SetOrderDirection(orderDirection *string) {
	o.OrderDirection = orderDirection
}

// WithQuery adds the query to the list rule waivers params
func (o *ListRuleWaiversParams) WithQuery(query *string) *ListRuleWaiversParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the list rule waivers params
func (o *ListRuleWaiversParams) SetQuery(query *string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *ListRuleWaiversParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MaxItems != nil {

		// query param max_items
		var qrMaxItems int64
		if o.MaxItems != nil {
			qrMaxItems = *o.MaxItems
		}
		qMaxItems := swag.FormatInt64(qrMaxItems)
		if qMaxItems != "" {
			if err := r.SetQueryParam("max_items", qMaxItems); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.OrderBy != nil {

		// query param order_by
		var qrOrderBy string
		if o.OrderBy != nil {
			qrOrderBy = *o.OrderBy
		}
		qOrderBy := qrOrderBy
		if qOrderBy != "" {
			if err := r.SetQueryParam("order_by", qOrderBy); err != nil {
				return err
			}
		}

	}

	if o.OrderDirection != nil {

		// query param order_direction
		var qrOrderDirection string
		if o.OrderDirection != nil {
			qrOrderDirection = *o.OrderDirection
		}
		qOrderDirection := qrOrderDirection
		if qOrderDirection != "" {
			if err := r.SetQueryParam("order_direction", qOrderDirection); err != nil {
				return err
			}
		}

	}

	if o.Query != nil {

		// query param query
		var qrQuery string
		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {
			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
