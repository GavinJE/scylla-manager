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

// ColumnFamilyMetricsMinRowSizeByNameGetReader is a Reader for the ColumnFamilyMetricsMinRowSizeByNameGet structure.
type ColumnFamilyMetricsMinRowSizeByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsMinRowSizeByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsMinRowSizeByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsMinRowSizeByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsMinRowSizeByNameGetOK creates a ColumnFamilyMetricsMinRowSizeByNameGetOK with default headers values
func NewColumnFamilyMetricsMinRowSizeByNameGetOK() *ColumnFamilyMetricsMinRowSizeByNameGetOK {
	return &ColumnFamilyMetricsMinRowSizeByNameGetOK{}
}

/*ColumnFamilyMetricsMinRowSizeByNameGetOK handles this case with default header values.

ColumnFamilyMetricsMinRowSizeByNameGetOK column family metrics min row size by name get o k
*/
type ColumnFamilyMetricsMinRowSizeByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsMinRowSizeByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsMinRowSizeByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsMinRowSizeByNameGetDefault creates a ColumnFamilyMetricsMinRowSizeByNameGetDefault with default headers values
func NewColumnFamilyMetricsMinRowSizeByNameGetDefault(code int) *ColumnFamilyMetricsMinRowSizeByNameGetDefault {
	return &ColumnFamilyMetricsMinRowSizeByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsMinRowSizeByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsMinRowSizeByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics min row size by name get default response
func (o *ColumnFamilyMetricsMinRowSizeByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsMinRowSizeByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsMinRowSizeByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsMinRowSizeByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
