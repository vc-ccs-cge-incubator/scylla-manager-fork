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

// StorageServiceMetricsTotalHintsGetReader is a Reader for the StorageServiceMetricsTotalHintsGet structure.
type StorageServiceMetricsTotalHintsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceMetricsTotalHintsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceMetricsTotalHintsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageServiceMetricsTotalHintsGetOK creates a StorageServiceMetricsTotalHintsGetOK with default headers values
func NewStorageServiceMetricsTotalHintsGetOK() *StorageServiceMetricsTotalHintsGetOK {
	return &StorageServiceMetricsTotalHintsGetOK{}
}

/*StorageServiceMetricsTotalHintsGetOK handles this case with default header values.

StorageServiceMetricsTotalHintsGetOK storage service metrics total hints get o k
*/
type StorageServiceMetricsTotalHintsGetOK struct {
	Payload int32
}

func (o *StorageServiceMetricsTotalHintsGetOK) Error() string {
	return fmt.Sprintf("[GET /storage_service/metrics/total_hints][%d] storageServiceMetricsTotalHintsGetOK  %+v", 200, o.Payload)
}

func (o *StorageServiceMetricsTotalHintsGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *StorageServiceMetricsTotalHintsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
