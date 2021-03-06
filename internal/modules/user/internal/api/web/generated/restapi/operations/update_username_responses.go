// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/Meat-Hook/point-bank/internal/modules/user/internal/api/web/generated/models"
)

// UpdateUsernameNoContentCode is the HTTP code returned for type UpdateUsernameNoContent
const UpdateUsernameNoContentCode int = 204

/*UpdateUsernameNoContent The server successfully processed the request and is not returning any content.

swagger:response updateUsernameNoContent
*/
type UpdateUsernameNoContent struct {
}

// NewUpdateUsernameNoContent creates UpdateUsernameNoContent with default headers values
func NewUpdateUsernameNoContent() *UpdateUsernameNoContent {

	return &UpdateUsernameNoContent{}
}

// WriteResponse to the client
func (o *UpdateUsernameNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *UpdateUsernameNoContent) UpdateUsernameResponder() {}

/*UpdateUsernameDefault Generic error response.

swagger:response updateUsernameDefault
*/
type UpdateUsernameDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateUsernameDefault creates UpdateUsernameDefault with default headers values
func NewUpdateUsernameDefault(code int) *UpdateUsernameDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateUsernameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update username default response
func (o *UpdateUsernameDefault) WithStatusCode(code int) *UpdateUsernameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update username default response
func (o *UpdateUsernameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update username default response
func (o *UpdateUsernameDefault) WithPayload(payload *models.Error) *UpdateUsernameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update username default response
func (o *UpdateUsernameDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUsernameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *UpdateUsernameDefault) UpdateUsernameResponder() {}

type UpdateUsernameNotImplementedResponder struct {
	middleware.Responder
}

func (*UpdateUsernameNotImplementedResponder) UpdateUsernameResponder() {}

func UpdateUsernameNotImplemented() UpdateUsernameResponder {
	return &UpdateUsernameNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.UpdateUsername has not yet been implemented",
		),
	}
}

type UpdateUsernameResponder interface {
	middleware.Responder
	UpdateUsernameResponder()
}
