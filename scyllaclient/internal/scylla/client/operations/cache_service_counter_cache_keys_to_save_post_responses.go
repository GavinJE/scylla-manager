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

// CacheServiceCounterCacheKeysToSavePostReader is a Reader for the CacheServiceCounterCacheKeysToSavePost structure.
type CacheServiceCounterCacheKeysToSavePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceCounterCacheKeysToSavePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceCounterCacheKeysToSavePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceCounterCacheKeysToSavePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceCounterCacheKeysToSavePostOK creates a CacheServiceCounterCacheKeysToSavePostOK with default headers values
func NewCacheServiceCounterCacheKeysToSavePostOK() *CacheServiceCounterCacheKeysToSavePostOK {
	return &CacheServiceCounterCacheKeysToSavePostOK{}
}

/*CacheServiceCounterCacheKeysToSavePostOK handles this case with default header values.

CacheServiceCounterCacheKeysToSavePostOK cache service counter cache keys to save post o k
*/
type CacheServiceCounterCacheKeysToSavePostOK struct {
}

func (o *CacheServiceCounterCacheKeysToSavePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCacheServiceCounterCacheKeysToSavePostDefault creates a CacheServiceCounterCacheKeysToSavePostDefault with default headers values
func NewCacheServiceCounterCacheKeysToSavePostDefault(code int) *CacheServiceCounterCacheKeysToSavePostDefault {
	return &CacheServiceCounterCacheKeysToSavePostDefault{
		_statusCode: code,
	}
}

/*CacheServiceCounterCacheKeysToSavePostDefault handles this case with default header values.

internal server error
*/
type CacheServiceCounterCacheKeysToSavePostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service counter cache keys to save post default response
func (o *CacheServiceCounterCacheKeysToSavePostDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceCounterCacheKeysToSavePostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceCounterCacheKeysToSavePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceCounterCacheKeysToSavePostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
