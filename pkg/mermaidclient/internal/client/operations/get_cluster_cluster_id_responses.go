// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/pkg/mermaidclient/internal/models"
)

// GetClusterClusterIDReader is a Reader for the GetClusterClusterID structure.
type GetClusterClusterIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterClusterIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterClusterIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetClusterClusterIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetClusterClusterIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetClusterClusterIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetClusterClusterIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterClusterIDOK creates a GetClusterClusterIDOK with default headers values
func NewGetClusterClusterIDOK() *GetClusterClusterIDOK {
	return &GetClusterClusterIDOK{}
}

/*GetClusterClusterIDOK handles this case with default header values.

Cluster info
*/
type GetClusterClusterIDOK struct {
	Payload *models.Cluster
}

func (o *GetClusterClusterIDOK) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}][%d] getClusterClusterIdOK  %+v", 200, o.Payload)
}

func (o *GetClusterClusterIDOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *GetClusterClusterIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterClusterIDBadRequest creates a GetClusterClusterIDBadRequest with default headers values
func NewGetClusterClusterIDBadRequest() *GetClusterClusterIDBadRequest {
	return &GetClusterClusterIDBadRequest{}
}

/*GetClusterClusterIDBadRequest handles this case with default header values.

Bad Request
*/
type GetClusterClusterIDBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetClusterClusterIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}][%d] getClusterClusterIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetClusterClusterIDBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterClusterIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterClusterIDNotFound creates a GetClusterClusterIDNotFound with default headers values
func NewGetClusterClusterIDNotFound() *GetClusterClusterIDNotFound {
	return &GetClusterClusterIDNotFound{}
}

/*GetClusterClusterIDNotFound handles this case with default header values.

Not found
*/
type GetClusterClusterIDNotFound struct {
	Payload *models.ErrorResponse
}

func (o *GetClusterClusterIDNotFound) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}][%d] getClusterClusterIdNotFound  %+v", 404, o.Payload)
}

func (o *GetClusterClusterIDNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterClusterIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterClusterIDInternalServerError creates a GetClusterClusterIDInternalServerError with default headers values
func NewGetClusterClusterIDInternalServerError() *GetClusterClusterIDInternalServerError {
	return &GetClusterClusterIDInternalServerError{}
}

/*GetClusterClusterIDInternalServerError handles this case with default header values.

Server error
*/
type GetClusterClusterIDInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetClusterClusterIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}][%d] getClusterClusterIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetClusterClusterIDInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterClusterIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterClusterIDDefault creates a GetClusterClusterIDDefault with default headers values
func NewGetClusterClusterIDDefault(code int) *GetClusterClusterIDDefault {
	return &GetClusterClusterIDDefault{
		_statusCode: code,
	}
}

/*GetClusterClusterIDDefault handles this case with default header values.

Unexpected error
*/
type GetClusterClusterIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get cluster cluster ID default response
func (o *GetClusterClusterIDDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterClusterIDDefault) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}][%d] GetClusterClusterID default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterClusterIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterClusterIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}