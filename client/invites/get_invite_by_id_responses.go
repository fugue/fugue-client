// Code generated by go-swagger; DO NOT EDIT.

package invites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fugue/fugue-client/models"
)

// GetInviteByIDReader is a Reader for the GetInviteByID structure.
type GetInviteByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInviteByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInviteByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInviteByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInviteByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInviteByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInviteByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInviteByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInviteByIDOK creates a GetInviteByIDOK with default headers values
func NewGetInviteByIDOK() *GetInviteByIDOK {
	return &GetInviteByIDOK{}
}

/*GetInviteByIDOK handles this case with default header values.

Invite details.
*/
type GetInviteByIDOK struct {
	Payload *models.Invite
}

func (o *GetInviteByIDOK) Error() string {
	return fmt.Sprintf("[GET /invites/{invite_id}][%d] getInviteByIdOK  %+v", 200, o.Payload)
}

func (o *GetInviteByIDOK) GetPayload() *models.Invite {
	return o.Payload
}

func (o *GetInviteByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Invite)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInviteByIDBadRequest creates a GetInviteByIDBadRequest with default headers values
func NewGetInviteByIDBadRequest() *GetInviteByIDBadRequest {
	return &GetInviteByIDBadRequest{}
}

/*GetInviteByIDBadRequest handles this case with default header values.

BadRequestError
*/
type GetInviteByIDBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetInviteByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /invites/{invite_id}][%d] getInviteByIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetInviteByIDBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetInviteByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInviteByIDUnauthorized creates a GetInviteByIDUnauthorized with default headers values
func NewGetInviteByIDUnauthorized() *GetInviteByIDUnauthorized {
	return &GetInviteByIDUnauthorized{}
}

/*GetInviteByIDUnauthorized handles this case with default header values.

AuthenticationError
*/
type GetInviteByIDUnauthorized struct {
	Payload *models.AuthenticationError
}

func (o *GetInviteByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /invites/{invite_id}][%d] getInviteByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInviteByIDUnauthorized) GetPayload() *models.AuthenticationError {
	return o.Payload
}

func (o *GetInviteByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInviteByIDForbidden creates a GetInviteByIDForbidden with default headers values
func NewGetInviteByIDForbidden() *GetInviteByIDForbidden {
	return &GetInviteByIDForbidden{}
}

/*GetInviteByIDForbidden handles this case with default header values.

AuthorizationError
*/
type GetInviteByIDForbidden struct {
	Payload *models.AuthorizationError
}

func (o *GetInviteByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /invites/{invite_id}][%d] getInviteByIdForbidden  %+v", 403, o.Payload)
}

func (o *GetInviteByIDForbidden) GetPayload() *models.AuthorizationError {
	return o.Payload
}

func (o *GetInviteByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInviteByIDNotFound creates a GetInviteByIDNotFound with default headers values
func NewGetInviteByIDNotFound() *GetInviteByIDNotFound {
	return &GetInviteByIDNotFound{}
}

/*GetInviteByIDNotFound handles this case with default header values.

NotFoundError
*/
type GetInviteByIDNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetInviteByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /invites/{invite_id}][%d] getInviteByIdNotFound  %+v", 404, o.Payload)
}

func (o *GetInviteByIDNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetInviteByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInviteByIDInternalServerError creates a GetInviteByIDInternalServerError with default headers values
func NewGetInviteByIDInternalServerError() *GetInviteByIDInternalServerError {
	return &GetInviteByIDInternalServerError{}
}

/*GetInviteByIDInternalServerError handles this case with default header values.

InternalServerError
*/
type GetInviteByIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetInviteByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /invites/{invite_id}][%d] getInviteByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInviteByIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetInviteByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
