// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Kartochnik010/test-effectivemobile-jan/internal/models"
)

// PutPeopleIDOKCode is the HTTP code returned for type PutPeopleIDOK
const PutPeopleIDOKCode int = 200

/*
PutPeopleIDOK Person updated successfully

swagger:response putPeopleIdOK
*/
type PutPeopleIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Person `json:"body,omitempty"`
}

// NewPutPeopleIDOK creates PutPeopleIDOK with default headers values
func NewPutPeopleIDOK() *PutPeopleIDOK {

	return &PutPeopleIDOK{}
}

// WithPayload adds the payload to the put people Id o k response
func (o *PutPeopleIDOK) WithPayload(payload *models.Person) *PutPeopleIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put people Id o k response
func (o *PutPeopleIDOK) SetPayload(payload *models.Person) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutPeopleIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
