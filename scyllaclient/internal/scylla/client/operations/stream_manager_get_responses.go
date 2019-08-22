// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/scylla/models"
)

// StreamManagerGetReader is a Reader for the StreamManagerGet structure.
type StreamManagerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StreamManagerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStreamManagerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStreamManagerGetOK creates a StreamManagerGetOK with default headers values
func NewStreamManagerGetOK() *StreamManagerGetOK {
	return &StreamManagerGetOK{}
}

/*StreamManagerGetOK handles this case with default header values.

StreamManagerGetOK stream manager get o k
*/
type StreamManagerGetOK struct {
	Payload []*models.StreamState
}

func (o *StreamManagerGetOK) Error() string {
	return fmt.Sprintf("[GET /stream_manager/][%d] streamManagerGetOK  %+v", 200, o.Payload)
}

func (o *StreamManagerGetOK) GetPayload() []*models.StreamState {
	return o.Payload
}

func (o *StreamManagerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
