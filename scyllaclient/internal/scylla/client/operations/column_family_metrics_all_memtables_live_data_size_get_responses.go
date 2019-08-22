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

// ColumnFamilyMetricsAllMemtablesLiveDataSizeGetReader is a Reader for the ColumnFamilyMetricsAllMemtablesLiveDataSizeGet structure.
type ColumnFamilyMetricsAllMemtablesLiveDataSizeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsAllMemtablesLiveDataSizeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsAllMemtablesLiveDataSizeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsAllMemtablesLiveDataSizeGetOK creates a ColumnFamilyMetricsAllMemtablesLiveDataSizeGetOK with default headers values
func NewColumnFamilyMetricsAllMemtablesLiveDataSizeGetOK() *ColumnFamilyMetricsAllMemtablesLiveDataSizeGetOK {
	return &ColumnFamilyMetricsAllMemtablesLiveDataSizeGetOK{}
}

/*ColumnFamilyMetricsAllMemtablesLiveDataSizeGetOK handles this case with default header values.

ColumnFamilyMetricsAllMemtablesLiveDataSizeGetOK column family metrics all memtables live data size get o k
*/
type ColumnFamilyMetricsAllMemtablesLiveDataSizeGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsAllMemtablesLiveDataSizeGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/all_memtables_live_data_size][%d] columnFamilyMetricsAllMemtablesLiveDataSizeGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsAllMemtablesLiveDataSizeGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsAllMemtablesLiveDataSizeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
