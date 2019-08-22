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

// FindConfigMemtableOffheapSpaceInMbReader is a Reader for the FindConfigMemtableOffheapSpaceInMb structure.
type FindConfigMemtableOffheapSpaceInMbReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigMemtableOffheapSpaceInMbReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigMemtableOffheapSpaceInMbOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigMemtableOffheapSpaceInMbDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigMemtableOffheapSpaceInMbOK creates a FindConfigMemtableOffheapSpaceInMbOK with default headers values
func NewFindConfigMemtableOffheapSpaceInMbOK() *FindConfigMemtableOffheapSpaceInMbOK {
	return &FindConfigMemtableOffheapSpaceInMbOK{}
}

/*FindConfigMemtableOffheapSpaceInMbOK handles this case with default header values.

Config value
*/
type FindConfigMemtableOffheapSpaceInMbOK struct {
	Payload int64
}

func (o *FindConfigMemtableOffheapSpaceInMbOK) Error() string {
	return fmt.Sprintf("[GET /config/memtable_offheap_space_in_mb][%d] findConfigMemtableOffheapSpaceInMbOK  %+v", 200, o.Payload)
}

func (o *FindConfigMemtableOffheapSpaceInMbOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigMemtableOffheapSpaceInMbOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigMemtableOffheapSpaceInMbDefault creates a FindConfigMemtableOffheapSpaceInMbDefault with default headers values
func NewFindConfigMemtableOffheapSpaceInMbDefault(code int) *FindConfigMemtableOffheapSpaceInMbDefault {
	return &FindConfigMemtableOffheapSpaceInMbDefault{
		_statusCode: code,
	}
}

/*FindConfigMemtableOffheapSpaceInMbDefault handles this case with default header values.

unexpected error
*/
type FindConfigMemtableOffheapSpaceInMbDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config memtable offheap space in mb default response
func (o *FindConfigMemtableOffheapSpaceInMbDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigMemtableOffheapSpaceInMbDefault) Error() string {
	return fmt.Sprintf("[GET /config/memtable_offheap_space_in_mb][%d] find_config_memtable_offheap_space_in_mb default  %+v", o._statusCode, o.Payload)
}

func (o *FindConfigMemtableOffheapSpaceInMbDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigMemtableOffheapSpaceInMbDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
