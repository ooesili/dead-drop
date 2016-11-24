package internal

import (
	"html/template"
	"io"

	"github.com/ooesili/dead-drop/internal/bindata/templates"
)

//go:generate go-bindata -nometadata -o bindata/templates/bindata.go -pkg templates -prefix ../templates ../templates
func init() {
	paths := []string{
		"index",
		"complete",
	}
	view = templateEngine{}

	for _, path := range paths {
		baseTmpl := string(templates.MustAsset("layout.html.tmpl"))
		viewTmpl := string(templates.MustAsset(path + ".html.tmpl"))

		view[path] = template.Must(template.New("layout").Parse(baseTmpl))
		template.Must(view[path].New(path).Parse(viewTmpl))
	}
}

type templateEngine map[string]*template.Template

func (e templateEngine) Render(out io.Writer, name string, data interface{}) error {
	return e[name].Execute(out, data)
}

var view templateEngine
