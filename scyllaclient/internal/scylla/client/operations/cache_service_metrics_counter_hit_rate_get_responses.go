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

// CacheServiceMetricsCounterHitRateGetReader is a Reader for the CacheServiceMetricsCounterHitRateGet structure.
type CacheServiceMetricsCounterHitRateGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceMetricsCounterHitRateGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceMetricsCounterHitRateGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceMetricsCounterHitRateGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceMetricsCounterHitRateGetOK creates a CacheServiceMetricsCounterHitRateGetOK with default headers values
func NewCacheServiceMetricsCounterHitRateGetOK() *CacheServiceMetricsCounterHitRateGetOK {
	return &CacheServiceMetricsCounterHitRateGetOK{}
}

/*CacheServiceMetricsCounterHitRateGetOK handles this case with default header values.

CacheServiceMetricsCounterHitRateGetOK cache service metrics counter hit rate get o k
*/
type CacheServiceMetricsCounterHitRateGetOK struct {
	Payload interface{}
}

func (o *CacheServiceMetricsCounterHitRateGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CacheServiceMetricsCounterHitRateGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCacheServiceMetricsCounterHitRateGetDefault creates a CacheServiceMetricsCounterHitRateGetDefault with default headers values
func NewCacheServiceMetricsCounterHitRateGetDefault(code int) *CacheServiceMetricsCounterHitRateGetDefault {
	return &CacheServiceMetricsCounterHitRateGetDefault{
		_statusCode: code,
	}
}

/*CacheServiceMetricsCounterHitRateGetDefault handles this case with default header values.

internal server error
*/
type CacheServiceMetricsCounterHitRateGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service metrics counter hit rate get default response
func (o *CacheServiceMetricsCounterHitRateGetDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceMetricsCounterHitRateGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceMetricsCounterHitRateGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceMetricsCounterHitRateGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
