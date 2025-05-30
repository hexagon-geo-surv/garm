// Code generated by go-swagger; DO NOT EDIT.

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	apiserver_params "github.com/cloudbase/garm/apiserver/params"
)

// DeleteGiteaCredentialsReader is a Reader for the DeleteGiteaCredentials structure.
type DeleteGiteaCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGiteaCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeleteGiteaCredentialsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeleteGiteaCredentialsDefault creates a DeleteGiteaCredentialsDefault with default headers values
func NewDeleteGiteaCredentialsDefault(code int) *DeleteGiteaCredentialsDefault {
	return &DeleteGiteaCredentialsDefault{
		_statusCode: code,
	}
}

/*
DeleteGiteaCredentialsDefault describes a response with status code -1, with default header values.

APIErrorResponse
*/
type DeleteGiteaCredentialsDefault struct {
	_statusCode int

	Payload apiserver_params.APIErrorResponse
}

// IsSuccess returns true when this delete gitea credentials default response has a 2xx status code
func (o *DeleteGiteaCredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete gitea credentials default response has a 3xx status code
func (o *DeleteGiteaCredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete gitea credentials default response has a 4xx status code
func (o *DeleteGiteaCredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete gitea credentials default response has a 5xx status code
func (o *DeleteGiteaCredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete gitea credentials default response a status code equal to that given
func (o *DeleteGiteaCredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete gitea credentials default response
func (o *DeleteGiteaCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *DeleteGiteaCredentialsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /gitea/credentials/{id}][%d] DeleteGiteaCredentials default %s", o._statusCode, payload)
}

func (o *DeleteGiteaCredentialsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /gitea/credentials/{id}][%d] DeleteGiteaCredentials default %s", o._statusCode, payload)
}

func (o *DeleteGiteaCredentialsDefault) GetPayload() apiserver_params.APIErrorResponse {
	return o.Payload
}

func (o *DeleteGiteaCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
