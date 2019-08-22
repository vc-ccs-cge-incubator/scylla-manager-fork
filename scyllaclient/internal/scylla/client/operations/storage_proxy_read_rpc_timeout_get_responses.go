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

// StorageProxyReadRPCTimeoutGetReader is a Reader for the StorageProxyReadRPCTimeoutGet structure.
type StorageProxyReadRPCTimeoutGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyReadRPCTimeoutGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyReadRPCTimeoutGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageProxyReadRPCTimeoutGetOK creates a StorageProxyReadRPCTimeoutGetOK with default headers values
func NewStorageProxyReadRPCTimeoutGetOK() *StorageProxyReadRPCTimeoutGetOK {
	return &StorageProxyReadRPCTimeoutGetOK{}
}

/*StorageProxyReadRPCTimeoutGetOK handles this case with default header values.

StorageProxyReadRPCTimeoutGetOK storage proxy read Rpc timeout get o k
*/
type StorageProxyReadRPCTimeoutGetOK struct {
	Payload interface{}
}

func (o *StorageProxyReadRPCTimeoutGetOK) Error() string {
	return fmt.Sprintf("[GET /storage_proxy/read_rpc_timeout][%d] storageProxyReadRpcTimeoutGetOK  %+v", 200, o.Payload)
}

func (o *StorageProxyReadRPCTimeoutGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *StorageProxyReadRPCTimeoutGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
