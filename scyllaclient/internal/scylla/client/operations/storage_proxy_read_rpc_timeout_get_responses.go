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

// StorageProxyReadRPCTimeoutGetReader is a Reader for the StorageProxyReadRPCTimeoutGet structure.
type StorageProxyReadRPCTimeoutGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyReadRPCTimeoutGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyReadRPCTimeoutGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyReadRPCTimeoutGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyReadRPCTimeoutGetOK creates a StorageProxyReadRPCTimeoutGetOK with default headers values
func NewStorageProxyReadRPCTimeoutGetOK() *StorageProxyReadRPCTimeoutGetOK {
	return &StorageProxyReadRPCTimeoutGetOK{}
}

/*StorageProxyReadRPCTimeoutGetOK handles this case with default header values.

StorageProxyReadRPCTimeoutGetOK storage proxy read Rpc timeout get o k
*/
type StorageProxyReadRPCTimeoutGetOK struct {
	Payload interface{}
}

func (o *StorageProxyReadRPCTimeoutGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *StorageProxyReadRPCTimeoutGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageProxyReadRPCTimeoutGetDefault creates a StorageProxyReadRPCTimeoutGetDefault with default headers values
func NewStorageProxyReadRPCTimeoutGetDefault(code int) *StorageProxyReadRPCTimeoutGetDefault {
	return &StorageProxyReadRPCTimeoutGetDefault{
		_statusCode: code,
	}
}

/*StorageProxyReadRPCTimeoutGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyReadRPCTimeoutGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy read Rpc timeout get default response
func (o *StorageProxyReadRPCTimeoutGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyReadRPCTimeoutGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyReadRPCTimeoutGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyReadRPCTimeoutGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
