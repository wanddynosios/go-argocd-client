// Code generated by go-swagger; DO NOT EDIT.

package certificate_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/wanddynosios/go-argocd-client/models"
)

// CertificateServiceDeleteCertificateReader is a Reader for the CertificateServiceDeleteCertificate structure.
type CertificateServiceDeleteCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CertificateServiceDeleteCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCertificateServiceDeleteCertificateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCertificateServiceDeleteCertificateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCertificateServiceDeleteCertificateOK creates a CertificateServiceDeleteCertificateOK with default headers values
func NewCertificateServiceDeleteCertificateOK() *CertificateServiceDeleteCertificateOK {
	return &CertificateServiceDeleteCertificateOK{}
}

/* CertificateServiceDeleteCertificateOK describes a response with status code 200, with default header values.

A successful response.
*/
type CertificateServiceDeleteCertificateOK struct {
	Payload *models.V1alpha1RepositoryCertificateList
}

func (o *CertificateServiceDeleteCertificateOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/certificates][%d] certificateServiceDeleteCertificateOK  %+v", 200, o.Payload)
}
func (o *CertificateServiceDeleteCertificateOK) GetPayload() *models.V1alpha1RepositoryCertificateList {
	return o.Payload
}

func (o *CertificateServiceDeleteCertificateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1RepositoryCertificateList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCertificateServiceDeleteCertificateDefault creates a CertificateServiceDeleteCertificateDefault with default headers values
func NewCertificateServiceDeleteCertificateDefault(code int) *CertificateServiceDeleteCertificateDefault {
	return &CertificateServiceDeleteCertificateDefault{
		_statusCode: code,
	}
}

/* CertificateServiceDeleteCertificateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CertificateServiceDeleteCertificateDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the certificate service delete certificate default response
func (o *CertificateServiceDeleteCertificateDefault) Code() int {
	return o._statusCode
}

func (o *CertificateServiceDeleteCertificateDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/certificates][%d] CertificateService_DeleteCertificate default  %+v", o._statusCode, o.Payload)
}
func (o *CertificateServiceDeleteCertificateDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *CertificateServiceDeleteCertificateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
