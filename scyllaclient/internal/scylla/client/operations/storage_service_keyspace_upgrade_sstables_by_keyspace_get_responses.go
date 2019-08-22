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

// StorageServiceKeyspaceUpgradeSstablesByKeyspaceGetReader is a Reader for the StorageServiceKeyspaceUpgradeSstablesByKeyspaceGet structure.
type StorageServiceKeyspaceUpgradeSstablesByKeyspaceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceKeyspaceUpgradeSstablesByKeyspaceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceKeyspaceUpgradeSstablesByKeyspaceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageServiceKeyspaceUpgradeSstablesByKeyspaceGetOK creates a StorageServiceKeyspaceUpgradeSstablesByKeyspaceGetOK with default headers values
func NewStorageServiceKeyspaceUpgradeSstablesByKeyspaceGetOK() *StorageServiceKeyspaceUpgradeSstablesByKeyspaceGetOK {
	return &StorageServiceKeyspaceUpgradeSstablesByKeyspaceGetOK{}
}

/*StorageServiceKeyspaceUpgradeSstablesByKeyspaceGetOK handles this case with default header values.

StorageServiceKeyspaceUpgradeSstablesByKeyspaceGetOK storage service keyspace upgrade sstables by keyspace get o k
*/
type StorageServiceKeyspaceUpgradeSstablesByKeyspaceGetOK struct {
	Payload int32
}

func (o *StorageServiceKeyspaceUpgradeSstablesByKeyspaceGetOK) Error() string {
	return fmt.Sprintf("[GET /storage_service/keyspace_upgrade_sstables/{keyspace}][%d] storageServiceKeyspaceUpgradeSstablesByKeyspaceGetOK  %+v", 200, o.Payload)
}

func (o *StorageServiceKeyspaceUpgradeSstablesByKeyspaceGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *StorageServiceKeyspaceUpgradeSstablesByKeyspaceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
