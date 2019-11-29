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

// ColumnFamilyMetricsWriteLatencyGetReader is a Reader for the ColumnFamilyMetricsWriteLatencyGet structure.
type ColumnFamilyMetricsWriteLatencyGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsWriteLatencyGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsWriteLatencyGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsWriteLatencyGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsWriteLatencyGetOK creates a ColumnFamilyMetricsWriteLatencyGetOK with default headers values
func NewColumnFamilyMetricsWriteLatencyGetOK() *ColumnFamilyMetricsWriteLatencyGetOK {
	return &ColumnFamilyMetricsWriteLatencyGetOK{}
}

/*ColumnFamilyMetricsWriteLatencyGetOK handles this case with default header values.

ColumnFamilyMetricsWriteLatencyGetOK column family metrics write latency get o k
*/
type ColumnFamilyMetricsWriteLatencyGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsWriteLatencyGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *ColumnFamilyMetricsWriteLatencyGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsWriteLatencyGetDefault creates a ColumnFamilyMetricsWriteLatencyGetDefault with default headers values
func NewColumnFamilyMetricsWriteLatencyGetDefault(code int) *ColumnFamilyMetricsWriteLatencyGetDefault {
	return &ColumnFamilyMetricsWriteLatencyGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsWriteLatencyGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsWriteLatencyGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics write latency get default response
func (o *ColumnFamilyMetricsWriteLatencyGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsWriteLatencyGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsWriteLatencyGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsWriteLatencyGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
