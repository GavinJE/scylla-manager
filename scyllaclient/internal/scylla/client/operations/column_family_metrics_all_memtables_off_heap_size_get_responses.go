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

// ColumnFamilyMetricsAllMemtablesOffHeapSizeGetReader is a Reader for the ColumnFamilyMetricsAllMemtablesOffHeapSizeGet structure.
type ColumnFamilyMetricsAllMemtablesOffHeapSizeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK creates a ColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK with default headers values
func NewColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK() *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK {
	return &ColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK{}
}

/*ColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK handles this case with default header values.

ColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK column family metrics all memtables off heap size get o k
*/
type ColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/all_memtables_off_heap_size][%d] columnFamilyMetricsAllMemtablesOffHeapSizeGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsAllMemtablesOffHeapSizeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}