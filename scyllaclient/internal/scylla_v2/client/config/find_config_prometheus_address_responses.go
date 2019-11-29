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

// FindConfigPrometheusAddressReader is a Reader for the FindConfigPrometheusAddress structure.
type FindConfigPrometheusAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigPrometheusAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigPrometheusAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigPrometheusAddressDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigPrometheusAddressOK creates a FindConfigPrometheusAddressOK with default headers values
func NewFindConfigPrometheusAddressOK() *FindConfigPrometheusAddressOK {
	return &FindConfigPrometheusAddressOK{}
}

/*FindConfigPrometheusAddressOK handles this case with default header values.

Config value
*/
type FindConfigPrometheusAddressOK struct {
	Payload string
}

func (o *FindConfigPrometheusAddressOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigPrometheusAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigPrometheusAddressDefault creates a FindConfigPrometheusAddressDefault with default headers values
func NewFindConfigPrometheusAddressDefault(code int) *FindConfigPrometheusAddressDefault {
	return &FindConfigPrometheusAddressDefault{
		_statusCode: code,
	}
}

/*FindConfigPrometheusAddressDefault handles this case with default header values.

unexpected error
*/
type FindConfigPrometheusAddressDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config prometheus address default response
func (o *FindConfigPrometheusAddressDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigPrometheusAddressDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigPrometheusAddressDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigPrometheusAddressDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
