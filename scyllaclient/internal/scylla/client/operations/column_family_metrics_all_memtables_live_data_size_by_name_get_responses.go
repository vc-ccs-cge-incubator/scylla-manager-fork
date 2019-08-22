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

// ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetReader is a Reader for the ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGet structure.
type ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK creates a ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK with default headers values
func NewColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK() *ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK {
	return &ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK{}
}

/*ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK handles this case with default header values.

ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK column family metrics all memtables live data size by name get o k
*/
type ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/all_memtables_live_data_size/{name}][%d] columnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsAllMemtablesLiveDataSizeByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
