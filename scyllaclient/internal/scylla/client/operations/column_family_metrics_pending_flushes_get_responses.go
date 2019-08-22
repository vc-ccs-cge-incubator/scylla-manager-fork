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

// ColumnFamilyMetricsPendingFlushesGetReader is a Reader for the ColumnFamilyMetricsPendingFlushesGet structure.
type ColumnFamilyMetricsPendingFlushesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsPendingFlushesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsPendingFlushesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsPendingFlushesGetOK creates a ColumnFamilyMetricsPendingFlushesGetOK with default headers values
func NewColumnFamilyMetricsPendingFlushesGetOK() *ColumnFamilyMetricsPendingFlushesGetOK {
	return &ColumnFamilyMetricsPendingFlushesGetOK{}
}

/*ColumnFamilyMetricsPendingFlushesGetOK handles this case with default header values.

ColumnFamilyMetricsPendingFlushesGetOK column family metrics pending flushes get o k
*/
type ColumnFamilyMetricsPendingFlushesGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsPendingFlushesGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/pending_flushes][%d] columnFamilyMetricsPendingFlushesGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsPendingFlushesGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *ColumnFamilyMetricsPendingFlushesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
