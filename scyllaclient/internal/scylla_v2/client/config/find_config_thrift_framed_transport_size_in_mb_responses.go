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

// FindConfigThriftFramedTransportSizeInMbReader is a Reader for the FindConfigThriftFramedTransportSizeInMb structure.
type FindConfigThriftFramedTransportSizeInMbReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigThriftFramedTransportSizeInMbReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigThriftFramedTransportSizeInMbOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigThriftFramedTransportSizeInMbDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigThriftFramedTransportSizeInMbOK creates a FindConfigThriftFramedTransportSizeInMbOK with default headers values
func NewFindConfigThriftFramedTransportSizeInMbOK() *FindConfigThriftFramedTransportSizeInMbOK {
	return &FindConfigThriftFramedTransportSizeInMbOK{}
}

/*FindConfigThriftFramedTransportSizeInMbOK handles this case with default header values.

Config value
*/
type FindConfigThriftFramedTransportSizeInMbOK struct {
	Payload int64
}

func (o *FindConfigThriftFramedTransportSizeInMbOK) Error() string {
	return fmt.Sprintf("[GET /config/thrift_framed_transport_size_in_mb][%d] findConfigThriftFramedTransportSizeInMbOK  %+v", 200, o.Payload)
}

func (o *FindConfigThriftFramedTransportSizeInMbOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigThriftFramedTransportSizeInMbOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigThriftFramedTransportSizeInMbDefault creates a FindConfigThriftFramedTransportSizeInMbDefault with default headers values
func NewFindConfigThriftFramedTransportSizeInMbDefault(code int) *FindConfigThriftFramedTransportSizeInMbDefault {
	return &FindConfigThriftFramedTransportSizeInMbDefault{
		_statusCode: code,
	}
}

/*FindConfigThriftFramedTransportSizeInMbDefault handles this case with default header values.

unexpected error
*/
type FindConfigThriftFramedTransportSizeInMbDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config thrift framed transport size in mb default response
func (o *FindConfigThriftFramedTransportSizeInMbDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigThriftFramedTransportSizeInMbDefault) Error() string {
	return fmt.Sprintf("[GET /config/thrift_framed_transport_size_in_mb][%d] find_config_thrift_framed_transport_size_in_mb default  %+v", o._statusCode, o.Payload)
}

func (o *FindConfigThriftFramedTransportSizeInMbDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigThriftFramedTransportSizeInMbDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
