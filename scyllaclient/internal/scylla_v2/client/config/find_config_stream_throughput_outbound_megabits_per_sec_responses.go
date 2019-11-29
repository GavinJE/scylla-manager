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

// FindConfigStreamThroughputOutboundMegabitsPerSecReader is a Reader for the FindConfigStreamThroughputOutboundMegabitsPerSec structure.
type FindConfigStreamThroughputOutboundMegabitsPerSecReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigStreamThroughputOutboundMegabitsPerSecReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigStreamThroughputOutboundMegabitsPerSecOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigStreamThroughputOutboundMegabitsPerSecDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigStreamThroughputOutboundMegabitsPerSecOK creates a FindConfigStreamThroughputOutboundMegabitsPerSecOK with default headers values
func NewFindConfigStreamThroughputOutboundMegabitsPerSecOK() *FindConfigStreamThroughputOutboundMegabitsPerSecOK {
	return &FindConfigStreamThroughputOutboundMegabitsPerSecOK{}
}

/*FindConfigStreamThroughputOutboundMegabitsPerSecOK handles this case with default header values.

Config value
*/
type FindConfigStreamThroughputOutboundMegabitsPerSecOK struct {
	Payload int64
}

func (o *FindConfigStreamThroughputOutboundMegabitsPerSecOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigStreamThroughputOutboundMegabitsPerSecOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigStreamThroughputOutboundMegabitsPerSecDefault creates a FindConfigStreamThroughputOutboundMegabitsPerSecDefault with default headers values
func NewFindConfigStreamThroughputOutboundMegabitsPerSecDefault(code int) *FindConfigStreamThroughputOutboundMegabitsPerSecDefault {
	return &FindConfigStreamThroughputOutboundMegabitsPerSecDefault{
		_statusCode: code,
	}
}

/*FindConfigStreamThroughputOutboundMegabitsPerSecDefault handles this case with default header values.

unexpected error
*/
type FindConfigStreamThroughputOutboundMegabitsPerSecDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config stream throughput outbound megabits per sec default response
func (o *FindConfigStreamThroughputOutboundMegabitsPerSecDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigStreamThroughputOutboundMegabitsPerSecDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigStreamThroughputOutboundMegabitsPerSecDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigStreamThroughputOutboundMegabitsPerSecDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
