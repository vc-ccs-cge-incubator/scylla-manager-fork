// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetReader is a Reader for the ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGet structure.
type ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK creates a ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK with default headers values
func NewColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK() *ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK {
	return &ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK{}
}

/*ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK handles this case with default header values.

ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK column family metrics write latency estimated histogram by name get o k
*/
type ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK struct {
}

func (o *ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/write_latency/estimated_histogram/{name}][%d] columnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK ", 200)
}

func (o *ColumnFamilyMetricsWriteLatencyEstimatedHistogramByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
