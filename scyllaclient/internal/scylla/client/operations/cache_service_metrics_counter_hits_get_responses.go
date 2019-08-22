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

// CacheServiceMetricsCounterHitsGetReader is a Reader for the CacheServiceMetricsCounterHitsGet structure.
type CacheServiceMetricsCounterHitsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceMetricsCounterHitsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceMetricsCounterHitsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCacheServiceMetricsCounterHitsGetOK creates a CacheServiceMetricsCounterHitsGetOK with default headers values
func NewCacheServiceMetricsCounterHitsGetOK() *CacheServiceMetricsCounterHitsGetOK {
	return &CacheServiceMetricsCounterHitsGetOK{}
}

/*CacheServiceMetricsCounterHitsGetOK handles this case with default header values.

CacheServiceMetricsCounterHitsGetOK cache service metrics counter hits get o k
*/
type CacheServiceMetricsCounterHitsGetOK struct {
	Payload interface{}
}

func (o *CacheServiceMetricsCounterHitsGetOK) Error() string {
	return fmt.Sprintf("[GET /cache_service/metrics/counter/hits][%d] cacheServiceMetricsCounterHitsGetOK  %+v", 200, o.Payload)
}

func (o *CacheServiceMetricsCounterHitsGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CacheServiceMetricsCounterHitsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
