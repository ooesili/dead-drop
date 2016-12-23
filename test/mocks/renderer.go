package mocks

import (
	"fmt"
	"io"
	"net/http"
)

type Renderer struct {
	RenderCall struct {
		Recieved struct {
			Template string
			Request  *http.Request
		}
		Renders string
		Returns struct {
			Err error
		}
	}
}

func (r *Renderer) Render(template string, out io.Writer, request *http.Request) error {
	call := &r.RenderCall
	call.Recieved.Template = template
	call.Recieved.Request = request

	fmt.Fprint(out, call.Renders)
	return call.Returns.Err
}
