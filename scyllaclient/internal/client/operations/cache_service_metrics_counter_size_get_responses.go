// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CacheServiceMetricsCounterSizeGetReader is a Reader for the CacheServiceMetricsCounterSizeGet structure.
type CacheServiceMetricsCounterSizeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceMetricsCounterSizeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCacheServiceMetricsCounterSizeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCacheServiceMetricsCounterSizeGetOK creates a CacheServiceMetricsCounterSizeGetOK with default headers values
func NewCacheServiceMetricsCounterSizeGetOK() *CacheServiceMetricsCounterSizeGetOK {
	return &CacheServiceMetricsCounterSizeGetOK{}
}

/*CacheServiceMetricsCounterSizeGetOK handles this case with default header values.

CacheServiceMetricsCounterSizeGetOK cache service metrics counter size get o k
*/
type CacheServiceMetricsCounterSizeGetOK struct {
	Payload interface{}
}

func (o *CacheServiceMetricsCounterSizeGetOK) Error() string {
	return fmt.Sprintf("[GET /cache_service/metrics/counter/size][%d] cacheServiceMetricsCounterSizeGetOK  %+v", 200, o.Payload)
}

func (o *CacheServiceMetricsCounterSizeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}