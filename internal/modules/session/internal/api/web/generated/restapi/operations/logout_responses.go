// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/Meat-Hook/point-bank/internal/modules/session/internal/api/web/generated/models"
)

// LogoutNoContentCode is the HTTP code returned for type LogoutNoContent
const LogoutNoContentCode int = 204

/*LogoutNoContent The server successfully processed the request and is not returning any content.

swagger:response logoutNoContent
*/
type LogoutNoContent struct {
}

// NewLogoutNoContent creates LogoutNoContent with default headers values
func NewLogoutNoContent() *LogoutNoContent {

	return &LogoutNoContent{}
}

// WriteResponse to the client
func (o *LogoutNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *LogoutNoContent) LogoutResponder() {}

/*LogoutDefault Generic error response.

swagger:response logoutDefault
*/
type LogoutDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewLogoutDefault creates LogoutDefault with default headers values
func NewLogoutDefault(code int) *LogoutDefault {
	if code <= 0 {
		code = 500
	}

	return &LogoutDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the logout default response
func (o *LogoutDefault) WithStatusCode(code int) *LogoutDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the logout default response
func (o *LogoutDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the logout default response
func (o *LogoutDefault) WithPayload(payload *models.Error) *LogoutDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the logout default response
func (o *LogoutDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LogoutDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *LogoutDefault) LogoutResponder() {}

type LogoutNotImplementedResponder struct {
	middleware.Responder
}

func (*LogoutNotImplementedResponder) LogoutResponder() {}

func LogoutNotImplemented() LogoutResponder {
	return &LogoutNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.Logout has not yet been implemented",
		),
	}
}

type LogoutResponder interface {
	middleware.Responder
	LogoutResponder()
}
