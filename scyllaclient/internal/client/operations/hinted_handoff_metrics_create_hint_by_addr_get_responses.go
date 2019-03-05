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

// HintedHandoffMetricsCreateHintByAddrGetReader is a Reader for the HintedHandoffMetricsCreateHintByAddrGet structure.
type HintedHandoffMetricsCreateHintByAddrGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HintedHandoffMetricsCreateHintByAddrGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewHintedHandoffMetricsCreateHintByAddrGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHintedHandoffMetricsCreateHintByAddrGetOK creates a HintedHandoffMetricsCreateHintByAddrGetOK with default headers values
func NewHintedHandoffMetricsCreateHintByAddrGetOK() *HintedHandoffMetricsCreateHintByAddrGetOK {
	return &HintedHandoffMetricsCreateHintByAddrGetOK{}
}

/*HintedHandoffMetricsCreateHintByAddrGetOK handles this case with default header values.

HintedHandoffMetricsCreateHintByAddrGetOK hinted handoff metrics create hint by addr get o k
*/
type HintedHandoffMetricsCreateHintByAddrGetOK struct {
	Payload int32
}

func (o *HintedHandoffMetricsCreateHintByAddrGetOK) Error() string {
	return fmt.Sprintf("[GET /hinted_handoff/metrics/create_hint/{addr}][%d] hintedHandoffMetricsCreateHintByAddrGetOK  %+v", 200, o.Payload)
}

func (o *HintedHandoffMetricsCreateHintByAddrGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}