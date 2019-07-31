// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/fugue/fugue-client/models"
)

// CreatePolicyReader is a Reader for the CreatePolicy structure.
type CreatePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreatePolicyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreatePolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreatePolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreatePolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreatePolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePolicyCreated creates a CreatePolicyCreated with default headers values
func NewCreatePolicyCreated() *CreatePolicyCreated {
	return &CreatePolicyCreated{}
}

/*CreatePolicyCreated handles this case with default header values.

Permissions for surveying and remediating the specified resource types.
*/
type CreatePolicyCreated struct {
	Payload *models.Permissions
}

func (o *CreatePolicyCreated) Error() string {
	return fmt.Sprintf("[POST /metadata/{provider}/permissions][%d] createPolicyCreated  %+v", 201, o.Payload)
}

func (o *CreatePolicyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Permissions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePolicyBadRequest creates a CreatePolicyBadRequest with default headers values
func NewCreatePolicyBadRequest() *CreatePolicyBadRequest {
	return &CreatePolicyBadRequest{}
}

/*CreatePolicyBadRequest handles this case with default header values.

Bad request error.
*/
type CreatePolicyBadRequest struct {
	Payload *models.BadRequestError
}

func (o *CreatePolicyBadRequest) Error() string {
	return fmt.Sprintf("[POST /metadata/{provider}/permissions][%d] createPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *CreatePolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePolicyUnauthorized creates a CreatePolicyUnauthorized with default headers values
func NewCreatePolicyUnauthorized() *CreatePolicyUnauthorized {
	return &CreatePolicyUnauthorized{}
}

/*CreatePolicyUnauthorized handles this case with default header values.

Authentication error.
*/
type CreatePolicyUnauthorized struct {
	Payload *models.AuthenticationError
}

func (o *CreatePolicyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /metadata/{provider}/permissions][%d] createPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *CreatePolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePolicyForbidden creates a CreatePolicyForbidden with default headers values
func NewCreatePolicyForbidden() *CreatePolicyForbidden {
	return &CreatePolicyForbidden{}
}

/*CreatePolicyForbidden handles this case with default header values.

Authorization error.
*/
type CreatePolicyForbidden struct {
	Payload *models.AuthorizationError
}

func (o *CreatePolicyForbidden) Error() string {
	return fmt.Sprintf("[POST /metadata/{provider}/permissions][%d] createPolicyForbidden  %+v", 403, o.Payload)
}

func (o *CreatePolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePolicyInternalServerError creates a CreatePolicyInternalServerError with default headers values
func NewCreatePolicyInternalServerError() *CreatePolicyInternalServerError {
	return &CreatePolicyInternalServerError{}
}

/*CreatePolicyInternalServerError handles this case with default header values.

Internal server error.
*/
type CreatePolicyInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *CreatePolicyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /metadata/{provider}/permissions][%d] createPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *CreatePolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
