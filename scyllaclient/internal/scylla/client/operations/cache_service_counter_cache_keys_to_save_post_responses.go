// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CacheServiceCounterCacheKeysToSavePostReader is a Reader for the CacheServiceCounterCacheKeysToSavePost structure.
type CacheServiceCounterCacheKeysToSavePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceCounterCacheKeysToSavePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceCounterCacheKeysToSavePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCacheServiceCounterCacheKeysToSavePostOK creates a CacheServiceCounterCacheKeysToSavePostOK with default headers values
func NewCacheServiceCounterCacheKeysToSavePostOK() *CacheServiceCounterCacheKeysToSavePostOK {
	return &CacheServiceCounterCacheKeysToSavePostOK{}
}

/*CacheServiceCounterCacheKeysToSavePostOK handles this case with default header values.

CacheServiceCounterCacheKeysToSavePostOK cache service counter cache keys to save post o k
*/
type CacheServiceCounterCacheKeysToSavePostOK struct {
}

func (o *CacheServiceCounterCacheKeysToSavePostOK) Error() string {
	return fmt.Sprintf("[POST /cache_service/counter_cache_keys_to_save][%d] cacheServiceCounterCacheKeysToSavePostOK ", 200)
}

func (o *CacheServiceCounterCacheKeysToSavePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
