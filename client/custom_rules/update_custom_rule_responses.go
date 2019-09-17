// Code generated by go-swagger; DO NOT EDIT.

package custom_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/fugue/fugue-client/models"
)

// UpdateCustomRuleReader is a Reader for the UpdateCustomRule structure.
type UpdateCustomRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCustomRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateCustomRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewUpdateCustomRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateCustomRuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateCustomRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateCustomRuleOK creates a UpdateCustomRuleOK with default headers values
func NewUpdateCustomRuleOK() *UpdateCustomRuleOK {
	return &UpdateCustomRuleOK{}
}

/*UpdateCustomRuleOK handles this case with default header values.

New custom rule details.
*/
type UpdateCustomRuleOK struct {
	Payload *models.CustomRule
}

func (o *UpdateCustomRuleOK) Error() string {
	return fmt.Sprintf("[PATCH /rules/{rule_id}][%d] updateCustomRuleOK  %+v", 200, o.Payload)
}

func (o *UpdateCustomRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomRuleUnauthorized creates a UpdateCustomRuleUnauthorized with default headers values
func NewUpdateCustomRuleUnauthorized() *UpdateCustomRuleUnauthorized {
	return &UpdateCustomRuleUnauthorized{}
}

/*UpdateCustomRuleUnauthorized handles this case with default header values.

AuthenticationError
*/
type UpdateCustomRuleUnauthorized struct {
	Payload *models.AuthenticationError
}

func (o *UpdateCustomRuleUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /rules/{rule_id}][%d] updateCustomRuleUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateCustomRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomRuleForbidden creates a UpdateCustomRuleForbidden with default headers values
func NewUpdateCustomRuleForbidden() *UpdateCustomRuleForbidden {
	return &UpdateCustomRuleForbidden{}
}

/*UpdateCustomRuleForbidden handles this case with default header values.

AuthorizationError
*/
type UpdateCustomRuleForbidden struct {
	Payload *models.AuthorizationError
}

func (o *UpdateCustomRuleForbidden) Error() string {
	return fmt.Sprintf("[PATCH /rules/{rule_id}][%d] updateCustomRuleForbidden  %+v", 403, o.Payload)
}

func (o *UpdateCustomRuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomRuleInternalServerError creates a UpdateCustomRuleInternalServerError with default headers values
func NewUpdateCustomRuleInternalServerError() *UpdateCustomRuleInternalServerError {
	return &UpdateCustomRuleInternalServerError{}
}

/*UpdateCustomRuleInternalServerError handles this case with default header values.

InternalServerError
*/
type UpdateCustomRuleInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *UpdateCustomRuleInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /rules/{rule_id}][%d] updateCustomRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateCustomRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
