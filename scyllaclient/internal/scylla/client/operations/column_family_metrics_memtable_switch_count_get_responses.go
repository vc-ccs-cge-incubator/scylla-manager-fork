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

// ColumnFamilyMetricsMemtableSwitchCountGetReader is a Reader for the ColumnFamilyMetricsMemtableSwitchCountGet structure.
type ColumnFamilyMetricsMemtableSwitchCountGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsMemtableSwitchCountGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsMemtableSwitchCountGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsMemtableSwitchCountGetOK creates a ColumnFamilyMetricsMemtableSwitchCountGetOK with default headers values
func NewColumnFamilyMetricsMemtableSwitchCountGetOK() *ColumnFamilyMetricsMemtableSwitchCountGetOK {
	return &ColumnFamilyMetricsMemtableSwitchCountGetOK{}
}

/*ColumnFamilyMetricsMemtableSwitchCountGetOK handles this case with default header values.

ColumnFamilyMetricsMemtableSwitchCountGetOK column family metrics memtable switch count get o k
*/
type ColumnFamilyMetricsMemtableSwitchCountGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsMemtableSwitchCountGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/memtable_switch_count][%d] columnFamilyMetricsMemtableSwitchCountGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsMemtableSwitchCountGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *ColumnFamilyMetricsMemtableSwitchCountGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
