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

// StorageProxyMetricsWriteTimeoutsRatesGetReader is a Reader for the StorageProxyMetricsWriteTimeoutsRatesGet structure.
type StorageProxyMetricsWriteTimeoutsRatesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMetricsWriteTimeoutsRatesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyMetricsWriteTimeoutsRatesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyMetricsWriteTimeoutsRatesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyMetricsWriteTimeoutsRatesGetOK creates a StorageProxyMetricsWriteTimeoutsRatesGetOK with default headers values
func NewStorageProxyMetricsWriteTimeoutsRatesGetOK() *StorageProxyMetricsWriteTimeoutsRatesGetOK {
	return &StorageProxyMetricsWriteTimeoutsRatesGetOK{}
}

/*StorageProxyMetricsWriteTimeoutsRatesGetOK handles this case with default header values.

StorageProxyMetricsWriteTimeoutsRatesGetOK storage proxy metrics write timeouts rates get o k
*/
type StorageProxyMetricsWriteTimeoutsRatesGetOK struct {
	Payload *models.RateMovingAverage
}

func (o *StorageProxyMetricsWriteTimeoutsRatesGetOK) GetPayload() *models.RateMovingAverage {
	return o.Payload
}

func (o *StorageProxyMetricsWriteTimeoutsRatesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RateMovingAverage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageProxyMetricsWriteTimeoutsRatesGetDefault creates a StorageProxyMetricsWriteTimeoutsRatesGetDefault with default headers values
func NewStorageProxyMetricsWriteTimeoutsRatesGetDefault(code int) *StorageProxyMetricsWriteTimeoutsRatesGetDefault {
	return &StorageProxyMetricsWriteTimeoutsRatesGetDefault{
		_statusCode: code,
	}
}

/*StorageProxyMetricsWriteTimeoutsRatesGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyMetricsWriteTimeoutsRatesGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy metrics write timeouts rates get default response
func (o *StorageProxyMetricsWriteTimeoutsRatesGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyMetricsWriteTimeoutsRatesGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyMetricsWriteTimeoutsRatesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyMetricsWriteTimeoutsRatesGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
