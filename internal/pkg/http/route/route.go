package route

import "github.com/quicklygabbing/http/pkg/http/route"

type Interface interface {
	GetRoutes() []route.Routes
}
