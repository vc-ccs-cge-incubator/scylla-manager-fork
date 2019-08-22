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

// FindConfigAutoSnapshotReader is a Reader for the FindConfigAutoSnapshot structure.
type FindConfigAutoSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigAutoSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigAutoSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigAutoSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigAutoSnapshotOK creates a FindConfigAutoSnapshotOK with default headers values
func NewFindConfigAutoSnapshotOK() *FindConfigAutoSnapshotOK {
	return &FindConfigAutoSnapshotOK{}
}

/*FindConfigAutoSnapshotOK handles this case with default header values.

Config value
*/
type FindConfigAutoSnapshotOK struct {
	Payload bool
}

func (o *FindConfigAutoSnapshotOK) Error() string {
	return fmt.Sprintf("[GET /config/auto_snapshot][%d] findConfigAutoSnapshotOK  %+v", 200, o.Payload)
}

func (o *FindConfigAutoSnapshotOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigAutoSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigAutoSnapshotDefault creates a FindConfigAutoSnapshotDefault with default headers values
func NewFindConfigAutoSnapshotDefault(code int) *FindConfigAutoSnapshotDefault {
	return &FindConfigAutoSnapshotDefault{
		_statusCode: code,
	}
}

/*FindConfigAutoSnapshotDefault handles this case with default header values.

unexpected error
*/
type FindConfigAutoSnapshotDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config auto snapshot default response
func (o *FindConfigAutoSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigAutoSnapshotDefault) Error() string {
	return fmt.Sprintf("[GET /config/auto_snapshot][%d] find_config_auto_snapshot default  %+v", o._statusCode, o.Payload)
}

func (o *FindConfigAutoSnapshotDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigAutoSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
