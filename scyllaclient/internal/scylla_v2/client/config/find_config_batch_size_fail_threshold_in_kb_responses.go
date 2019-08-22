// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/scylla_v2/models"
)

// FindConfigBatchSizeFailThresholdInKbReader is a Reader for the FindConfigBatchSizeFailThresholdInKb structure.
type FindConfigBatchSizeFailThresholdInKbReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigBatchSizeFailThresholdInKbReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigBatchSizeFailThresholdInKbOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigBatchSizeFailThresholdInKbDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigBatchSizeFailThresholdInKbOK creates a FindConfigBatchSizeFailThresholdInKbOK with default headers values
func NewFindConfigBatchSizeFailThresholdInKbOK() *FindConfigBatchSizeFailThresholdInKbOK {
	return &FindConfigBatchSizeFailThresholdInKbOK{}
}

/*FindConfigBatchSizeFailThresholdInKbOK handles this case with default header values.

Config value
*/
type FindConfigBatchSizeFailThresholdInKbOK struct {
	Payload int64
}

func (o *FindConfigBatchSizeFailThresholdInKbOK) Error() string {
	return fmt.Sprintf("[GET /config/batch_size_fail_threshold_in_kb][%d] findConfigBatchSizeFailThresholdInKbOK  %+v", 200, o.Payload)
}

func (o *FindConfigBatchSizeFailThresholdInKbOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigBatchSizeFailThresholdInKbOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigBatchSizeFailThresholdInKbDefault creates a FindConfigBatchSizeFailThresholdInKbDefault with default headers values
func NewFindConfigBatchSizeFailThresholdInKbDefault(code int) *FindConfigBatchSizeFailThresholdInKbDefault {
	return &FindConfigBatchSizeFailThresholdInKbDefault{
		_statusCode: code,
	}
}

/*FindConfigBatchSizeFailThresholdInKbDefault handles this case with default header values.

unexpected error
*/
type FindConfigBatchSizeFailThresholdInKbDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config batch size fail threshold in kb default response
func (o *FindConfigBatchSizeFailThresholdInKbDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigBatchSizeFailThresholdInKbDefault) Error() string {
	return fmt.Sprintf("[GET /config/batch_size_fail_threshold_in_kb][%d] find_config_batch_size_fail_threshold_in_kb default  %+v", o._statusCode, o.Payload)
}

func (o *FindConfigBatchSizeFailThresholdInKbDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigBatchSizeFailThresholdInKbDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
