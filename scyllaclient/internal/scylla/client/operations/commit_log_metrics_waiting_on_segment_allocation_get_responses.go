// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CommitLogMetricsWaitingOnSegmentAllocationGetReader is a Reader for the CommitLogMetricsWaitingOnSegmentAllocationGet structure.
type CommitLogMetricsWaitingOnSegmentAllocationGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CommitLogMetricsWaitingOnSegmentAllocationGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCommitLogMetricsWaitingOnSegmentAllocationGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCommitLogMetricsWaitingOnSegmentAllocationGetOK creates a CommitLogMetricsWaitingOnSegmentAllocationGetOK with default headers values
func NewCommitLogMetricsWaitingOnSegmentAllocationGetOK() *CommitLogMetricsWaitingOnSegmentAllocationGetOK {
	return &CommitLogMetricsWaitingOnSegmentAllocationGetOK{}
}

/*CommitLogMetricsWaitingOnSegmentAllocationGetOK handles this case with default header values.

CommitLogMetricsWaitingOnSegmentAllocationGetOK commit log metrics waiting on segment allocation get o k
*/
type CommitLogMetricsWaitingOnSegmentAllocationGetOK struct {
}

func (o *CommitLogMetricsWaitingOnSegmentAllocationGetOK) Error() string {
	return fmt.Sprintf("[GET /commit_log/metrics/waiting_on_segment_allocation][%d] commitLogMetricsWaitingOnSegmentAllocationGetOK ", 200)
}

func (o *CommitLogMetricsWaitingOnSegmentAllocationGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
