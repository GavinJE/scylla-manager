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

// FindConfigRequestSchedulerOptionsReader is a Reader for the FindConfigRequestSchedulerOptions structure.
type FindConfigRequestSchedulerOptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigRequestSchedulerOptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigRequestSchedulerOptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigRequestSchedulerOptionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigRequestSchedulerOptionsOK creates a FindConfigRequestSchedulerOptionsOK with default headers values
func NewFindConfigRequestSchedulerOptionsOK() *FindConfigRequestSchedulerOptionsOK {
	return &FindConfigRequestSchedulerOptionsOK{}
}

/*FindConfigRequestSchedulerOptionsOK handles this case with default header values.

Config value
*/
type FindConfigRequestSchedulerOptionsOK struct {
	Payload []string
}

func (o *FindConfigRequestSchedulerOptionsOK) GetPayload() []string {
	return o.Payload
}

func (o *FindConfigRequestSchedulerOptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigRequestSchedulerOptionsDefault creates a FindConfigRequestSchedulerOptionsDefault with default headers values
func NewFindConfigRequestSchedulerOptionsDefault(code int) *FindConfigRequestSchedulerOptionsDefault {
	return &FindConfigRequestSchedulerOptionsDefault{
		_statusCode: code,
	}
}

/*FindConfigRequestSchedulerOptionsDefault handles this case with default header values.

unexpected error
*/
type FindConfigRequestSchedulerOptionsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config request scheduler options default response
func (o *FindConfigRequestSchedulerOptionsDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigRequestSchedulerOptionsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigRequestSchedulerOptionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigRequestSchedulerOptionsDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
