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

// ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetReader is a Reader for the ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGet structure.
type ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK creates a ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK with default headers values
func NewColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK() *ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK {
	return &ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK{}
}

/*ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK handles this case with default header values.

ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK column family metrics write latency estimated histogram by name get o k
*/
type ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK struct {
}

func (o *ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetDefault creates a ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetDefault with default headers values
func NewColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetDefault(code int) *ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetDefault {
	return &ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics write latency estimated histogram by name get default response
func (o *ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
