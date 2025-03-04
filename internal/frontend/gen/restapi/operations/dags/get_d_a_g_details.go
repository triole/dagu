// Code generated by go-swagger; DO NOT EDIT.

package dags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetDAGDetailsHandlerFunc turns a function with the right signature into a get d a g details handler
type GetDAGDetailsHandlerFunc func(GetDAGDetailsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDAGDetailsHandlerFunc) Handle(params GetDAGDetailsParams) middleware.Responder {
	return fn(params)
}

// GetDAGDetailsHandler interface for that can handle valid get d a g details params
type GetDAGDetailsHandler interface {
	Handle(GetDAGDetailsParams) middleware.Responder
}

// NewGetDAGDetails creates a new http.Handler for the get d a g details operation
func NewGetDAGDetails(ctx *middleware.Context, handler GetDAGDetailsHandler) *GetDAGDetails {
	return &GetDAGDetails{Context: ctx, Handler: handler}
}

/*
	GetDAGDetails swagger:route GET /dags/{dagId} dags getDAGDetails

# Get DAG details

Returns details of a DAG, including files, logs, etc.
*/
type GetDAGDetails struct {
	Context *middleware.Context
	Handler GetDAGDetailsHandler
}

func (o *GetDAGDetails) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetDAGDetailsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
