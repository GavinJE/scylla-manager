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

// CacheServiceMetricsKeyRequestsMovingAvrageGetReader is a Reader for the CacheServiceMetricsKeyRequestsMovingAvrageGet structure.
type CacheServiceMetricsKeyRequestsMovingAvrageGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceMetricsKeyRequestsMovingAvrageGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceMetricsKeyRequestsMovingAvrageGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceMetricsKeyRequestsMovingAvrageGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceMetricsKeyRequestsMovingAvrageGetOK creates a CacheServiceMetricsKeyRequestsMovingAvrageGetOK with default headers values
func NewCacheServiceMetricsKeyRequestsMovingAvrageGetOK() *CacheServiceMetricsKeyRequestsMovingAvrageGetOK {
	return &CacheServiceMetricsKeyRequestsMovingAvrageGetOK{}
}

/*CacheServiceMetricsKeyRequestsMovingAvrageGetOK handles this case with default header values.

CacheServiceMetricsKeyRequestsMovingAvrageGetOK cache service metrics key requests moving avrage get o k
*/
type CacheServiceMetricsKeyRequestsMovingAvrageGetOK struct {
	Payload *models.RateMovingAverage
}

func (o *CacheServiceMetricsKeyRequestsMovingAvrageGetOK) GetPayload() *models.RateMovingAverage {
	return o.Payload
}

func (o *CacheServiceMetricsKeyRequestsMovingAvrageGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RateMovingAverage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCacheServiceMetricsKeyRequestsMovingAvrageGetDefault creates a CacheServiceMetricsKeyRequestsMovingAvrageGetDefault with default headers values
func NewCacheServiceMetricsKeyRequestsMovingAvrageGetDefault(code int) *CacheServiceMetricsKeyRequestsMovingAvrageGetDefault {
	return &CacheServiceMetricsKeyRequestsMovingAvrageGetDefault{
		_statusCode: code,
	}
}

/*CacheServiceMetricsKeyRequestsMovingAvrageGetDefault handles this case with default header values.

internal server error
*/
type CacheServiceMetricsKeyRequestsMovingAvrageGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service metrics key requests moving avrage get default response
func (o *CacheServiceMetricsKeyRequestsMovingAvrageGetDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceMetricsKeyRequestsMovingAvrageGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceMetricsKeyRequestsMovingAvrageGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceMetricsKeyRequestsMovingAvrageGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
