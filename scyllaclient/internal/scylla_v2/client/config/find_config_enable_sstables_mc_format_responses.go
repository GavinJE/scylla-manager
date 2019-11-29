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

// FindConfigEnableSstablesMcFormatReader is a Reader for the FindConfigEnableSstablesMcFormat structure.
type FindConfigEnableSstablesMcFormatReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigEnableSstablesMcFormatReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigEnableSstablesMcFormatOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigEnableSstablesMcFormatDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigEnableSstablesMcFormatOK creates a FindConfigEnableSstablesMcFormatOK with default headers values
func NewFindConfigEnableSstablesMcFormatOK() *FindConfigEnableSstablesMcFormatOK {
	return &FindConfigEnableSstablesMcFormatOK{}
}

/*FindConfigEnableSstablesMcFormatOK handles this case with default header values.

Config value
*/
type FindConfigEnableSstablesMcFormatOK struct {
	Payload bool
}

func (o *FindConfigEnableSstablesMcFormatOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigEnableSstablesMcFormatOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigEnableSstablesMcFormatDefault creates a FindConfigEnableSstablesMcFormatDefault with default headers values
func NewFindConfigEnableSstablesMcFormatDefault(code int) *FindConfigEnableSstablesMcFormatDefault {
	return &FindConfigEnableSstablesMcFormatDefault{
		_statusCode: code,
	}
}

/*FindConfigEnableSstablesMcFormatDefault handles this case with default header values.

unexpected error
*/
type FindConfigEnableSstablesMcFormatDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config enable sstables mc format default response
func (o *FindConfigEnableSstablesMcFormatDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigEnableSstablesMcFormatDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigEnableSstablesMcFormatDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigEnableSstablesMcFormatDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
