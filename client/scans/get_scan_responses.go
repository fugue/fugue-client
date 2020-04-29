// Code generated by go-swagger; DO NOT EDIT.

package scans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fugue/fugue-client/models"
)

// GetScanReader is a Reader for the GetScan structure.
type GetScanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetScanUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScanForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetScanOK creates a GetScanOK with default headers values
func NewGetScanOK() *GetScanOK {
	return &GetScanOK{}
}

/*GetScanOK handles this case with default header values.

Scan details.
*/
type GetScanOK struct {
	Payload *models.ScanWithSummary
}

func (o *GetScanOK) Error() string {
	return fmt.Sprintf("[GET /scans/{scan_id}][%d] getScanOK  %+v", 200, o.Payload)
}

func (o *GetScanOK) GetPayload() *models.ScanWithSummary {
	return o.Payload
}

func (o *GetScanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScanWithSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScanUnauthorized creates a GetScanUnauthorized with default headers values
func NewGetScanUnauthorized() *GetScanUnauthorized {
	return &GetScanUnauthorized{}
}

/*GetScanUnauthorized handles this case with default header values.

Authentication error.
*/
type GetScanUnauthorized struct {
	Payload *models.AuthenticationError
}

func (o *GetScanUnauthorized) Error() string {
	return fmt.Sprintf("[GET /scans/{scan_id}][%d] getScanUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScanUnauthorized) GetPayload() *models.AuthenticationError {
	return o.Payload
}

func (o *GetScanUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScanForbidden creates a GetScanForbidden with default headers values
func NewGetScanForbidden() *GetScanForbidden {
	return &GetScanForbidden{}
}

/*GetScanForbidden handles this case with default header values.

Authorization error.
*/
type GetScanForbidden struct {
	Payload *models.AuthorizationError
}

func (o *GetScanForbidden) Error() string {
	return fmt.Sprintf("[GET /scans/{scan_id}][%d] getScanForbidden  %+v", 403, o.Payload)
}

func (o *GetScanForbidden) GetPayload() *models.AuthorizationError {
	return o.Payload
}

func (o *GetScanForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScanNotFound creates a GetScanNotFound with default headers values
func NewGetScanNotFound() *GetScanNotFound {
	return &GetScanNotFound{}
}

/*GetScanNotFound handles this case with default header values.

Not found error.
*/
type GetScanNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetScanNotFound) Error() string {
	return fmt.Sprintf("[GET /scans/{scan_id}][%d] getScanNotFound  %+v", 404, o.Payload)
}

func (o *GetScanNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetScanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScanInternalServerError creates a GetScanInternalServerError with default headers values
func NewGetScanInternalServerError() *GetScanInternalServerError {
	return &GetScanInternalServerError{}
}

/*GetScanInternalServerError handles this case with default header values.

Internal server error.
*/
type GetScanInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetScanInternalServerError) Error() string {
	return fmt.Sprintf("[GET /scans/{scan_id}][%d] getScanInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScanInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetScanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
