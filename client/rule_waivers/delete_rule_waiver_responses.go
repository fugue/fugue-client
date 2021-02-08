// Code generated by go-swagger; DO NOT EDIT.

package rule_waivers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fugue/fugue-client/models"
)

// DeleteRuleWaiverReader is a Reader for the DeleteRuleWaiver structure.
type DeleteRuleWaiverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRuleWaiverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRuleWaiverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteRuleWaiverUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRuleWaiverForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRuleWaiverNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRuleWaiverInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRuleWaiverOK creates a DeleteRuleWaiverOK with default headers values
func NewDeleteRuleWaiverOK() *DeleteRuleWaiverOK {
	return &DeleteRuleWaiverOK{}
}

/*DeleteRuleWaiverOK handles this case with default header values.

Rule waiver deleted.
*/
type DeleteRuleWaiverOK struct {
}

func (o *DeleteRuleWaiverOK) Error() string {
	return fmt.Sprintf("[DELETE /rule_waivers/{rule_waiver_id}][%d] deleteRuleWaiverOK ", 200)
}

func (o *DeleteRuleWaiverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRuleWaiverUnauthorized creates a DeleteRuleWaiverUnauthorized with default headers values
func NewDeleteRuleWaiverUnauthorized() *DeleteRuleWaiverUnauthorized {
	return &DeleteRuleWaiverUnauthorized{}
}

/*DeleteRuleWaiverUnauthorized handles this case with default header values.

AuthenticationError
*/
type DeleteRuleWaiverUnauthorized struct {
	Payload *models.AuthenticationError
}

func (o *DeleteRuleWaiverUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /rule_waivers/{rule_waiver_id}][%d] deleteRuleWaiverUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRuleWaiverUnauthorized) GetPayload() *models.AuthenticationError {
	return o.Payload
}

func (o *DeleteRuleWaiverUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRuleWaiverForbidden creates a DeleteRuleWaiverForbidden with default headers values
func NewDeleteRuleWaiverForbidden() *DeleteRuleWaiverForbidden {
	return &DeleteRuleWaiverForbidden{}
}

/*DeleteRuleWaiverForbidden handles this case with default header values.

AuthorizationError
*/
type DeleteRuleWaiverForbidden struct {
	Payload *models.AuthorizationError
}

func (o *DeleteRuleWaiverForbidden) Error() string {
	return fmt.Sprintf("[DELETE /rule_waivers/{rule_waiver_id}][%d] deleteRuleWaiverForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRuleWaiverForbidden) GetPayload() *models.AuthorizationError {
	return o.Payload
}

func (o *DeleteRuleWaiverForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRuleWaiverNotFound creates a DeleteRuleWaiverNotFound with default headers values
func NewDeleteRuleWaiverNotFound() *DeleteRuleWaiverNotFound {
	return &DeleteRuleWaiverNotFound{}
}

/*DeleteRuleWaiverNotFound handles this case with default header values.

NotFoundError
*/
type DeleteRuleWaiverNotFound struct {
	Payload *models.NotFoundError
}

func (o *DeleteRuleWaiverNotFound) Error() string {
	return fmt.Sprintf("[DELETE /rule_waivers/{rule_waiver_id}][%d] deleteRuleWaiverNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRuleWaiverNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *DeleteRuleWaiverNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRuleWaiverInternalServerError creates a DeleteRuleWaiverInternalServerError with default headers values
func NewDeleteRuleWaiverInternalServerError() *DeleteRuleWaiverInternalServerError {
	return &DeleteRuleWaiverInternalServerError{}
}

/*DeleteRuleWaiverInternalServerError handles this case with default header values.

InternalServerError
*/
type DeleteRuleWaiverInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *DeleteRuleWaiverInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /rule_waivers/{rule_waiver_id}][%d] deleteRuleWaiverInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRuleWaiverInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *DeleteRuleWaiverInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
