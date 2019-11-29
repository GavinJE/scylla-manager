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

// FindConfigAbortOnLsaBadAllocReader is a Reader for the FindConfigAbortOnLsaBadAlloc structure.
type FindConfigAbortOnLsaBadAllocReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigAbortOnLsaBadAllocReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigAbortOnLsaBadAllocOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigAbortOnLsaBadAllocDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigAbortOnLsaBadAllocOK creates a FindConfigAbortOnLsaBadAllocOK with default headers values
func NewFindConfigAbortOnLsaBadAllocOK() *FindConfigAbortOnLsaBadAllocOK {
	return &FindConfigAbortOnLsaBadAllocOK{}
}

/*FindConfigAbortOnLsaBadAllocOK handles this case with default header values.

Config value
*/
type FindConfigAbortOnLsaBadAllocOK struct {
	Payload bool
}

func (o *FindConfigAbortOnLsaBadAllocOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigAbortOnLsaBadAllocOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigAbortOnLsaBadAllocDefault creates a FindConfigAbortOnLsaBadAllocDefault with default headers values
func NewFindConfigAbortOnLsaBadAllocDefault(code int) *FindConfigAbortOnLsaBadAllocDefault {
	return &FindConfigAbortOnLsaBadAllocDefault{
		_statusCode: code,
	}
}

/*FindConfigAbortOnLsaBadAllocDefault handles this case with default header values.

unexpected error
*/
type FindConfigAbortOnLsaBadAllocDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config abort on lsa bad alloc default response
func (o *FindConfigAbortOnLsaBadAllocDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigAbortOnLsaBadAllocDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigAbortOnLsaBadAllocDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigAbortOnLsaBadAllocDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
