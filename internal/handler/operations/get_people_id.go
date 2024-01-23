// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetPeopleIDHandlerFunc turns a function with the right signature into a get people ID handler
type GetPeopleIDHandlerFunc func(GetPeopleIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPeopleIDHandlerFunc) Handle(params GetPeopleIDParams) middleware.Responder {
	return fn(params)
}

// GetPeopleIDHandler interface for that can handle valid get people ID params
type GetPeopleIDHandler interface {
	Handle(GetPeopleIDParams) middleware.Responder
}

// NewGetPeopleID creates a new http.Handler for the get people ID operation
func NewGetPeopleID(ctx *middleware.Context, handler GetPeopleIDHandler) *GetPeopleID {
	return &GetPeopleID{Context: ctx, Handler: handler}
}

/*
GetPeopleID swagger:route GET /people/{id} getPeopleId

Get person by ID
*/
type GetPeopleID struct {
	Context *middleware.Context
	Handler GetPeopleIDHandler
}

func (o *GetPeopleID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPeopleIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
