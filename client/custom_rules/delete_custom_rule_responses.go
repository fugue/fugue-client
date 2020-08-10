// Code generated by go-swagger; DO NOT EDIT.

package custom_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fugue/fugue-client/models"
)

// DeleteCustomRuleReader is a Reader for the DeleteCustomRule structure.
type DeleteCustomRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCustomRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCustomRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCustomRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteCustomRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCustomRuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCustomRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCustomRuleNoContent creates a DeleteCustomRuleNoContent with default headers values
func NewDeleteCustomRuleNoContent() *DeleteCustomRuleNoContent {
	return &DeleteCustomRuleNoContent{}
}

/*DeleteCustomRuleNoContent handles this case with default header values.

Custom rule deleted.
*/
type DeleteCustomRuleNoContent struct {
}

func (o *DeleteCustomRuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /rules/{rule_id}][%d] deleteCustomRuleNoContent ", 204)
}

func (o *DeleteCustomRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomRuleBadRequest creates a DeleteCustomRuleBadRequest with default headers values
func NewDeleteCustomRuleBadRequest() *DeleteCustomRuleBadRequest {
	return &DeleteCustomRuleBadRequest{}
}

/*DeleteCustomRuleBadRequest handles this case with default header values.

Bad request error.
*/
type DeleteCustomRuleBadRequest struct {
	Payload *models.BadRequestError
}

func (o *DeleteCustomRuleBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /rules/{rule_id}][%d] deleteCustomRuleBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCustomRuleBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *DeleteCustomRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCustomRuleUnauthorized creates a DeleteCustomRuleUnauthorized with default headers values
func NewDeleteCustomRuleUnauthorized() *DeleteCustomRuleUnauthorized {
	return &DeleteCustomRuleUnauthorized{}
}

/*DeleteCustomRuleUnauthorized handles this case with default header values.

AuthenticationError
*/
type DeleteCustomRuleUnauthorized struct {
	Payload *models.AuthenticationError
}

func (o *DeleteCustomRuleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /rules/{rule_id}][%d] deleteCustomRuleUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteCustomRuleUnauthorized) GetPayload() *models.AuthenticationError {
	return o.Payload
}

func (o *DeleteCustomRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCustomRuleForbidden creates a DeleteCustomRuleForbidden with default headers values
func NewDeleteCustomRuleForbidden() *DeleteCustomRuleForbidden {
	return &DeleteCustomRuleForbidden{}
}

/*DeleteCustomRuleForbidden handles this case with default header values.

AuthorizationError
*/
type DeleteCustomRuleForbidden struct {
	Payload *models.AuthorizationError
}

func (o *DeleteCustomRuleForbidden) Error() string {
	return fmt.Sprintf("[DELETE /rules/{rule_id}][%d] deleteCustomRuleForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCustomRuleForbidden) GetPayload() *models.AuthorizationError {
	return o.Payload
}

func (o *DeleteCustomRuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCustomRuleInternalServerError creates a DeleteCustomRuleInternalServerError with default headers values
func NewDeleteCustomRuleInternalServerError() *DeleteCustomRuleInternalServerError {
	return &DeleteCustomRuleInternalServerError{}
}

/*DeleteCustomRuleInternalServerError handles this case with default header values.

InternalServerError
*/
type DeleteCustomRuleInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *DeleteCustomRuleInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /rules/{rule_id}][%d] deleteCustomRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCustomRuleInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *DeleteCustomRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
