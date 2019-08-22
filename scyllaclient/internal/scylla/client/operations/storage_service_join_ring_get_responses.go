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

// StorageServiceJoinRingGetReader is a Reader for the StorageServiceJoinRingGet structure.
type StorageServiceJoinRingGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceJoinRingGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceJoinRingGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageServiceJoinRingGetOK creates a StorageServiceJoinRingGetOK with default headers values
func NewStorageServiceJoinRingGetOK() *StorageServiceJoinRingGetOK {
	return &StorageServiceJoinRingGetOK{}
}

/*StorageServiceJoinRingGetOK handles this case with default header values.

StorageServiceJoinRingGetOK storage service join ring get o k
*/
type StorageServiceJoinRingGetOK struct {
	Payload bool
}

func (o *StorageServiceJoinRingGetOK) Error() string {
	return fmt.Sprintf("[GET /storage_service/join_ring][%d] storageServiceJoinRingGetOK  %+v", 200, o.Payload)
}

func (o *StorageServiceJoinRingGetOK) GetPayload() bool {
	return o.Payload
}

func (o *StorageServiceJoinRingGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
