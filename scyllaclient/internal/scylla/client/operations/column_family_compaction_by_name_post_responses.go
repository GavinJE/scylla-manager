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

// ColumnFamilyCompactionByNamePostReader is a Reader for the ColumnFamilyCompactionByNamePost structure.
type ColumnFamilyCompactionByNamePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyCompactionByNamePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyCompactionByNamePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyCompactionByNamePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyCompactionByNamePostOK creates a ColumnFamilyCompactionByNamePostOK with default headers values
func NewColumnFamilyCompactionByNamePostOK() *ColumnFamilyCompactionByNamePostOK {
	return &ColumnFamilyCompactionByNamePostOK{}
}

/*ColumnFamilyCompactionByNamePostOK handles this case with default header values.

ColumnFamilyCompactionByNamePostOK column family compaction by name post o k
*/
type ColumnFamilyCompactionByNamePostOK struct {
	Payload string
}

func (o *ColumnFamilyCompactionByNamePostOK) GetPayload() string {
	return o.Payload
}

func (o *ColumnFamilyCompactionByNamePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyCompactionByNamePostDefault creates a ColumnFamilyCompactionByNamePostDefault with default headers values
func NewColumnFamilyCompactionByNamePostDefault(code int) *ColumnFamilyCompactionByNamePostDefault {
	return &ColumnFamilyCompactionByNamePostDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyCompactionByNamePostDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyCompactionByNamePostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family compaction by name post default response
func (o *ColumnFamilyCompactionByNamePostDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyCompactionByNamePostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyCompactionByNamePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyCompactionByNamePostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
