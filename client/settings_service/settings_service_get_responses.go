// Code generated by go-swagger; DO NOT EDIT.

package settings_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/wanddynosios/go-argocd-client/models"
)

// SettingsServiceGetReader is a Reader for the SettingsServiceGet structure.
type SettingsServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SettingsServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSettingsServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSettingsServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSettingsServiceGetOK creates a SettingsServiceGetOK with default headers values
func NewSettingsServiceGetOK() *SettingsServiceGetOK {
	return &SettingsServiceGetOK{}
}

/* SettingsServiceGetOK describes a response with status code 200, with default header values.

A successful response.
*/
type SettingsServiceGetOK struct {
	Payload *models.ClusterSettings
}

func (o *SettingsServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/settings][%d] settingsServiceGetOK  %+v", 200, o.Payload)
}
func (o *SettingsServiceGetOK) GetPayload() *models.ClusterSettings {
	return o.Payload
}

func (o *SettingsServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSettingsServiceGetDefault creates a SettingsServiceGetDefault with default headers values
func NewSettingsServiceGetDefault(code int) *SettingsServiceGetDefault {
	return &SettingsServiceGetDefault{
		_statusCode: code,
	}
}

/* SettingsServiceGetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SettingsServiceGetDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the settings service get default response
func (o *SettingsServiceGetDefault) Code() int {
	return o._statusCode
}

func (o *SettingsServiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/settings][%d] SettingsService_Get default  %+v", o._statusCode, o.Payload)
}
func (o *SettingsServiceGetDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *SettingsServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
