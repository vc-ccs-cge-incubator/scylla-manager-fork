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

// StorageServiceRPCServerGetReader is a Reader for the StorageServiceRPCServerGet structure.
type StorageServiceRPCServerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceRPCServerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceRPCServerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageServiceRPCServerGetOK creates a StorageServiceRPCServerGetOK with default headers values
func NewStorageServiceRPCServerGetOK() *StorageServiceRPCServerGetOK {
	return &StorageServiceRPCServerGetOK{}
}

/*StorageServiceRPCServerGetOK handles this case with default header values.

StorageServiceRPCServerGetOK storage service Rpc server get o k
*/
type StorageServiceRPCServerGetOK struct {
	Payload bool
}

func (o *StorageServiceRPCServerGetOK) Error() string {
	return fmt.Sprintf("[GET /storage_service/rpc_server][%d] storageServiceRpcServerGetOK  %+v", 200, o.Payload)
}

func (o *StorageServiceRPCServerGetOK) GetPayload() bool {
	return o.Payload
}

func (o *StorageServiceRPCServerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
