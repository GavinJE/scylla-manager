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

// FindConfigPermissionsCacheMaxEntriesReader is a Reader for the FindConfigPermissionsCacheMaxEntries structure.
type FindConfigPermissionsCacheMaxEntriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigPermissionsCacheMaxEntriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigPermissionsCacheMaxEntriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigPermissionsCacheMaxEntriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigPermissionsCacheMaxEntriesOK creates a FindConfigPermissionsCacheMaxEntriesOK with default headers values
func NewFindConfigPermissionsCacheMaxEntriesOK() *FindConfigPermissionsCacheMaxEntriesOK {
	return &FindConfigPermissionsCacheMaxEntriesOK{}
}

/*FindConfigPermissionsCacheMaxEntriesOK handles this case with default header values.

Config value
*/
type FindConfigPermissionsCacheMaxEntriesOK struct {
	Payload int64
}

func (o *FindConfigPermissionsCacheMaxEntriesOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigPermissionsCacheMaxEntriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigPermissionsCacheMaxEntriesDefault creates a FindConfigPermissionsCacheMaxEntriesDefault with default headers values
func NewFindConfigPermissionsCacheMaxEntriesDefault(code int) *FindConfigPermissionsCacheMaxEntriesDefault {
	return &FindConfigPermissionsCacheMaxEntriesDefault{
		_statusCode: code,
	}
}

/*FindConfigPermissionsCacheMaxEntriesDefault handles this case with default header values.

unexpected error
*/
type FindConfigPermissionsCacheMaxEntriesDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config permissions cache max entries default response
func (o *FindConfigPermissionsCacheMaxEntriesDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigPermissionsCacheMaxEntriesDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigPermissionsCacheMaxEntriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigPermissionsCacheMaxEntriesDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
