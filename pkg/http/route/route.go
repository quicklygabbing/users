package route

import (
	httpRoute "github.com/quicklygabbing/http/pkg/http/route"
	"github.com/quicklygabbing/users/internal/pkg/http/request"
	internalRoute "github.com/quicklygabbing/users/internal/pkg/http/route"
)

type route struct {
	request request.Interface
}

func NewRoute() internalRoute.Interface {
	return &route{request: request.NewRequest()}
}

func (r *route) GetRoutes() []httpRoute.Routes {
	return []httpRoute.Routes{
		r.registration(),
		r.signIn(),
	}
}

func (r *route) registration() httpRoute.Routes {
	return httpRoute.Routes{
		Route:   `/registration`,
		Methods: []string{httpRoute.POST, httpRoute.GET},
		Handle:  r.request.Registration,
	}
}

func (r *route) signIn() httpRoute.Routes {
	return httpRoute.Routes{
		Route:   `/sign-in`,
		Methods: []string{httpRoute.POST, httpRoute.GET},
		Handle:  r.request.SignIn,
	}
}
