// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/scylla/models"
)

// ColumnFamilyMetricsSpeculativeRetriesGetReader is a Reader for the ColumnFamilyMetricsSpeculativeRetriesGet structure.
type ColumnFamilyMetricsSpeculativeRetriesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsSpeculativeRetriesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsSpeculativeRetriesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsSpeculativeRetriesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsSpeculativeRetriesGetOK creates a ColumnFamilyMetricsSpeculativeRetriesGetOK with default headers values
func NewColumnFamilyMetricsSpeculativeRetriesGetOK() *ColumnFamilyMetricsSpeculativeRetriesGetOK {
	return &ColumnFamilyMetricsSpeculativeRetriesGetOK{}
}

/*ColumnFamilyMetricsSpeculativeRetriesGetOK handles this case with default header values.

ColumnFamilyMetricsSpeculativeRetriesGetOK column family metrics speculative retries get o k
*/
type ColumnFamilyMetricsSpeculativeRetriesGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsSpeculativeRetriesGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *ColumnFamilyMetricsSpeculativeRetriesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsSpeculativeRetriesGetDefault creates a ColumnFamilyMetricsSpeculativeRetriesGetDefault with default headers values
func NewColumnFamilyMetricsSpeculativeRetriesGetDefault(code int) *ColumnFamilyMetricsSpeculativeRetriesGetDefault {
	return &ColumnFamilyMetricsSpeculativeRetriesGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsSpeculativeRetriesGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsSpeculativeRetriesGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics speculative retries get default response
func (o *ColumnFamilyMetricsSpeculativeRetriesGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsSpeculativeRetriesGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsSpeculativeRetriesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsSpeculativeRetriesGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
