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

// ColumnFamilyCrcCheckChanceByNamePostReader is a Reader for the ColumnFamilyCrcCheckChanceByNamePost structure.
type ColumnFamilyCrcCheckChanceByNamePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyCrcCheckChanceByNamePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyCrcCheckChanceByNamePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyCrcCheckChanceByNamePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyCrcCheckChanceByNamePostOK creates a ColumnFamilyCrcCheckChanceByNamePostOK with default headers values
func NewColumnFamilyCrcCheckChanceByNamePostOK() *ColumnFamilyCrcCheckChanceByNamePostOK {
	return &ColumnFamilyCrcCheckChanceByNamePostOK{}
}

/*ColumnFamilyCrcCheckChanceByNamePostOK handles this case with default header values.

ColumnFamilyCrcCheckChanceByNamePostOK column family crc check chance by name post o k
*/
type ColumnFamilyCrcCheckChanceByNamePostOK struct {
}

func (o *ColumnFamilyCrcCheckChanceByNamePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewColumnFamilyCrcCheckChanceByNamePostDefault creates a ColumnFamilyCrcCheckChanceByNamePostDefault with default headers values
func NewColumnFamilyCrcCheckChanceByNamePostDefault(code int) *ColumnFamilyCrcCheckChanceByNamePostDefault {
	return &ColumnFamilyCrcCheckChanceByNamePostDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyCrcCheckChanceByNamePostDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyCrcCheckChanceByNamePostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family crc check chance by name post default response
func (o *ColumnFamilyCrcCheckChanceByNamePostDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyCrcCheckChanceByNamePostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyCrcCheckChanceByNamePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyCrcCheckChanceByNamePostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
