// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/Meat-Hook/back-template/internal/modules/user/internal/api/web/generated/models"
)

// VerificationEmailNoContentCode is the HTTP code returned for type VerificationEmailNoContent
const VerificationEmailNoContentCode int = 204

/*VerificationEmailNoContent The server successfully processed the request and is not returning any content.

swagger:response verificationEmailNoContent
*/
type VerificationEmailNoContent struct {
}

// NewVerificationEmailNoContent creates VerificationEmailNoContent with default headers values
func NewVerificationEmailNoContent() *VerificationEmailNoContent {

	return &VerificationEmailNoContent{}
}

// WriteResponse to the client
func (o *VerificationEmailNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *VerificationEmailNoContent) VerificationEmailResponder() {}

/*VerificationEmailDefault Generic error response.

swagger:response verificationEmailDefault
*/
type VerificationEmailDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewVerificationEmailDefault creates VerificationEmailDefault with default headers values
func NewVerificationEmailDefault(code int) *VerificationEmailDefault {
	if code <= 0 {
		code = 500
	}

	return &VerificationEmailDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the verification email default response
func (o *VerificationEmailDefault) WithStatusCode(code int) *VerificationEmailDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the verification email default response
func (o *VerificationEmailDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the verification email default response
func (o *VerificationEmailDefault) WithPayload(payload *models.Error) *VerificationEmailDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the verification email default response
func (o *VerificationEmailDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VerificationEmailDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *VerificationEmailDefault) VerificationEmailResponder() {}

type VerificationEmailNotImplementedResponder struct {
	middleware.Responder
}

func (*VerificationEmailNotImplementedResponder) VerificationEmailResponder() {}

func VerificationEmailNotImplemented() VerificationEmailResponder {
	return &VerificationEmailNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.VerificationEmail has not yet been implemented",
		),
	}
}

type VerificationEmailResponder interface {
	middleware.Responder
	VerificationEmailResponder()
}
