package request

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Interface interface {
	Registration(http.ResponseWriter, *http.Request, httprouter.Params)
	SignIn(http.ResponseWriter, *http.Request, httprouter.Params)
}

type request struct {
}

func NewRequest() Interface {
	return &request{}
}

func (r *request) Registration(http.ResponseWriter, *http.Request, httprouter.Params) {

}

func (r *request) SignIn(http.ResponseWriter, *http.Request, httprouter.Params) {

}
