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

// StorageServiceOperationModeGetReader is a Reader for the StorageServiceOperationModeGet structure.
type StorageServiceOperationModeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceOperationModeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceOperationModeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageServiceOperationModeGetOK creates a StorageServiceOperationModeGetOK with default headers values
func NewStorageServiceOperationModeGetOK() *StorageServiceOperationModeGetOK {
	return &StorageServiceOperationModeGetOK{}
}

/*StorageServiceOperationModeGetOK handles this case with default header values.

StorageServiceOperationModeGetOK storage service operation mode get o k
*/
type StorageServiceOperationModeGetOK struct {
	Payload string
}

func (o *StorageServiceOperationModeGetOK) Error() string {
	return fmt.Sprintf("[GET /storage_service/operation_mode][%d] storageServiceOperationModeGetOK  %+v", 200, o.Payload)
}

func (o *StorageServiceOperationModeGetOK) GetPayload() string {
	return o.Payload
}

func (o *StorageServiceOperationModeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
