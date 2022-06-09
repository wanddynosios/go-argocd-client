// Code generated by go-swagger; DO NOT EDIT.

package repository_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/wanddynosios/go-argocd-client/models"
)

// RepositoryServiceGetReader is a Reader for the RepositoryServiceGet structure.
type RepositoryServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepositoryServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepositoryServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRepositoryServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRepositoryServiceGetOK creates a RepositoryServiceGetOK with default headers values
func NewRepositoryServiceGetOK() *RepositoryServiceGetOK {
	return &RepositoryServiceGetOK{}
}

/* RepositoryServiceGetOK describes a response with status code 200, with default header values.

A successful response.
*/
type RepositoryServiceGetOK struct {
	Payload *models.V1alpha1Repository
}

func (o *RepositoryServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/repositories/{repo}][%d] repositoryServiceGetOK  %+v", 200, o.Payload)
}
func (o *RepositoryServiceGetOK) GetPayload() *models.V1alpha1Repository {
	return o.Payload
}

func (o *RepositoryServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1Repository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepositoryServiceGetDefault creates a RepositoryServiceGetDefault with default headers values
func NewRepositoryServiceGetDefault(code int) *RepositoryServiceGetDefault {
	return &RepositoryServiceGetDefault{
		_statusCode: code,
	}
}

/* RepositoryServiceGetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RepositoryServiceGetDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the repository service get default response
func (o *RepositoryServiceGetDefault) Code() int {
	return o._statusCode
}

func (o *RepositoryServiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/repositories/{repo}][%d] RepositoryService_Get default  %+v", o._statusCode, o.Payload)
}
func (o *RepositoryServiceGetDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *RepositoryServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
