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

// FindConfigPermissionsValidityInMsReader is a Reader for the FindConfigPermissionsValidityInMs structure.
type FindConfigPermissionsValidityInMsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigPermissionsValidityInMsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigPermissionsValidityInMsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigPermissionsValidityInMsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigPermissionsValidityInMsOK creates a FindConfigPermissionsValidityInMsOK with default headers values
func NewFindConfigPermissionsValidityInMsOK() *FindConfigPermissionsValidityInMsOK {
	return &FindConfigPermissionsValidityInMsOK{}
}

/*FindConfigPermissionsValidityInMsOK handles this case with default header values.

Config value
*/
type FindConfigPermissionsValidityInMsOK struct {
	Payload int64
}

func (o *FindConfigPermissionsValidityInMsOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigPermissionsValidityInMsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigPermissionsValidityInMsDefault creates a FindConfigPermissionsValidityInMsDefault with default headers values
func NewFindConfigPermissionsValidityInMsDefault(code int) *FindConfigPermissionsValidityInMsDefault {
	return &FindConfigPermissionsValidityInMsDefault{
		_statusCode: code,
	}
}

/*FindConfigPermissionsValidityInMsDefault handles this case with default header values.

unexpected error
*/
type FindConfigPermissionsValidityInMsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config permissions validity in ms default response
func (o *FindConfigPermissionsValidityInMsDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigPermissionsValidityInMsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigPermissionsValidityInMsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigPermissionsValidityInMsDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
