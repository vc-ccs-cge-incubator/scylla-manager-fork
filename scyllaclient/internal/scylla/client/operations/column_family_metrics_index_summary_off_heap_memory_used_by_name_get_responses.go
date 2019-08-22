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

// ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedByNameGetReader is a Reader for the ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedByNameGet structure.
type ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedByNameGetOK creates a ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedByNameGetOK with default headers values
func NewColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedByNameGetOK() *ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedByNameGetOK {
	return &ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedByNameGetOK{}
}

/*ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedByNameGetOK handles this case with default header values.

ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedByNameGetOK column family metrics index summary off heap memory used by name get o k
*/
type ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/index_summary_off_heap_memory_used/{name}][%d] columnFamilyMetricsIndexSummaryOffHeapMemoryUsedByNameGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsIndexSummaryOffHeapMemoryUsedByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
