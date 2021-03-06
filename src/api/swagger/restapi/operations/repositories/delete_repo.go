package repositories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteRepoHandlerFunc turns a function with the right signature into a delete repo handler
type DeleteRepoHandlerFunc func(DeleteRepoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteRepoHandlerFunc) Handle(params DeleteRepoParams) middleware.Responder {
	return fn(params)
}

// DeleteRepoHandler interface for that can handle valid delete repo params
type DeleteRepoHandler interface {
	Handle(DeleteRepoParams) middleware.Responder
}

// NewDeleteRepo creates a new http.Handler for the delete repo operation
func NewDeleteRepo(ctx *middleware.Context, handler DeleteRepoHandler) *DeleteRepo {
	return &DeleteRepo{Context: ctx, Handler: handler}
}

/*DeleteRepo swagger:route DELETE /v1/repos/{repoName} repositories deleteRepo

delete a repository enabled in the backend

*/
type DeleteRepo struct {
	Context *middleware.Context
	Handler DeleteRepoHandler
}

func (o *DeleteRepo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDeleteRepoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
