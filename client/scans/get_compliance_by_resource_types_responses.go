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

// GetComplianceByResourceTypesReader is a Reader for the GetComplianceByResourceTypes structure.
type GetComplianceByResourceTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComplianceByResourceTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComplianceByResourceTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetComplianceByResourceTypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetComplianceByResourceTypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetComplianceByResourceTypesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetComplianceByResourceTypesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetComplianceByResourceTypesOK creates a GetComplianceByResourceTypesOK with default headers values
func NewGetComplianceByResourceTypesOK() *GetComplianceByResourceTypesOK {
	return &GetComplianceByResourceTypesOK{}
}

/*GetComplianceByResourceTypesOK handles this case with default header values.

List of compliance results from a scan grouped by resource type.
*/
type GetComplianceByResourceTypesOK struct {
	Payload *models.ComplianceByResourceTypeOutput
}

func (o *GetComplianceByResourceTypesOK) Error() string {
	return fmt.Sprintf("[GET /scans/{scan_id}/compliance_by_resource_types][%d] getComplianceByResourceTypesOK  %+v", 200, o.Payload)
}

func (o *GetComplianceByResourceTypesOK) GetPayload() *models.ComplianceByResourceTypeOutput {
	return o.Payload
}

func (o *GetComplianceByResourceTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComplianceByResourceTypeOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComplianceByResourceTypesUnauthorized creates a GetComplianceByResourceTypesUnauthorized with default headers values
func NewGetComplianceByResourceTypesUnauthorized() *GetComplianceByResourceTypesUnauthorized {
	return &GetComplianceByResourceTypesUnauthorized{}
}

/*GetComplianceByResourceTypesUnauthorized handles this case with default header values.

Authentication error.
*/
type GetComplianceByResourceTypesUnauthorized struct {
	Payload *models.AuthenticationError
}

func (o *GetComplianceByResourceTypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /scans/{scan_id}/compliance_by_resource_types][%d] getComplianceByResourceTypesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetComplianceByResourceTypesUnauthorized) GetPayload() *models.AuthenticationError {
	return o.Payload
}

func (o *GetComplianceByResourceTypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComplianceByResourceTypesForbidden creates a GetComplianceByResourceTypesForbidden with default headers values
func NewGetComplianceByResourceTypesForbidden() *GetComplianceByResourceTypesForbidden {
	return &GetComplianceByResourceTypesForbidden{}
}

/*GetComplianceByResourceTypesForbidden handles this case with default header values.

Authorization error.
*/
type GetComplianceByResourceTypesForbidden struct {
	Payload *models.AuthorizationError
}

func (o *GetComplianceByResourceTypesForbidden) Error() string {
	return fmt.Sprintf("[GET /scans/{scan_id}/compliance_by_resource_types][%d] getComplianceByResourceTypesForbidden  %+v", 403, o.Payload)
}

func (o *GetComplianceByResourceTypesForbidden) GetPayload() *models.AuthorizationError {
	return o.Payload
}

func (o *GetComplianceByResourceTypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComplianceByResourceTypesNotFound creates a GetComplianceByResourceTypesNotFound with default headers values
func NewGetComplianceByResourceTypesNotFound() *GetComplianceByResourceTypesNotFound {
	return &GetComplianceByResourceTypesNotFound{}
}

/*GetComplianceByResourceTypesNotFound handles this case with default header values.

Not found error.
*/
type GetComplianceByResourceTypesNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetComplianceByResourceTypesNotFound) Error() string {
	return fmt.Sprintf("[GET /scans/{scan_id}/compliance_by_resource_types][%d] getComplianceByResourceTypesNotFound  %+v", 404, o.Payload)
}

func (o *GetComplianceByResourceTypesNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetComplianceByResourceTypesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComplianceByResourceTypesInternalServerError creates a GetComplianceByResourceTypesInternalServerError with default headers values
func NewGetComplianceByResourceTypesInternalServerError() *GetComplianceByResourceTypesInternalServerError {
	return &GetComplianceByResourceTypesInternalServerError{}
}

/*GetComplianceByResourceTypesInternalServerError handles this case with default header values.

Internal server error.
*/
type GetComplianceByResourceTypesInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetComplianceByResourceTypesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /scans/{scan_id}/compliance_by_resource_types][%d] getComplianceByResourceTypesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetComplianceByResourceTypesInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetComplianceByResourceTypesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
