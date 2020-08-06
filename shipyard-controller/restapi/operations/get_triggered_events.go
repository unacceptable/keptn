// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetTriggeredEventsHandlerFunc turns a function with the right signature into a get triggered events handler
type GetTriggeredEventsHandlerFunc func(GetTriggeredEventsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTriggeredEventsHandlerFunc) Handle(params GetTriggeredEventsParams) middleware.Responder {
	return fn(params)
}

// GetTriggeredEventsHandler interface for that can handle valid get triggered events params
type GetTriggeredEventsHandler interface {
	Handle(GetTriggeredEventsParams) middleware.Responder
}

// NewGetTriggeredEvents creates a new http.Handler for the get triggered events operation
func NewGetTriggeredEvents(ctx *middleware.Context, handler GetTriggeredEventsHandler) *GetTriggeredEvents {
	return &GetTriggeredEvents{Context: ctx, Handler: handler}
}

/*GetTriggeredEvents swagger:route GET /event/triggered/{eventType} getTriggeredEvents

Get list of triggered events

*/
type GetTriggeredEvents struct {
	Context *middleware.Context
	Handler GetTriggeredEventsHandler
}

func (o *GetTriggeredEvents) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTriggeredEventsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
