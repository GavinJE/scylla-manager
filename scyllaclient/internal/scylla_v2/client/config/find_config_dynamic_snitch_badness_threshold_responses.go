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

// FindConfigDynamicSnitchBadnessThresholdReader is a Reader for the FindConfigDynamicSnitchBadnessThreshold structure.
type FindConfigDynamicSnitchBadnessThresholdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigDynamicSnitchBadnessThresholdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigDynamicSnitchBadnessThresholdOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigDynamicSnitchBadnessThresholdDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigDynamicSnitchBadnessThresholdOK creates a FindConfigDynamicSnitchBadnessThresholdOK with default headers values
func NewFindConfigDynamicSnitchBadnessThresholdOK() *FindConfigDynamicSnitchBadnessThresholdOK {
	return &FindConfigDynamicSnitchBadnessThresholdOK{}
}

/*FindConfigDynamicSnitchBadnessThresholdOK handles this case with default header values.

Config value
*/
type FindConfigDynamicSnitchBadnessThresholdOK struct {
	Payload float64
}

func (o *FindConfigDynamicSnitchBadnessThresholdOK) GetPayload() float64 {
	return o.Payload
}

func (o *FindConfigDynamicSnitchBadnessThresholdOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigDynamicSnitchBadnessThresholdDefault creates a FindConfigDynamicSnitchBadnessThresholdDefault with default headers values
func NewFindConfigDynamicSnitchBadnessThresholdDefault(code int) *FindConfigDynamicSnitchBadnessThresholdDefault {
	return &FindConfigDynamicSnitchBadnessThresholdDefault{
		_statusCode: code,
	}
}

/*FindConfigDynamicSnitchBadnessThresholdDefault handles this case with default header values.

unexpected error
*/
type FindConfigDynamicSnitchBadnessThresholdDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config dynamic snitch badness threshold default response
func (o *FindConfigDynamicSnitchBadnessThresholdDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigDynamicSnitchBadnessThresholdDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigDynamicSnitchBadnessThresholdDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigDynamicSnitchBadnessThresholdDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
