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

// ColumnFamilyMetricsMemtableOffHeapSizeByNameGetReader is a Reader for the ColumnFamilyMetricsMemtableOffHeapSizeByNameGet structure.
type ColumnFamilyMetricsMemtableOffHeapSizeByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsMemtableOffHeapSizeByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsMemtableOffHeapSizeByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsMemtableOffHeapSizeByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsMemtableOffHeapSizeByNameGetOK creates a ColumnFamilyMetricsMemtableOffHeapSizeByNameGetOK with default headers values
func NewColumnFamilyMetricsMemtableOffHeapSizeByNameGetOK() *ColumnFamilyMetricsMemtableOffHeapSizeByNameGetOK {
	return &ColumnFamilyMetricsMemtableOffHeapSizeByNameGetOK{}
}

/*ColumnFamilyMetricsMemtableOffHeapSizeByNameGetOK handles this case with default header values.

ColumnFamilyMetricsMemtableOffHeapSizeByNameGetOK column family metrics memtable off heap size by name get o k
*/
type ColumnFamilyMetricsMemtableOffHeapSizeByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsMemtableOffHeapSizeByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsMemtableOffHeapSizeByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsMemtableOffHeapSizeByNameGetDefault creates a ColumnFamilyMetricsMemtableOffHeapSizeByNameGetDefault with default headers values
func NewColumnFamilyMetricsMemtableOffHeapSizeByNameGetDefault(code int) *ColumnFamilyMetricsMemtableOffHeapSizeByNameGetDefault {
	return &ColumnFamilyMetricsMemtableOffHeapSizeByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsMemtableOffHeapSizeByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsMemtableOffHeapSizeByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics memtable off heap size by name get default response
func (o *ColumnFamilyMetricsMemtableOffHeapSizeByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsMemtableOffHeapSizeByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsMemtableOffHeapSizeByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsMemtableOffHeapSizeByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
