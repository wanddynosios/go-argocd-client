// Code generated by go-swagger; DO NOT EDIT.

package account_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/wanddynosios/go-argocd-client/models"
)

// AccountServiceGetAccountReader is a Reader for the AccountServiceGetAccount structure.
type AccountServiceGetAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountServiceGetAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountServiceGetAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAccountServiceGetAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccountServiceGetAccountOK creates a AccountServiceGetAccountOK with default headers values
func NewAccountServiceGetAccountOK() *AccountServiceGetAccountOK {
	return &AccountServiceGetAccountOK{}
}

/* AccountServiceGetAccountOK describes a response with status code 200, with default header values.

A successful response.
*/
type AccountServiceGetAccountOK struct {
	Payload *models.AccountAccount
}

func (o *AccountServiceGetAccountOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/account/{name}][%d] accountServiceGetAccountOK  %+v", 200, o.Payload)
}
func (o *AccountServiceGetAccountOK) GetPayload() *models.AccountAccount {
	return o.Payload
}

func (o *AccountServiceGetAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountServiceGetAccountDefault creates a AccountServiceGetAccountDefault with default headers values
func NewAccountServiceGetAccountDefault(code int) *AccountServiceGetAccountDefault {
	return &AccountServiceGetAccountDefault{
		_statusCode: code,
	}
}

/* AccountServiceGetAccountDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AccountServiceGetAccountDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the account service get account default response
func (o *AccountServiceGetAccountDefault) Code() int {
	return o._statusCode
}

func (o *AccountServiceGetAccountDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/account/{name}][%d] AccountService_GetAccount default  %+v", o._statusCode, o.Payload)
}
func (o *AccountServiceGetAccountDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AccountServiceGetAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
