// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListEventsParams creates a new ListEventsParams object
// with the default values initialized.
func NewListEventsParams() *ListEventsParams {
	var (
		maxItemsDefault = int64(100)
		offsetDefault   = int64(0)
	)
	return &ListEventsParams{
		MaxItems: &maxItemsDefault,
		Offset:   &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListEventsParamsWithTimeout creates a new ListEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListEventsParamsWithTimeout(timeout time.Duration) *ListEventsParams {
	var (
		maxItemsDefault = int64(100)
		offsetDefault   = int64(0)
	)
	return &ListEventsParams{
		MaxItems: &maxItemsDefault,
		Offset:   &offsetDefault,

		timeout: timeout,
	}
}

// NewListEventsParamsWithContext creates a new ListEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListEventsParamsWithContext(ctx context.Context) *ListEventsParams {
	var (
		maxItemsDefault = int64(100)
		offsetDefault   = int64(0)
	)
	return &ListEventsParams{
		MaxItems: &maxItemsDefault,
		Offset:   &offsetDefault,

		Context: ctx,
	}
}

// NewListEventsParamsWithHTTPClient creates a new ListEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListEventsParamsWithHTTPClient(client *http.Client) *ListEventsParams {
	var (
		maxItemsDefault = int64(100)
		offsetDefault   = int64(0)
	)
	return &ListEventsParams{
		MaxItems:   &maxItemsDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*ListEventsParams contains all the parameters to send to the API endpoint
for the list events operation typically these are written to a http.Request
*/
type ListEventsParams struct {

	/*Change
	  Type of change made in the event to filter by. When not specified, all change types will be returned.

	*/
	Change []string
	/*EnvironmentID
	  Environment ID.

	*/
	EnvironmentID string
	/*EventType
	  Event type to filter by. When not specified, all event types will be returned.

	*/
	EventType []string
	/*MaxItems
	  Maximum number of items to return.

	*/
	MaxItems *int64
	/*Offset
	  Number of items to skip before returning. This parameter is used when the number of items spans multiple pages.

	*/
	Offset *int64
	/*RangeFrom
	  Earliest created_at time to return events from.

	*/
	RangeFrom *int64
	/*RangeTo
	  Latest created_at time to return events from.

	*/
	RangeTo *int64
	/*Remediated
	  Filter remediation results for an event by success or failure. When not specified, all remediation results will be returned.

	*/
	Remediated []string
	/*ResourceType
	  Resource types in the event to filter by. When not specified, all resource types will be returned.

	*/
	ResourceType []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list events params
func (o *ListEventsParams) WithTimeout(timeout time.Duration) *ListEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list events params
func (o *ListEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list events params
func (o *ListEventsParams) WithContext(ctx context.Context) *ListEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list events params
func (o *ListEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list events params
func (o *ListEventsParams) WithHTTPClient(client *http.Client) *ListEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list events params
func (o *ListEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChange adds the change to the list events params
func (o *ListEventsParams) WithChange(change []string) *ListEventsParams {
	o.SetChange(change)
	return o
}

// SetChange adds the change to the list events params
func (o *ListEventsParams) SetChange(change []string) {
	o.Change = change
}

// WithEnvironmentID adds the environmentID to the list events params
func (o *ListEventsParams) WithEnvironmentID(environmentID string) *ListEventsParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the list events params
func (o *ListEventsParams) SetEnvironmentID(environmentID string) {
	o.EnvironmentID = environmentID
}

// WithEventType adds the eventType to the list events params
func (o *ListEventsParams) WithEventType(eventType []string) *ListEventsParams {
	o.SetEventType(eventType)
	return o
}

// SetEventType adds the eventType to the list events params
func (o *ListEventsParams) SetEventType(eventType []string) {
	o.EventType = eventType
}

// WithMaxItems adds the maxItems to the list events params
func (o *ListEventsParams) WithMaxItems(maxItems *int64) *ListEventsParams {
	o.SetMaxItems(maxItems)
	return o
}

// SetMaxItems adds the maxItems to the list events params
func (o *ListEventsParams) SetMaxItems(maxItems *int64) {
	o.MaxItems = maxItems
}

// WithOffset adds the offset to the list events params
func (o *ListEventsParams) WithOffset(offset *int64) *ListEventsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list events params
func (o *ListEventsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithRangeFrom adds the rangeFrom to the list events params
func (o *ListEventsParams) WithRangeFrom(rangeFrom *int64) *ListEventsParams {
	o.SetRangeFrom(rangeFrom)
	return o
}

// SetRangeFrom adds the rangeFrom to the list events params
func (o *ListEventsParams) SetRangeFrom(rangeFrom *int64) {
	o.RangeFrom = rangeFrom
}

// WithRangeTo adds the rangeTo to the list events params
func (o *ListEventsParams) WithRangeTo(rangeTo *int64) *ListEventsParams {
	o.SetRangeTo(rangeTo)
	return o
}

// SetRangeTo adds the rangeTo to the list events params
func (o *ListEventsParams) SetRangeTo(rangeTo *int64) {
	o.RangeTo = rangeTo
}

// WithRemediated adds the remediated to the list events params
func (o *ListEventsParams) WithRemediated(remediated []string) *ListEventsParams {
	o.SetRemediated(remediated)
	return o
}

// SetRemediated adds the remediated to the list events params
func (o *ListEventsParams) SetRemediated(remediated []string) {
	o.Remediated = remediated
}

// WithResourceType adds the resourceType to the list events params
func (o *ListEventsParams) WithResourceType(resourceType []string) *ListEventsParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the list events params
func (o *ListEventsParams) SetResourceType(resourceType []string) {
	o.ResourceType = resourceType
}

// WriteToRequest writes these params to a swagger request
func (o *ListEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesChange := o.Change

	joinedChange := swag.JoinByFormat(valuesChange, "multi")
	// query array param change
	if err := r.SetQueryParam("change", joinedChange...); err != nil {
		return err
	}

	// query param environment_id
	qrEnvironmentID := o.EnvironmentID
	qEnvironmentID := qrEnvironmentID
	if qEnvironmentID != "" {
		if err := r.SetQueryParam("environment_id", qEnvironmentID); err != nil {
			return err
		}
	}

	valuesEventType := o.EventType

	joinedEventType := swag.JoinByFormat(valuesEventType, "multi")
	// query array param event_type
	if err := r.SetQueryParam("event_type", joinedEventType...); err != nil {
		return err
	}

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

	if o.RangeFrom != nil {

		// query param range_from
		var qrRangeFrom int64
		if o.RangeFrom != nil {
			qrRangeFrom = *o.RangeFrom
		}
		qRangeFrom := swag.FormatInt64(qrRangeFrom)
		if qRangeFrom != "" {
			if err := r.SetQueryParam("range_from", qRangeFrom); err != nil {
				return err
			}
		}

	}

	if o.RangeTo != nil {

		// query param range_to
		var qrRangeTo int64
		if o.RangeTo != nil {
			qrRangeTo = *o.RangeTo
		}
		qRangeTo := swag.FormatInt64(qrRangeTo)
		if qRangeTo != "" {
			if err := r.SetQueryParam("range_to", qRangeTo); err != nil {
				return err
			}
		}

	}

	valuesRemediated := o.Remediated

	joinedRemediated := swag.JoinByFormat(valuesRemediated, "multi")
	// query array param remediated
	if err := r.SetQueryParam("remediated", joinedRemediated...); err != nil {
		return err
	}

	valuesResourceType := o.ResourceType

	joinedResourceType := swag.JoinByFormat(valuesResourceType, "multi")
	// query array param resource_type
	if err := r.SetQueryParam("resource_type", joinedResourceType...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
