package internal

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/julienschmidt/httprouter"
)

type DropHandler struct{}

func (d DropHandler) DropForm(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	view.Render(resp, "index", map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(req),
	})
}

func (d DropHandler) Drop(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	err := Drop(req)
	if err != nil {
		status := http.StatusInternalServerError

		if statusErr, ok := err.(interface {
			Status() int
		}); ok {
			status = statusErr.Status()
		}

		http.Error(resp, err.Error(), status)
	}

	http.Redirect(resp, req, "/complete", http.StatusSeeOther)
}

func (d DropHandler) Complete(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	view.Render(resp, "complete", nil)
}
