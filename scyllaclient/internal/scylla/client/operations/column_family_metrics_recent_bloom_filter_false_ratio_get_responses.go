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

// ColumnFamilyMetricsRecentBloomFilterFalseRatioGetReader is a Reader for the ColumnFamilyMetricsRecentBloomFilterFalseRatioGet structure.
type ColumnFamilyMetricsRecentBloomFilterFalseRatioGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsRecentBloomFilterFalseRatioGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK creates a ColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK with default headers values
func NewColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK() *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK {
	return &ColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK{}
}

/*ColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK handles this case with default header values.

ColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK column family metrics recent bloom filter false ratio get o k
*/
type ColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsRecentBloomFilterFalseRatioGetDefault creates a ColumnFamilyMetricsRecentBloomFilterFalseRatioGetDefault with default headers values
func NewColumnFamilyMetricsRecentBloomFilterFalseRatioGetDefault(code int) *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetDefault {
	return &ColumnFamilyMetricsRecentBloomFilterFalseRatioGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsRecentBloomFilterFalseRatioGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsRecentBloomFilterFalseRatioGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics recent bloom filter false ratio get default response
func (o *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
