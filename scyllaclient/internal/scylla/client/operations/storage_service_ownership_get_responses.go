// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/scylla/models"
)

// StorageServiceOwnershipGetReader is a Reader for the StorageServiceOwnershipGet structure.
type StorageServiceOwnershipGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceOwnershipGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceOwnershipGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceOwnershipGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceOwnershipGetOK creates a StorageServiceOwnershipGetOK with default headers values
func NewStorageServiceOwnershipGetOK() *StorageServiceOwnershipGetOK {
	return &StorageServiceOwnershipGetOK{}
}

/*StorageServiceOwnershipGetOK handles this case with default header values.

StorageServiceOwnershipGetOK storage service ownership get o k
*/
type StorageServiceOwnershipGetOK struct {
	Payload []*models.Mapper
}

func (o *StorageServiceOwnershipGetOK) GetPayload() []*models.Mapper {
	return o.Payload
}

func (o *StorageServiceOwnershipGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceOwnershipGetDefault creates a StorageServiceOwnershipGetDefault with default headers values
func NewStorageServiceOwnershipGetDefault(code int) *StorageServiceOwnershipGetDefault {
	return &StorageServiceOwnershipGetDefault{
		_statusCode: code,
	}
}

/*StorageServiceOwnershipGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceOwnershipGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service ownership get default response
func (o *StorageServiceOwnershipGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceOwnershipGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceOwnershipGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceOwnershipGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
