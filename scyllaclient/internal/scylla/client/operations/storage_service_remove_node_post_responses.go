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

// StorageServiceRemoveNodePostReader is a Reader for the StorageServiceRemoveNodePost structure.
type StorageServiceRemoveNodePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceRemoveNodePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceRemoveNodePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceRemoveNodePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceRemoveNodePostOK creates a StorageServiceRemoveNodePostOK with default headers values
func NewStorageServiceRemoveNodePostOK() *StorageServiceRemoveNodePostOK {
	return &StorageServiceRemoveNodePostOK{}
}

/*StorageServiceRemoveNodePostOK handles this case with default header values.

StorageServiceRemoveNodePostOK storage service remove node post o k
*/
type StorageServiceRemoveNodePostOK struct {
}

func (o *StorageServiceRemoveNodePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceRemoveNodePostDefault creates a StorageServiceRemoveNodePostDefault with default headers values
func NewStorageServiceRemoveNodePostDefault(code int) *StorageServiceRemoveNodePostDefault {
	return &StorageServiceRemoveNodePostDefault{
		_statusCode: code,
	}
}

/*StorageServiceRemoveNodePostDefault handles this case with default header values.

internal server error
*/
type StorageServiceRemoveNodePostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service remove node post default response
func (o *StorageServiceRemoveNodePostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceRemoveNodePostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceRemoveNodePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceRemoveNodePostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
