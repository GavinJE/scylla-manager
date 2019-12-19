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

	models "github.com/scylladb/mermaid/pkg/scyllaclient/internal/scylla/models"
)

// ColumnFamilyMetricsMemtableLiveDataSizeByNameGetReader is a Reader for the ColumnFamilyMetricsMemtableLiveDataSizeByNameGet structure.
type ColumnFamilyMetricsMemtableLiveDataSizeByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsMemtableLiveDataSizeByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsMemtableLiveDataSizeByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsMemtableLiveDataSizeByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsMemtableLiveDataSizeByNameGetOK creates a ColumnFamilyMetricsMemtableLiveDataSizeByNameGetOK with default headers values
func NewColumnFamilyMetricsMemtableLiveDataSizeByNameGetOK() *ColumnFamilyMetricsMemtableLiveDataSizeByNameGetOK {
	return &ColumnFamilyMetricsMemtableLiveDataSizeByNameGetOK{}
}

/*ColumnFamilyMetricsMemtableLiveDataSizeByNameGetOK handles this case with default header values.

ColumnFamilyMetricsMemtableLiveDataSizeByNameGetOK column family metrics memtable live data size by name get o k
*/
type ColumnFamilyMetricsMemtableLiveDataSizeByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsMemtableLiveDataSizeByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsMemtableLiveDataSizeByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsMemtableLiveDataSizeByNameGetDefault creates a ColumnFamilyMetricsMemtableLiveDataSizeByNameGetDefault with default headers values
func NewColumnFamilyMetricsMemtableLiveDataSizeByNameGetDefault(code int) *ColumnFamilyMetricsMemtableLiveDataSizeByNameGetDefault {
	return &ColumnFamilyMetricsMemtableLiveDataSizeByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsMemtableLiveDataSizeByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsMemtableLiveDataSizeByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics memtable live data size by name get default response
func (o *ColumnFamilyMetricsMemtableLiveDataSizeByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsMemtableLiveDataSizeByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsMemtableLiveDataSizeByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsMemtableLiveDataSizeByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}