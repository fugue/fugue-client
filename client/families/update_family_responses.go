// Code generated by go-swagger; DO NOT EDIT.

package families

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fugue/fugue-client/models"
)

// UpdateFamilyReader is a Reader for the UpdateFamily structure.
type UpdateFamilyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFamilyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateFamilyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateFamilyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateFamilyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateFamilyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateFamilyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateFamilyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateFamilyOK creates a UpdateFamilyOK with default headers values
func NewUpdateFamilyOK() *UpdateFamilyOK {
	return &UpdateFamilyOK{}
}

/*UpdateFamilyOK handles this case with default header values.

The updated Family.
*/
type UpdateFamilyOK struct {
	Payload *models.Family
}

func (o *UpdateFamilyOK) Error() string {
	return fmt.Sprintf("[PATCH /families/{family_id}][%d] updateFamilyOK  %+v", 200, o.Payload)
}

func (o *UpdateFamilyOK) GetPayload() *models.Family {
	return o.Payload
}

func (o *UpdateFamilyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Family)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFamilyBadRequest creates a UpdateFamilyBadRequest with default headers values
func NewUpdateFamilyBadRequest() *UpdateFamilyBadRequest {
	return &UpdateFamilyBadRequest{}
}

/*UpdateFamilyBadRequest handles this case with default header values.

BadRequestError
*/
type UpdateFamilyBadRequest struct {
	Payload *models.BadRequestError
}

func (o *UpdateFamilyBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /families/{family_id}][%d] updateFamilyBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateFamilyBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *UpdateFamilyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFamilyUnauthorized creates a UpdateFamilyUnauthorized with default headers values
func NewUpdateFamilyUnauthorized() *UpdateFamilyUnauthorized {
	return &UpdateFamilyUnauthorized{}
}

/*UpdateFamilyUnauthorized handles this case with default header values.

AuthenticationError
*/
type UpdateFamilyUnauthorized struct {
	Payload *models.AuthenticationError
}

func (o *UpdateFamilyUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /families/{family_id}][%d] updateFamilyUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateFamilyUnauthorized) GetPayload() *models.AuthenticationError {
	return o.Payload
}

func (o *UpdateFamilyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFamilyForbidden creates a UpdateFamilyForbidden with default headers values
func NewUpdateFamilyForbidden() *UpdateFamilyForbidden {
	return &UpdateFamilyForbidden{}
}

/*UpdateFamilyForbidden handles this case with default header values.

AuthorizationError
*/
type UpdateFamilyForbidden struct {
	Payload *models.AuthorizationError
}

func (o *UpdateFamilyForbidden) Error() string {
	return fmt.Sprintf("[PATCH /families/{family_id}][%d] updateFamilyForbidden  %+v", 403, o.Payload)
}

func (o *UpdateFamilyForbidden) GetPayload() *models.AuthorizationError {
	return o.Payload
}

func (o *UpdateFamilyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFamilyNotFound creates a UpdateFamilyNotFound with default headers values
func NewUpdateFamilyNotFound() *UpdateFamilyNotFound {
	return &UpdateFamilyNotFound{}
}

/*UpdateFamilyNotFound handles this case with default header values.

NotFoundError
*/
type UpdateFamilyNotFound struct {
	Payload *models.NotFoundError
}

func (o *UpdateFamilyNotFound) Error() string {
	return fmt.Sprintf("[PATCH /families/{family_id}][%d] updateFamilyNotFound  %+v", 404, o.Payload)
}

func (o *UpdateFamilyNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *UpdateFamilyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFamilyInternalServerError creates a UpdateFamilyInternalServerError with default headers values
func NewUpdateFamilyInternalServerError() *UpdateFamilyInternalServerError {
	return &UpdateFamilyInternalServerError{}
}

/*UpdateFamilyInternalServerError handles this case with default header values.

InternalServerError
*/
type UpdateFamilyInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *UpdateFamilyInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /families/{family_id}][%d] updateFamilyInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateFamilyInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *UpdateFamilyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
