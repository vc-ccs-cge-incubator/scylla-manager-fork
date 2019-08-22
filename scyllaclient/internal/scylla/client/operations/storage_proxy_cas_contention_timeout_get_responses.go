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

// StorageProxyCasContentionTimeoutGetReader is a Reader for the StorageProxyCasContentionTimeoutGet structure.
type StorageProxyCasContentionTimeoutGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyCasContentionTimeoutGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyCasContentionTimeoutGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageProxyCasContentionTimeoutGetOK creates a StorageProxyCasContentionTimeoutGetOK with default headers values
func NewStorageProxyCasContentionTimeoutGetOK() *StorageProxyCasContentionTimeoutGetOK {
	return &StorageProxyCasContentionTimeoutGetOK{}
}

/*StorageProxyCasContentionTimeoutGetOK handles this case with default header values.

StorageProxyCasContentionTimeoutGetOK storage proxy cas contention timeout get o k
*/
type StorageProxyCasContentionTimeoutGetOK struct {
	Payload interface{}
}

func (o *StorageProxyCasContentionTimeoutGetOK) Error() string {
	return fmt.Sprintf("[GET /storage_proxy/cas_contention_timeout][%d] storageProxyCasContentionTimeoutGetOK  %+v", 200, o.Payload)
}

func (o *StorageProxyCasContentionTimeoutGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *StorageProxyCasContentionTimeoutGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
