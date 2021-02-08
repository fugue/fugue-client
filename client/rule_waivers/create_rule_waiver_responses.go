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

// CreateRuleWaiverReader is a Reader for the CreateRuleWaiver structure.
type CreateRuleWaiverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRuleWaiverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRuleWaiverCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRuleWaiverBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateRuleWaiverUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRuleWaiverForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRuleWaiverInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateRuleWaiverCreated creates a CreateRuleWaiverCreated with default headers values
func NewCreateRuleWaiverCreated() *CreateRuleWaiverCreated {
	return &CreateRuleWaiverCreated{}
}

/*CreateRuleWaiverCreated handles this case with default header values.

New rule waiver details.
*/
type CreateRuleWaiverCreated struct {
	Payload *models.RuleWaivers
}

func (o *CreateRuleWaiverCreated) Error() string {
	return fmt.Sprintf("[POST /rule_waivers][%d] createRuleWaiverCreated  %+v", 201, o.Payload)
}

func (o *CreateRuleWaiverCreated) GetPayload() *models.RuleWaivers {
	return o.Payload
}

func (o *CreateRuleWaiverCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuleWaivers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRuleWaiverBadRequest creates a CreateRuleWaiverBadRequest with default headers values
func NewCreateRuleWaiverBadRequest() *CreateRuleWaiverBadRequest {
	return &CreateRuleWaiverBadRequest{}
}

/*CreateRuleWaiverBadRequest handles this case with default header values.

BadRequestError
*/
type CreateRuleWaiverBadRequest struct {
	Payload *models.BadRequestError
}

func (o *CreateRuleWaiverBadRequest) Error() string {
	return fmt.Sprintf("[POST /rule_waivers][%d] createRuleWaiverBadRequest  %+v", 400, o.Payload)
}

func (o *CreateRuleWaiverBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *CreateRuleWaiverBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRuleWaiverUnauthorized creates a CreateRuleWaiverUnauthorized with default headers values
func NewCreateRuleWaiverUnauthorized() *CreateRuleWaiverUnauthorized {
	return &CreateRuleWaiverUnauthorized{}
}

/*CreateRuleWaiverUnauthorized handles this case with default header values.

AuthenticationError
*/
type CreateRuleWaiverUnauthorized struct {
	Payload *models.AuthenticationError
}

func (o *CreateRuleWaiverUnauthorized) Error() string {
	return fmt.Sprintf("[POST /rule_waivers][%d] createRuleWaiverUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateRuleWaiverUnauthorized) GetPayload() *models.AuthenticationError {
	return o.Payload
}

func (o *CreateRuleWaiverUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRuleWaiverForbidden creates a CreateRuleWaiverForbidden with default headers values
func NewCreateRuleWaiverForbidden() *CreateRuleWaiverForbidden {
	return &CreateRuleWaiverForbidden{}
}

/*CreateRuleWaiverForbidden handles this case with default header values.

AuthorizationError
*/
type CreateRuleWaiverForbidden struct {
	Payload *models.AuthorizationError
}

func (o *CreateRuleWaiverForbidden) Error() string {
	return fmt.Sprintf("[POST /rule_waivers][%d] createRuleWaiverForbidden  %+v", 403, o.Payload)
}

func (o *CreateRuleWaiverForbidden) GetPayload() *models.AuthorizationError {
	return o.Payload
}

func (o *CreateRuleWaiverForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRuleWaiverInternalServerError creates a CreateRuleWaiverInternalServerError with default headers values
func NewCreateRuleWaiverInternalServerError() *CreateRuleWaiverInternalServerError {
	return &CreateRuleWaiverInternalServerError{}
}

/*CreateRuleWaiverInternalServerError handles this case with default header values.

InternalServerError
*/
type CreateRuleWaiverInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *CreateRuleWaiverInternalServerError) Error() string {
	return fmt.Sprintf("[POST /rule_waivers][%d] createRuleWaiverInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateRuleWaiverInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *CreateRuleWaiverInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
