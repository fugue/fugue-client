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

// TestCustomRuleReader is a Reader for the TestCustomRule structure.
type TestCustomRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestCustomRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTestCustomRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTestCustomRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewTestCustomRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTestCustomRuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTestCustomRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTestCustomRuleOK creates a TestCustomRuleOK with default headers values
func NewTestCustomRuleOK() *TestCustomRuleOK {
	return &TestCustomRuleOK{}
}

/*TestCustomRuleOK handles this case with default header values.

Validation results for the custom rule.
*/
type TestCustomRuleOK struct {
	Payload *models.TestCustomRuleOutput
}

func (o *TestCustomRuleOK) Error() string {
	return fmt.Sprintf("[POST /rules/test][%d] testCustomRuleOK  %+v", 200, o.Payload)
}

func (o *TestCustomRuleOK) GetPayload() *models.TestCustomRuleOutput {
	return o.Payload
}

func (o *TestCustomRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TestCustomRuleOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestCustomRuleBadRequest creates a TestCustomRuleBadRequest with default headers values
func NewTestCustomRuleBadRequest() *TestCustomRuleBadRequest {
	return &TestCustomRuleBadRequest{}
}

/*TestCustomRuleBadRequest handles this case with default header values.

Bad request error.
*/
type TestCustomRuleBadRequest struct {
	Payload *models.BadRequestError
}

func (o *TestCustomRuleBadRequest) Error() string {
	return fmt.Sprintf("[POST /rules/test][%d] testCustomRuleBadRequest  %+v", 400, o.Payload)
}

func (o *TestCustomRuleBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *TestCustomRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestCustomRuleUnauthorized creates a TestCustomRuleUnauthorized with default headers values
func NewTestCustomRuleUnauthorized() *TestCustomRuleUnauthorized {
	return &TestCustomRuleUnauthorized{}
}

/*TestCustomRuleUnauthorized handles this case with default header values.

AuthenticationError
*/
type TestCustomRuleUnauthorized struct {
	Payload *models.AuthenticationError
}

func (o *TestCustomRuleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /rules/test][%d] testCustomRuleUnauthorized  %+v", 401, o.Payload)
}

func (o *TestCustomRuleUnauthorized) GetPayload() *models.AuthenticationError {
	return o.Payload
}

func (o *TestCustomRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestCustomRuleForbidden creates a TestCustomRuleForbidden with default headers values
func NewTestCustomRuleForbidden() *TestCustomRuleForbidden {
	return &TestCustomRuleForbidden{}
}

/*TestCustomRuleForbidden handles this case with default header values.

AuthorizationError
*/
type TestCustomRuleForbidden struct {
	Payload *models.AuthorizationError
}

func (o *TestCustomRuleForbidden) Error() string {
	return fmt.Sprintf("[POST /rules/test][%d] testCustomRuleForbidden  %+v", 403, o.Payload)
}

func (o *TestCustomRuleForbidden) GetPayload() *models.AuthorizationError {
	return o.Payload
}

func (o *TestCustomRuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestCustomRuleInternalServerError creates a TestCustomRuleInternalServerError with default headers values
func NewTestCustomRuleInternalServerError() *TestCustomRuleInternalServerError {
	return &TestCustomRuleInternalServerError{}
}

/*TestCustomRuleInternalServerError handles this case with default header values.

InternalServerError
*/
type TestCustomRuleInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *TestCustomRuleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /rules/test][%d] testCustomRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *TestCustomRuleInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *TestCustomRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
