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

// CommitlogMetricsPendingTasksGetReader is a Reader for the CommitlogMetricsPendingTasksGet structure.
type CommitlogMetricsPendingTasksGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CommitlogMetricsPendingTasksGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCommitlogMetricsPendingTasksGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCommitlogMetricsPendingTasksGetOK creates a CommitlogMetricsPendingTasksGetOK with default headers values
func NewCommitlogMetricsPendingTasksGetOK() *CommitlogMetricsPendingTasksGetOK {
	return &CommitlogMetricsPendingTasksGetOK{}
}

/*CommitlogMetricsPendingTasksGetOK handles this case with default header values.

CommitlogMetricsPendingTasksGetOK commitlog metrics pending tasks get o k
*/
type CommitlogMetricsPendingTasksGetOK struct {
	Payload interface{}
}

func (o *CommitlogMetricsPendingTasksGetOK) Error() string {
	return fmt.Sprintf("[GET /commitlog/metrics/pending_tasks][%d] commitlogMetricsPendingTasksGetOK  %+v", 200, o.Payload)
}

func (o *CommitlogMetricsPendingTasksGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CommitlogMetricsPendingTasksGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
