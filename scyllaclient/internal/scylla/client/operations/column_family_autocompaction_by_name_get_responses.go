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

// ColumnFamilyAutocompactionByNameGetReader is a Reader for the ColumnFamilyAutocompactionByNameGet structure.
type ColumnFamilyAutocompactionByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyAutocompactionByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyAutocompactionByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyAutocompactionByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyAutocompactionByNameGetOK creates a ColumnFamilyAutocompactionByNameGetOK with default headers values
func NewColumnFamilyAutocompactionByNameGetOK() *ColumnFamilyAutocompactionByNameGetOK {
	return &ColumnFamilyAutocompactionByNameGetOK{}
}

/*ColumnFamilyAutocompactionByNameGetOK handles this case with default header values.

ColumnFamilyAutocompactionByNameGetOK column family autocompaction by name get o k
*/
type ColumnFamilyAutocompactionByNameGetOK struct {
	Payload bool
}

func (o *ColumnFamilyAutocompactionByNameGetOK) GetPayload() bool {
	return o.Payload
}

func (o *ColumnFamilyAutocompactionByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyAutocompactionByNameGetDefault creates a ColumnFamilyAutocompactionByNameGetDefault with default headers values
func NewColumnFamilyAutocompactionByNameGetDefault(code int) *ColumnFamilyAutocompactionByNameGetDefault {
	return &ColumnFamilyAutocompactionByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyAutocompactionByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyAutocompactionByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family autocompaction by name get default response
func (o *ColumnFamilyAutocompactionByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyAutocompactionByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyAutocompactionByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyAutocompactionByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
