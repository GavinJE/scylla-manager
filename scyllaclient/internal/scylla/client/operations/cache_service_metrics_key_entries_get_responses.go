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

// CacheServiceMetricsKeyEntriesGetReader is a Reader for the CacheServiceMetricsKeyEntriesGet structure.
type CacheServiceMetricsKeyEntriesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceMetricsKeyEntriesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceMetricsKeyEntriesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceMetricsKeyEntriesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceMetricsKeyEntriesGetOK creates a CacheServiceMetricsKeyEntriesGetOK with default headers values
func NewCacheServiceMetricsKeyEntriesGetOK() *CacheServiceMetricsKeyEntriesGetOK {
	return &CacheServiceMetricsKeyEntriesGetOK{}
}

/*CacheServiceMetricsKeyEntriesGetOK handles this case with default header values.

CacheServiceMetricsKeyEntriesGetOK cache service metrics key entries get o k
*/
type CacheServiceMetricsKeyEntriesGetOK struct {
	Payload int32
}

func (o *CacheServiceMetricsKeyEntriesGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *CacheServiceMetricsKeyEntriesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCacheServiceMetricsKeyEntriesGetDefault creates a CacheServiceMetricsKeyEntriesGetDefault with default headers values
func NewCacheServiceMetricsKeyEntriesGetDefault(code int) *CacheServiceMetricsKeyEntriesGetDefault {
	return &CacheServiceMetricsKeyEntriesGetDefault{
		_statusCode: code,
	}
}

/*CacheServiceMetricsKeyEntriesGetDefault handles this case with default header values.

internal server error
*/
type CacheServiceMetricsKeyEntriesGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service metrics key entries get default response
func (o *CacheServiceMetricsKeyEntriesGetDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceMetricsKeyEntriesGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceMetricsKeyEntriesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceMetricsKeyEntriesGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
