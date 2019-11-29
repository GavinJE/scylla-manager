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

// FindConfigDefragmentMemoryOnIdleReader is a Reader for the FindConfigDefragmentMemoryOnIdle structure.
type FindConfigDefragmentMemoryOnIdleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigDefragmentMemoryOnIdleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigDefragmentMemoryOnIdleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigDefragmentMemoryOnIdleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigDefragmentMemoryOnIdleOK creates a FindConfigDefragmentMemoryOnIdleOK with default headers values
func NewFindConfigDefragmentMemoryOnIdleOK() *FindConfigDefragmentMemoryOnIdleOK {
	return &FindConfigDefragmentMemoryOnIdleOK{}
}

/*FindConfigDefragmentMemoryOnIdleOK handles this case with default header values.

Config value
*/
type FindConfigDefragmentMemoryOnIdleOK struct {
	Payload bool
}

func (o *FindConfigDefragmentMemoryOnIdleOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigDefragmentMemoryOnIdleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigDefragmentMemoryOnIdleDefault creates a FindConfigDefragmentMemoryOnIdleDefault with default headers values
func NewFindConfigDefragmentMemoryOnIdleDefault(code int) *FindConfigDefragmentMemoryOnIdleDefault {
	return &FindConfigDefragmentMemoryOnIdleDefault{
		_statusCode: code,
	}
}

/*FindConfigDefragmentMemoryOnIdleDefault handles this case with default header values.

unexpected error
*/
type FindConfigDefragmentMemoryOnIdleDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config defragment memory on idle default response
func (o *FindConfigDefragmentMemoryOnIdleDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigDefragmentMemoryOnIdleDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigDefragmentMemoryOnIdleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigDefragmentMemoryOnIdleDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
