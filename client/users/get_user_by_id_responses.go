// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fugue/fugue-client/models"
)

// GetUserByIDReader is a Reader for the GetUserByID structure.
type GetUserByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUserByIDOK creates a GetUserByIDOK with default headers values
func NewGetUserByIDOK() *GetUserByIDOK {
	return &GetUserByIDOK{}
}

/*GetUserByIDOK handles this case with default header values.

User details.
*/
type GetUserByIDOK struct {
	Payload *models.User
}

func (o *GetUserByIDOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserByIdOK  %+v", 200, o.Payload)
}

func (o *GetUserByIDOK) GetPayload() *models.User {
	return o.Payload
}

func (o *GetUserByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserByIDBadRequest creates a GetUserByIDBadRequest with default headers values
func NewGetUserByIDBadRequest() *GetUserByIDBadRequest {
	return &GetUserByIDBadRequest{}
}

/*GetUserByIDBadRequest handles this case with default header values.

Bad request error.
*/
type GetUserByIDBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetUserByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserByIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserByIDBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetUserByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserByIDUnauthorized creates a GetUserByIDUnauthorized with default headers values
func NewGetUserByIDUnauthorized() *GetUserByIDUnauthorized {
	return &GetUserByIDUnauthorized{}
}

/*GetUserByIDUnauthorized handles this case with default header values.

Authentication error.
*/
type GetUserByIDUnauthorized struct {
	Payload *models.AuthenticationError
}

func (o *GetUserByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserByIDUnauthorized) GetPayload() *models.AuthenticationError {
	return o.Payload
}

func (o *GetUserByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserByIDForbidden creates a GetUserByIDForbidden with default headers values
func NewGetUserByIDForbidden() *GetUserByIDForbidden {
	return &GetUserByIDForbidden{}
}

/*GetUserByIDForbidden handles this case with default header values.

Authorization error.
*/
type GetUserByIDForbidden struct {
	Payload *models.AuthorizationError
}

func (o *GetUserByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserByIdForbidden  %+v", 403, o.Payload)
}

func (o *GetUserByIDForbidden) GetPayload() *models.AuthorizationError {
	return o.Payload
}

func (o *GetUserByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserByIDNotFound creates a GetUserByIDNotFound with default headers values
func NewGetUserByIDNotFound() *GetUserByIDNotFound {
	return &GetUserByIDNotFound{}
}

/*GetUserByIDNotFound handles this case with default header values.

Not found error.
*/
type GetUserByIDNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetUserByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserByIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUserByIDNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetUserByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserByIDInternalServerError creates a GetUserByIDInternalServerError with default headers values
func NewGetUserByIDInternalServerError() *GetUserByIDInternalServerError {
	return &GetUserByIDInternalServerError{}
}

/*GetUserByIDInternalServerError handles this case with default header values.

Internal server error.
*/
type GetUserByIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetUserByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserByIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUserByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
