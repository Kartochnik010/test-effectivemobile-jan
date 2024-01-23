// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostPeopleHandlerFunc turns a function with the right signature into a post people handler
type PostPeopleHandlerFunc func(PostPeopleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostPeopleHandlerFunc) Handle(params PostPeopleParams) middleware.Responder {
	return fn(params)
}

// PostPeopleHandler interface for that can handle valid post people params
type PostPeopleHandler interface {
	Handle(PostPeopleParams) middleware.Responder
}

// NewPostPeople creates a new http.Handler for the post people operation
func NewPostPeople(ctx *middleware.Context, handler PostPeopleHandler) *PostPeople {
	return &PostPeople{Context: ctx, Handler: handler}
}

/*
PostPeople swagger:route POST /people postPeople

Add a new person
*/
type PostPeople struct {
	Context *middleware.Context
	Handler PostPeopleHandler
}

func (o *PostPeople) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostPeopleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}