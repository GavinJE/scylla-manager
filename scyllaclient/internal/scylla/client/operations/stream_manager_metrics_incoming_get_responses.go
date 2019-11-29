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

// StreamManagerMetricsIncomingGetReader is a Reader for the StreamManagerMetricsIncomingGet structure.
type StreamManagerMetricsIncomingGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StreamManagerMetricsIncomingGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStreamManagerMetricsIncomingGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStreamManagerMetricsIncomingGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStreamManagerMetricsIncomingGetOK creates a StreamManagerMetricsIncomingGetOK with default headers values
func NewStreamManagerMetricsIncomingGetOK() *StreamManagerMetricsIncomingGetOK {
	return &StreamManagerMetricsIncomingGetOK{}
}

/*StreamManagerMetricsIncomingGetOK handles this case with default header values.

StreamManagerMetricsIncomingGetOK stream manager metrics incoming get o k
*/
type StreamManagerMetricsIncomingGetOK struct {
	Payload int32
}

func (o *StreamManagerMetricsIncomingGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *StreamManagerMetricsIncomingGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStreamManagerMetricsIncomingGetDefault creates a StreamManagerMetricsIncomingGetDefault with default headers values
func NewStreamManagerMetricsIncomingGetDefault(code int) *StreamManagerMetricsIncomingGetDefault {
	return &StreamManagerMetricsIncomingGetDefault{
		_statusCode: code,
	}
}

/*StreamManagerMetricsIncomingGetDefault handles this case with default header values.

internal server error
*/
type StreamManagerMetricsIncomingGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the stream manager metrics incoming get default response
func (o *StreamManagerMetricsIncomingGetDefault) Code() int {
	return o._statusCode
}

func (o *StreamManagerMetricsIncomingGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StreamManagerMetricsIncomingGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StreamManagerMetricsIncomingGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
