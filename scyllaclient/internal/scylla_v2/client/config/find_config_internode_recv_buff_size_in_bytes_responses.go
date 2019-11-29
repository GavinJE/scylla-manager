// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/scylla_v2/models"
)

// FindConfigInternodeRecvBuffSizeInBytesReader is a Reader for the FindConfigInternodeRecvBuffSizeInBytes structure.
type FindConfigInternodeRecvBuffSizeInBytesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigInternodeRecvBuffSizeInBytesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigInternodeRecvBuffSizeInBytesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigInternodeRecvBuffSizeInBytesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigInternodeRecvBuffSizeInBytesOK creates a FindConfigInternodeRecvBuffSizeInBytesOK with default headers values
func NewFindConfigInternodeRecvBuffSizeInBytesOK() *FindConfigInternodeRecvBuffSizeInBytesOK {
	return &FindConfigInternodeRecvBuffSizeInBytesOK{}
}

/*FindConfigInternodeRecvBuffSizeInBytesOK handles this case with default header values.

Config value
*/
type FindConfigInternodeRecvBuffSizeInBytesOK struct {
	Payload int64
}

func (o *FindConfigInternodeRecvBuffSizeInBytesOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigInternodeRecvBuffSizeInBytesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigInternodeRecvBuffSizeInBytesDefault creates a FindConfigInternodeRecvBuffSizeInBytesDefault with default headers values
func NewFindConfigInternodeRecvBuffSizeInBytesDefault(code int) *FindConfigInternodeRecvBuffSizeInBytesDefault {
	return &FindConfigInternodeRecvBuffSizeInBytesDefault{
		_statusCode: code,
	}
}

/*FindConfigInternodeRecvBuffSizeInBytesDefault handles this case with default header values.

unexpected error
*/
type FindConfigInternodeRecvBuffSizeInBytesDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config internode recv buff size in bytes default response
func (o *FindConfigInternodeRecvBuffSizeInBytesDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigInternodeRecvBuffSizeInBytesDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigInternodeRecvBuffSizeInBytesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigInternodeRecvBuffSizeInBytesDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
