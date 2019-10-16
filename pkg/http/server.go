package http

import (
	"github.com/pkg/errors"
	"github.com/quicklygabbing/http/pkg/http"
	"github.com/quicklygabbing/users/pkg/http/route"
)

func NewServer(address *string) (err error) {
	server := http.NewServer(address)
	err = server.Start(route.NewRoute().GetRoutes())
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
