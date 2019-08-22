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

// StorageProxyMaxHintWindowGetReader is a Reader for the StorageProxyMaxHintWindowGet structure.
type StorageProxyMaxHintWindowGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMaxHintWindowGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyMaxHintWindowGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageProxyMaxHintWindowGetOK creates a StorageProxyMaxHintWindowGetOK with default headers values
func NewStorageProxyMaxHintWindowGetOK() *StorageProxyMaxHintWindowGetOK {
	return &StorageProxyMaxHintWindowGetOK{}
}

/*StorageProxyMaxHintWindowGetOK handles this case with default header values.

StorageProxyMaxHintWindowGetOK storage proxy max hint window get o k
*/
type StorageProxyMaxHintWindowGetOK struct {
	Payload int32
}

func (o *StorageProxyMaxHintWindowGetOK) Error() string {
	return fmt.Sprintf("[GET /storage_proxy/max_hint_window][%d] storageProxyMaxHintWindowGetOK  %+v", 200, o.Payload)
}

func (o *StorageProxyMaxHintWindowGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *StorageProxyMaxHintWindowGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
