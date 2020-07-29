// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MakeTraefikDecisionReader is a Reader for the MakeTraefikDecision structure.
type MakeTraefikDecisionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MakeTraefikDecisionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMakeTraefikDecisionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewMakeTraefikDecisionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMakeTraefikDecisionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMakeTraefikDecisionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewMakeTraefikDecisionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMakeTraefikDecisionOK creates a MakeTraefikDecisionOK with default headers values
func NewMakeTraefikDecisionOK() *MakeTraefikDecisionOK {
	return &MakeTraefikDecisionOK{}
}

/*MakeTraefikDecisionOK handles this case with default header values.

An empty response
*/
type MakeTraefikDecisionOK struct {
}

func (o *MakeTraefikDecisionOK) Error() string {
	return fmt.Sprintf("[GET /decisions/traefik][%d] makeTraefikDecisionOK ", 200)
}

func (o *MakeTraefikDecisionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMakeTraefikDecisionUnauthorized creates a MakeTraefikDecisionUnauthorized with default headers values
func NewMakeTraefikDecisionUnauthorized() *MakeTraefikDecisionUnauthorized {
	return &MakeTraefikDecisionUnauthorized{}
}

/*MakeTraefikDecisionUnauthorized handles this case with default header values.

The standard error format
*/
type MakeTraefikDecisionUnauthorized struct {
	Payload *MakeTraefikDecisionUnauthorizedBody
}

func (o *MakeTraefikDecisionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /decisions/traefik][%d] makeTraefikDecisionUnauthorized  %+v", 401, o.Payload)
}

func (o *MakeTraefikDecisionUnauthorized) GetPayload() *MakeTraefikDecisionUnauthorizedBody {
	return o.Payload
}

func (o *MakeTraefikDecisionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(MakeTraefikDecisionUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMakeTraefikDecisionForbidden creates a MakeTraefikDecisionForbidden with default headers values
func NewMakeTraefikDecisionForbidden() *MakeTraefikDecisionForbidden {
	return &MakeTraefikDecisionForbidden{}
}

/*MakeTraefikDecisionForbidden handles this case with default header values.

The standard error format
*/
type MakeTraefikDecisionForbidden struct {
	Payload *MakeTraefikDecisionForbiddenBody
}

func (o *MakeTraefikDecisionForbidden) Error() string {
	return fmt.Sprintf("[GET /decisions/traefik][%d] makeTraefikDecisionForbidden  %+v", 403, o.Payload)
}

func (o *MakeTraefikDecisionForbidden) GetPayload() *MakeTraefikDecisionForbiddenBody {
	return o.Payload
}

func (o *MakeTraefikDecisionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(MakeTraefikDecisionForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMakeTraefikDecisionNotFound creates a MakeTraefikDecisionNotFound with default headers values
func NewMakeTraefikDecisionNotFound() *MakeTraefikDecisionNotFound {
	return &MakeTraefikDecisionNotFound{}
}

/*MakeTraefikDecisionNotFound handles this case with default header values.

The standard error format
*/
type MakeTraefikDecisionNotFound struct {
	Payload *MakeTraefikDecisionNotFoundBody
}

func (o *MakeTraefikDecisionNotFound) Error() string {
	return fmt.Sprintf("[GET /decisions/traefik][%d] makeTraefikDecisionNotFound  %+v", 404, o.Payload)
}

func (o *MakeTraefikDecisionNotFound) GetPayload() *MakeTraefikDecisionNotFoundBody {
	return o.Payload
}

func (o *MakeTraefikDecisionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(MakeTraefikDecisionNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMakeTraefikDecisionInternalServerError creates a MakeTraefikDecisionInternalServerError with default headers values
func NewMakeTraefikDecisionInternalServerError() *MakeTraefikDecisionInternalServerError {
	return &MakeTraefikDecisionInternalServerError{}
}

/*MakeTraefikDecisionInternalServerError handles this case with default header values.

The standard error format
*/
type MakeTraefikDecisionInternalServerError struct {
	Payload *MakeTraefikDecisionInternalServerErrorBody
}

func (o *MakeTraefikDecisionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /decisions/traefik][%d] makeTraefikDecisionInternalServerError  %+v", 500, o.Payload)
}

func (o *MakeTraefikDecisionInternalServerError) GetPayload() *MakeTraefikDecisionInternalServerErrorBody {
	return o.Payload
}

func (o *MakeTraefikDecisionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(MakeTraefikDecisionInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*MakeTraefikDecisionForbiddenBody make traefik decision forbidden body
swagger:model MakeTraefikDecisionForbiddenBody
*/
type MakeTraefikDecisionForbiddenBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// details
	Details []interface{} `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// request
	Request string `json:"request,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this make traefik decision forbidden body
func (o *MakeTraefikDecisionForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MakeTraefikDecisionForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MakeTraefikDecisionForbiddenBody) UnmarshalBinary(b []byte) error {
	var res MakeTraefikDecisionForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*MakeTraefikDecisionInternalServerErrorBody make traefik decision internal server error body
swagger:model MakeTraefikDecisionInternalServerErrorBody
*/
type MakeTraefikDecisionInternalServerErrorBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// details
	Details []interface{} `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// request
	Request string `json:"request,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this make traefik decision internal server error body
func (o *MakeTraefikDecisionInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MakeTraefikDecisionInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MakeTraefikDecisionInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res MakeTraefikDecisionInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*MakeTraefikDecisionNotFoundBody make traefik decision not found body
swagger:model MakeTraefikDecisionNotFoundBody
*/
type MakeTraefikDecisionNotFoundBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// details
	Details []interface{} `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// request
	Request string `json:"request,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this make traefik decision not found body
func (o *MakeTraefikDecisionNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MakeTraefikDecisionNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MakeTraefikDecisionNotFoundBody) UnmarshalBinary(b []byte) error {
	var res MakeTraefikDecisionNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*MakeTraefikDecisionUnauthorizedBody make traefik decision unauthorized body
swagger:model MakeTraefikDecisionUnauthorizedBody
*/
type MakeTraefikDecisionUnauthorizedBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// details
	Details []interface{} `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// request
	Request string `json:"request,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this make traefik decision unauthorized body
func (o *MakeTraefikDecisionUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MakeTraefikDecisionUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MakeTraefikDecisionUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res MakeTraefikDecisionUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}