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

// ColumnFamilyMinimumCompactionByNameGetReader is a Reader for the ColumnFamilyMinimumCompactionByNameGet structure.
type ColumnFamilyMinimumCompactionByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMinimumCompactionByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMinimumCompactionByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMinimumCompactionByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMinimumCompactionByNameGetOK creates a ColumnFamilyMinimumCompactionByNameGetOK with default headers values
func NewColumnFamilyMinimumCompactionByNameGetOK() *ColumnFamilyMinimumCompactionByNameGetOK {
	return &ColumnFamilyMinimumCompactionByNameGetOK{}
}

/*ColumnFamilyMinimumCompactionByNameGetOK handles this case with default header values.

ColumnFamilyMinimumCompactionByNameGetOK column family minimum compaction by name get o k
*/
type ColumnFamilyMinimumCompactionByNameGetOK struct {
	Payload string
}

func (o *ColumnFamilyMinimumCompactionByNameGetOK) GetPayload() string {
	return o.Payload
}

func (o *ColumnFamilyMinimumCompactionByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMinimumCompactionByNameGetDefault creates a ColumnFamilyMinimumCompactionByNameGetDefault with default headers values
func NewColumnFamilyMinimumCompactionByNameGetDefault(code int) *ColumnFamilyMinimumCompactionByNameGetDefault {
	return &ColumnFamilyMinimumCompactionByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMinimumCompactionByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMinimumCompactionByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family minimum compaction by name get default response
func (o *ColumnFamilyMinimumCompactionByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMinimumCompactionByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMinimumCompactionByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMinimumCompactionByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
