package response

import (
	"github.com/open-runtimes/types-for-go/v4/openruntimes"
)

type Router struct {
	// path -> method -> handler
	routes map[string]map[string]func(c *openruntimes.Context) openruntimes.Response
}

func NewRouter() *Router {
	return &Router{routes: make(map[string]map[string]func(c *openruntimes.Context) openruntimes.Response)}
}

func (r *Router) AddRoute(path string, method string, handler func(Context *openruntimes.Context) openruntimes.Response) {
	r.routes[path][method] = handler
}

func (r *Router) Handle(Context *openruntimes.Context) openruntimes.Response {
	method := Context.Req.Method
	path := Context.Req.Path
	handler, ok := r.routes[path][method]
	if !ok {
		return NewStatusErrorResponse(Context, 404)
	}
	return handler(Context)
}
