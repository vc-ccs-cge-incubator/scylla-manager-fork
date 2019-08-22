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

// StreamManagerMetricsIncomingGetReader is a Reader for the StreamManagerMetricsIncomingGet structure.
type StreamManagerMetricsIncomingGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StreamManagerMetricsIncomingGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStreamManagerMetricsIncomingGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStreamManagerMetricsIncomingGetOK creates a StreamManagerMetricsIncomingGetOK with default headers values
func NewStreamManagerMetricsIncomingGetOK() *StreamManagerMetricsIncomingGetOK {
	return &StreamManagerMetricsIncomingGetOK{}
}

/*StreamManagerMetricsIncomingGetOK handles this case with default header values.

StreamManagerMetricsIncomingGetOK stream manager metrics incoming get o k
*/
type StreamManagerMetricsIncomingGetOK struct {
	Payload int32
}

func (o *StreamManagerMetricsIncomingGetOK) Error() string {
	return fmt.Sprintf("[GET /stream_manager/metrics/incoming][%d] streamManagerMetricsIncomingGetOK  %+v", 200, o.Payload)
}

func (o *StreamManagerMetricsIncomingGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *StreamManagerMetricsIncomingGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
