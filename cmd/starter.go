package cmd

import (
	"github.com/pkg/errors"
	"github.com/quicklygabbing/users/pkg/http"
)

func Starter(address *string) (err error) {
	err = http.NewServer(address)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
