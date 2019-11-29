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

// StorageProxyReadRepairRepairedBlockingGetReader is a Reader for the StorageProxyReadRepairRepairedBlockingGet structure.
type StorageProxyReadRepairRepairedBlockingGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyReadRepairRepairedBlockingGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyReadRepairRepairedBlockingGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyReadRepairRepairedBlockingGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyReadRepairRepairedBlockingGetOK creates a StorageProxyReadRepairRepairedBlockingGetOK with default headers values
func NewStorageProxyReadRepairRepairedBlockingGetOK() *StorageProxyReadRepairRepairedBlockingGetOK {
	return &StorageProxyReadRepairRepairedBlockingGetOK{}
}

/*StorageProxyReadRepairRepairedBlockingGetOK handles this case with default header values.

StorageProxyReadRepairRepairedBlockingGetOK storage proxy read repair repaired blocking get o k
*/
type StorageProxyReadRepairRepairedBlockingGetOK struct {
	Payload interface{}
}

func (o *StorageProxyReadRepairRepairedBlockingGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *StorageProxyReadRepairRepairedBlockingGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageProxyReadRepairRepairedBlockingGetDefault creates a StorageProxyReadRepairRepairedBlockingGetDefault with default headers values
func NewStorageProxyReadRepairRepairedBlockingGetDefault(code int) *StorageProxyReadRepairRepairedBlockingGetDefault {
	return &StorageProxyReadRepairRepairedBlockingGetDefault{
		_statusCode: code,
	}
}

/*StorageProxyReadRepairRepairedBlockingGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyReadRepairRepairedBlockingGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy read repair repaired blocking get default response
func (o *StorageProxyReadRepairRepairedBlockingGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyReadRepairRepairedBlockingGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyReadRepairRepairedBlockingGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyReadRepairRepairedBlockingGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
