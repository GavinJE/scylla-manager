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

// ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetReader is a Reader for the ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGet structure.
type ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetOK creates a ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetOK with default headers values
func NewColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetOK() *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetOK {
	return &ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetOK{}
}

/*ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetOK handles this case with default header values.

ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetOK column family metrics recent bloom filter false positives by name get o k
*/
type ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetDefault creates a ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetDefault with default headers values
func NewColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetDefault(code int) *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetDefault {
	return &ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics recent bloom filter false positives by name get default response
func (o *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsRecentBloomFilterFalsePositivesByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
