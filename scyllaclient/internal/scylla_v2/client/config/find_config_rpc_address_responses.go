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

// FindConfigRPCAddressReader is a Reader for the FindConfigRPCAddress structure.
type FindConfigRPCAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigRPCAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigRPCAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigRPCAddressDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigRPCAddressOK creates a FindConfigRPCAddressOK with default headers values
func NewFindConfigRPCAddressOK() *FindConfigRPCAddressOK {
	return &FindConfigRPCAddressOK{}
}

/*FindConfigRPCAddressOK handles this case with default header values.

Config value
*/
type FindConfigRPCAddressOK struct {
	Payload string
}

func (o *FindConfigRPCAddressOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigRPCAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigRPCAddressDefault creates a FindConfigRPCAddressDefault with default headers values
func NewFindConfigRPCAddressDefault(code int) *FindConfigRPCAddressDefault {
	return &FindConfigRPCAddressDefault{
		_statusCode: code,
	}
}

/*FindConfigRPCAddressDefault handles this case with default header values.

unexpected error
*/
type FindConfigRPCAddressDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config rpc address default response
func (o *FindConfigRPCAddressDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigRPCAddressDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigRPCAddressDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigRPCAddressDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
