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
	/*QEnvironmentID
	  An environment ID associated with a rule waiver.

	*/
	QEnvironmentID []string
	/*QEnvironmentName
	  An environment name associated with a rule waiver.

	*/
	QEnvironmentName []string
	/*QEnvironmentProvider
	  An environment provider associated with a rule waiver

	*/
	QEnvironmentProvider []string
	/*QID
	  A specific rule waiver ID.

	*/
	QID []string
	/*QName
	  A name of a rule waiver.

	*/
	QName []string
	/*QProvider
	  Alias for q.environment_provider.

	*/
	QProvider []string
	/*QResourceID
	  A resource ID associated with a rule waiver.

	*/
	QResourceID []string
	/*QResourceProvider
	  A resource provider associated with a rule waiver

	*/
	QResourceProvider []string
	/*QResourceType
	  A resource ID associated with a rule waiver.

	*/
	QResourceType []string
	/*QRuleID
	  A rule ID associated with a rule waiver.

	*/
	QRuleID []string
	/*QStatus
	  A current waiver status.

	*/
	QStatus []string
	/*Query
	  Deprecated, use the q.<parameter> fields instead.

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

// WithQEnvironmentID adds the qEnvironmentID to the list rule waivers params
func (o *ListRuleWaiversParams) WithQEnvironmentID(qEnvironmentID []string) *ListRuleWaiversParams {
	o.SetQEnvironmentID(qEnvironmentID)
	return o
}

// SetQEnvironmentID adds the qEnvironmentId to the list rule waivers params
func (o *ListRuleWaiversParams) SetQEnvironmentID(qEnvironmentID []string) {
	o.QEnvironmentID = qEnvironmentID
}

// WithQEnvironmentName adds the qEnvironmentName to the list rule waivers params
func (o *ListRuleWaiversParams) WithQEnvironmentName(qEnvironmentName []string) *ListRuleWaiversParams {
	o.SetQEnvironmentName(qEnvironmentName)
	return o
}

// SetQEnvironmentName adds the qEnvironmentName to the list rule waivers params
func (o *ListRuleWaiversParams) SetQEnvironmentName(qEnvironmentName []string) {
	o.QEnvironmentName = qEnvironmentName
}

// WithQEnvironmentProvider adds the qEnvironmentProvider to the list rule waivers params
func (o *ListRuleWaiversParams) WithQEnvironmentProvider(qEnvironmentProvider []string) *ListRuleWaiversParams {
	o.SetQEnvironmentProvider(qEnvironmentProvider)
	return o
}

// SetQEnvironmentProvider adds the qEnvironmentProvider to the list rule waivers params
func (o *ListRuleWaiversParams) SetQEnvironmentProvider(qEnvironmentProvider []string) {
	o.QEnvironmentProvider = qEnvironmentProvider
}

// WithQID adds the qID to the list rule waivers params
func (o *ListRuleWaiversParams) WithQID(qID []string) *ListRuleWaiversParams {
	o.SetQID(qID)
	return o
}

// SetQID adds the qId to the list rule waivers params
func (o *ListRuleWaiversParams) SetQID(qID []string) {
	o.QID = qID
}

// WithQName adds the qName to the list rule waivers params
func (o *ListRuleWaiversParams) WithQName(qName []string) *ListRuleWaiversParams {
	o.SetQName(qName)
	return o
}

// SetQName adds the qName to the list rule waivers params
func (o *ListRuleWaiversParams) SetQName(qName []string) {
	o.QName = qName
}

// WithQProvider adds the qProvider to the list rule waivers params
func (o *ListRuleWaiversParams) WithQProvider(qProvider []string) *ListRuleWaiversParams {
	o.SetQProvider(qProvider)
	return o
}

// SetQProvider adds the qProvider to the list rule waivers params
func (o *ListRuleWaiversParams) SetQProvider(qProvider []string) {
	o.QProvider = qProvider
}

// WithQResourceID adds the qResourceID to the list rule waivers params
func (o *ListRuleWaiversParams) WithQResourceID(qResourceID []string) *ListRuleWaiversParams {
	o.SetQResourceID(qResourceID)
	return o
}

// SetQResourceID adds the qResourceId to the list rule waivers params
func (o *ListRuleWaiversParams) SetQResourceID(qResourceID []string) {
	o.QResourceID = qResourceID
}

// WithQResourceProvider adds the qResourceProvider to the list rule waivers params
func (o *ListRuleWaiversParams) WithQResourceProvider(qResourceProvider []string) *ListRuleWaiversParams {
	o.SetQResourceProvider(qResourceProvider)
	return o
}

// SetQResourceProvider adds the qResourceProvider to the list rule waivers params
func (o *ListRuleWaiversParams) SetQResourceProvider(qResourceProvider []string) {
	o.QResourceProvider = qResourceProvider
}

// WithQResourceType adds the qResourceType to the list rule waivers params
func (o *ListRuleWaiversParams) WithQResourceType(qResourceType []string) *ListRuleWaiversParams {
	o.SetQResourceType(qResourceType)
	return o
}

// SetQResourceType adds the qResourceType to the list rule waivers params
func (o *ListRuleWaiversParams) SetQResourceType(qResourceType []string) {
	o.QResourceType = qResourceType
}

// WithQRuleID adds the qRuleID to the list rule waivers params
func (o *ListRuleWaiversParams) WithQRuleID(qRuleID []string) *ListRuleWaiversParams {
	o.SetQRuleID(qRuleID)
	return o
}

// SetQRuleID adds the qRuleId to the list rule waivers params
func (o *ListRuleWaiversParams) SetQRuleID(qRuleID []string) {
	o.QRuleID = qRuleID
}

// WithQStatus adds the qStatus to the list rule waivers params
func (o *ListRuleWaiversParams) WithQStatus(qStatus []string) *ListRuleWaiversParams {
	o.SetQStatus(qStatus)
	return o
}

// SetQStatus adds the qStatus to the list rule waivers params
func (o *ListRuleWaiversParams) SetQStatus(qStatus []string) {
	o.QStatus = qStatus
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

	valuesQEnvironmentID := o.QEnvironmentID

	joinedQEnvironmentID := swag.JoinByFormat(valuesQEnvironmentID, "multi")
	// query array param q.environment_id
	if err := r.SetQueryParam("q.environment_id", joinedQEnvironmentID...); err != nil {
		return err
	}

	valuesQEnvironmentName := o.QEnvironmentName

	joinedQEnvironmentName := swag.JoinByFormat(valuesQEnvironmentName, "multi")
	// query array param q.environment_name
	if err := r.SetQueryParam("q.environment_name", joinedQEnvironmentName...); err != nil {
		return err
	}

	valuesQEnvironmentProvider := o.QEnvironmentProvider

	joinedQEnvironmentProvider := swag.JoinByFormat(valuesQEnvironmentProvider, "multi")
	// query array param q.environment_provider
	if err := r.SetQueryParam("q.environment_provider", joinedQEnvironmentProvider...); err != nil {
		return err
	}

	valuesQID := o.QID

	joinedQID := swag.JoinByFormat(valuesQID, "multi")
	// query array param q.id
	if err := r.SetQueryParam("q.id", joinedQID...); err != nil {
		return err
	}

	valuesQName := o.QName

	joinedQName := swag.JoinByFormat(valuesQName, "multi")
	// query array param q.name
	if err := r.SetQueryParam("q.name", joinedQName...); err != nil {
		return err
	}

	valuesQProvider := o.QProvider

	joinedQProvider := swag.JoinByFormat(valuesQProvider, "multi")
	// query array param q.provider
	if err := r.SetQueryParam("q.provider", joinedQProvider...); err != nil {
		return err
	}

	valuesQResourceID := o.QResourceID

	joinedQResourceID := swag.JoinByFormat(valuesQResourceID, "multi")
	// query array param q.resource_id
	if err := r.SetQueryParam("q.resource_id", joinedQResourceID...); err != nil {
		return err
	}

	valuesQResourceProvider := o.QResourceProvider

	joinedQResourceProvider := swag.JoinByFormat(valuesQResourceProvider, "multi")
	// query array param q.resource_provider
	if err := r.SetQueryParam("q.resource_provider", joinedQResourceProvider...); err != nil {
		return err
	}

	valuesQResourceType := o.QResourceType

	joinedQResourceType := swag.JoinByFormat(valuesQResourceType, "multi")
	// query array param q.resource_type
	if err := r.SetQueryParam("q.resource_type", joinedQResourceType...); err != nil {
		return err
	}

	valuesQRuleID := o.QRuleID

	joinedQRuleID := swag.JoinByFormat(valuesQRuleID, "multi")
	// query array param q.rule_id
	if err := r.SetQueryParam("q.rule_id", joinedQRuleID...); err != nil {
		return err
	}

	valuesQStatus := o.QStatus

	joinedQStatus := swag.JoinByFormat(valuesQStatus, "multi")
	// query array param q.status
	if err := r.SetQueryParam("q.status", joinedQStatus...); err != nil {
		return err
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
