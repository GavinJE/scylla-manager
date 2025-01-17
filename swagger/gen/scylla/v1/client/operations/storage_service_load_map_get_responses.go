// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/swagger/gen/scylla/v1/models"
)

// StorageServiceLoadMapGetReader is a Reader for the StorageServiceLoadMapGet structure.
type StorageServiceLoadMapGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceLoadMapGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceLoadMapGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceLoadMapGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceLoadMapGetOK creates a StorageServiceLoadMapGetOK with default headers values
func NewStorageServiceLoadMapGetOK() *StorageServiceLoadMapGetOK {
	return &StorageServiceLoadMapGetOK{}
}

/*StorageServiceLoadMapGetOK handles this case with default header values.

StorageServiceLoadMapGetOK storage service load map get o k
*/
type StorageServiceLoadMapGetOK struct {
	Payload []*models.MapStringDouble
}

func (o *StorageServiceLoadMapGetOK) GetPayload() []*models.MapStringDouble {
	return o.Payload
}

func (o *StorageServiceLoadMapGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceLoadMapGetDefault creates a StorageServiceLoadMapGetDefault with default headers values
func NewStorageServiceLoadMapGetDefault(code int) *StorageServiceLoadMapGetDefault {
	return &StorageServiceLoadMapGetDefault{
		_statusCode: code,
	}
}

/*StorageServiceLoadMapGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceLoadMapGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service load map get default response
func (o *StorageServiceLoadMapGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceLoadMapGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceLoadMapGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceLoadMapGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
