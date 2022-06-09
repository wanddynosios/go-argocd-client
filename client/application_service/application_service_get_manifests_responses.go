// Code generated by go-swagger; DO NOT EDIT.

package application_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/wanddynosios/go-argocd-client/models"
)

// ApplicationServiceGetManifestsReader is a Reader for the ApplicationServiceGetManifests structure.
type ApplicationServiceGetManifestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationServiceGetManifestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationServiceGetManifestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationServiceGetManifestsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationServiceGetManifestsOK creates a ApplicationServiceGetManifestsOK with default headers values
func NewApplicationServiceGetManifestsOK() *ApplicationServiceGetManifestsOK {
	return &ApplicationServiceGetManifestsOK{}
}

/* ApplicationServiceGetManifestsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ApplicationServiceGetManifestsOK struct {
	Payload *models.RepositoryManifestResponse
}

func (o *ApplicationServiceGetManifestsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/applications/{name}/manifests][%d] applicationServiceGetManifestsOK  %+v", 200, o.Payload)
}
func (o *ApplicationServiceGetManifestsOK) GetPayload() *models.RepositoryManifestResponse {
	return o.Payload
}

func (o *ApplicationServiceGetManifestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RepositoryManifestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationServiceGetManifestsDefault creates a ApplicationServiceGetManifestsDefault with default headers values
func NewApplicationServiceGetManifestsDefault(code int) *ApplicationServiceGetManifestsDefault {
	return &ApplicationServiceGetManifestsDefault{
		_statusCode: code,
	}
}

/* ApplicationServiceGetManifestsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ApplicationServiceGetManifestsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the application service get manifests default response
func (o *ApplicationServiceGetManifestsDefault) Code() int {
	return o._statusCode
}

func (o *ApplicationServiceGetManifestsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/applications/{name}/manifests][%d] ApplicationService_GetManifests default  %+v", o._statusCode, o.Payload)
}
func (o *ApplicationServiceGetManifestsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ApplicationServiceGetManifestsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
