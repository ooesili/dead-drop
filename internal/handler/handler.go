package handler

import (
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"
	"github.com/ooesili/dead-drop/internal"
)

func NewRouter(dropper internal.Dropper, renderer internal.Renderer) *httprouter.Router {
	h := handler{
		dropper:  dropper,
		renderer: renderer,
	}

	r := httprouter.New()
	r.HandlerFunc("GET", "/", h.home)
	r.HandlerFunc("POST", "/drop", h.drop)
	r.HandlerFunc("GET", "/complete", h.complete)
	return r
}

type handler struct {
	dropper  internal.Dropper
	renderer internal.Renderer
}

func (h handler) home(resp http.ResponseWriter, req *http.Request) {
	h.renderer.Render("index", resp, req)
}

func (h handler) drop(resp http.ResponseWriter, req *http.Request) {
	file, fileHeader, err := req.FormFile("file")
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := path.Base(fileHeader.Filename)
	h.dropper.Drop(filename, file)

	http.Redirect(resp, req, "/complete", http.StatusSeeOther)
}

func (h handler) complete(resp http.ResponseWriter, req *http.Request) {
	h.renderer.Render("complete", resp, req)
}
