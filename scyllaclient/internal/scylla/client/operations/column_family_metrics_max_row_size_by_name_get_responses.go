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

// ColumnFamilyMetricsMaxRowSizeByNameGetReader is a Reader for the ColumnFamilyMetricsMaxRowSizeByNameGet structure.
type ColumnFamilyMetricsMaxRowSizeByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsMaxRowSizeByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsMaxRowSizeByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsMaxRowSizeByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsMaxRowSizeByNameGetOK creates a ColumnFamilyMetricsMaxRowSizeByNameGetOK with default headers values
func NewColumnFamilyMetricsMaxRowSizeByNameGetOK() *ColumnFamilyMetricsMaxRowSizeByNameGetOK {
	return &ColumnFamilyMetricsMaxRowSizeByNameGetOK{}
}

/*ColumnFamilyMetricsMaxRowSizeByNameGetOK handles this case with default header values.

ColumnFamilyMetricsMaxRowSizeByNameGetOK column family metrics max row size by name get o k
*/
type ColumnFamilyMetricsMaxRowSizeByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsMaxRowSizeByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsMaxRowSizeByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsMaxRowSizeByNameGetDefault creates a ColumnFamilyMetricsMaxRowSizeByNameGetDefault with default headers values
func NewColumnFamilyMetricsMaxRowSizeByNameGetDefault(code int) *ColumnFamilyMetricsMaxRowSizeByNameGetDefault {
	return &ColumnFamilyMetricsMaxRowSizeByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsMaxRowSizeByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsMaxRowSizeByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics max row size by name get default response
func (o *ColumnFamilyMetricsMaxRowSizeByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsMaxRowSizeByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsMaxRowSizeByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsMaxRowSizeByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
