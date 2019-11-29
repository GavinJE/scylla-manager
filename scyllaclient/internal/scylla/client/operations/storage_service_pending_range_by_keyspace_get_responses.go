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

// StorageServicePendingRangeByKeyspaceGetReader is a Reader for the StorageServicePendingRangeByKeyspaceGet structure.
type StorageServicePendingRangeByKeyspaceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServicePendingRangeByKeyspaceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServicePendingRangeByKeyspaceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServicePendingRangeByKeyspaceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServicePendingRangeByKeyspaceGetOK creates a StorageServicePendingRangeByKeyspaceGetOK with default headers values
func NewStorageServicePendingRangeByKeyspaceGetOK() *StorageServicePendingRangeByKeyspaceGetOK {
	return &StorageServicePendingRangeByKeyspaceGetOK{}
}

/*StorageServicePendingRangeByKeyspaceGetOK handles this case with default header values.

StorageServicePendingRangeByKeyspaceGetOK storage service pending range by keyspace get o k
*/
type StorageServicePendingRangeByKeyspaceGetOK struct {
	Payload []*models.MaplistMapper
}

func (o *StorageServicePendingRangeByKeyspaceGetOK) GetPayload() []*models.MaplistMapper {
	return o.Payload
}

func (o *StorageServicePendingRangeByKeyspaceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServicePendingRangeByKeyspaceGetDefault creates a StorageServicePendingRangeByKeyspaceGetDefault with default headers values
func NewStorageServicePendingRangeByKeyspaceGetDefault(code int) *StorageServicePendingRangeByKeyspaceGetDefault {
	return &StorageServicePendingRangeByKeyspaceGetDefault{
		_statusCode: code,
	}
}

/*StorageServicePendingRangeByKeyspaceGetDefault handles this case with default header values.

internal server error
*/
type StorageServicePendingRangeByKeyspaceGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service pending range by keyspace get default response
func (o *StorageServicePendingRangeByKeyspaceGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServicePendingRangeByKeyspaceGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServicePendingRangeByKeyspaceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServicePendingRangeByKeyspaceGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
