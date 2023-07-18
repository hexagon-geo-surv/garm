// Code generated by go-swagger; DO NOT EDIT.

package repositories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	apiserver_params "github.com/cloudbase/garm/apiserver/params"
)

// DeleteRepoPoolReader is a Reader for the DeleteRepoPool structure.
type DeleteRepoPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRepoPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeleteRepoPoolDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeleteRepoPoolDefault creates a DeleteRepoPoolDefault with default headers values
func NewDeleteRepoPoolDefault(code int) *DeleteRepoPoolDefault {
	return &DeleteRepoPoolDefault{
		_statusCode: code,
	}
}

/*
DeleteRepoPoolDefault describes a response with status code -1, with default header values.

APIErrorResponse
*/
type DeleteRepoPoolDefault struct {
	_statusCode int

	Payload apiserver_params.APIErrorResponse
}

// IsSuccess returns true when this delete repo pool default response has a 2xx status code
func (o *DeleteRepoPoolDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete repo pool default response has a 3xx status code
func (o *DeleteRepoPoolDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete repo pool default response has a 4xx status code
func (o *DeleteRepoPoolDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete repo pool default response has a 5xx status code
func (o *DeleteRepoPoolDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete repo pool default response a status code equal to that given
func (o *DeleteRepoPoolDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete repo pool default response
func (o *DeleteRepoPoolDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRepoPoolDefault) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{repoID}/pools/{poolID}][%d] DeleteRepoPool default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRepoPoolDefault) String() string {
	return fmt.Sprintf("[DELETE /repositories/{repoID}/pools/{poolID}][%d] DeleteRepoPool default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRepoPoolDefault) GetPayload() apiserver_params.APIErrorResponse {
	return o.Payload
}

func (o *DeleteRepoPoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}