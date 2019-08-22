// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/rclone/models"
)

// CoreGroupListReader is a Reader for the CoreGroupList structure.
type CoreGroupListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CoreGroupListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCoreGroupListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCoreGroupListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCoreGroupListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCoreGroupListOK creates a CoreGroupListOK with default headers values
func NewCoreGroupListOK() *CoreGroupListOK {
	return &CoreGroupListOK{}
}

/*CoreGroupListOK handles this case with default header values.

Current groups
*/
type CoreGroupListOK struct {
	Payload *models.GroupList
}

func (o *CoreGroupListOK) Error() string {
	return fmt.Sprintf("[POST /core/group-list][%d] coreGroupListOK  %+v", 200, o.Payload)
}

func (o *CoreGroupListOK) GetPayload() *models.GroupList {
	return o.Payload
}

func (o *CoreGroupListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCoreGroupListNotFound creates a CoreGroupListNotFound with default headers values
func NewCoreGroupListNotFound() *CoreGroupListNotFound {
	return &CoreGroupListNotFound{}
}

/*CoreGroupListNotFound handles this case with default header values.

Not found
*/
type CoreGroupListNotFound struct {
	Payload *models.ErrorResponse
}

func (o *CoreGroupListNotFound) Error() string {
	return fmt.Sprintf("[POST /core/group-list][%d] coreGroupListNotFound  %+v", 404, o.Payload)
}

func (o *CoreGroupListNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CoreGroupListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCoreGroupListInternalServerError creates a CoreGroupListInternalServerError with default headers values
func NewCoreGroupListInternalServerError() *CoreGroupListInternalServerError {
	return &CoreGroupListInternalServerError{}
}

/*CoreGroupListInternalServerError handles this case with default header values.

Server error
*/
type CoreGroupListInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CoreGroupListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /core/group-list][%d] coreGroupListInternalServerError  %+v", 500, o.Payload)
}

func (o *CoreGroupListInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CoreGroupListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
