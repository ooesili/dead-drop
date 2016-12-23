package view

import (
	"html/template"
	"io"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/ooesili/dead-drop/internal/bindata/templates"
)

//go:generate go-bindata -nometadata -o ../bindata/templates/bindata.go -pkg templates -prefix ../../templates ../../templates
func New() View {
	paths := []string{
		"index",
		"complete",
	}
	view := View{
		templates: map[string]*template.Template{},
	}

	for _, path := range paths {
		baseTmpl := string(templates.MustAsset("layout.html.tmpl"))
		viewTmpl := string(templates.MustAsset(path + ".html.tmpl"))

		view.templates[path] = template.Must(template.New("layout").Parse(baseTmpl))
		template.Must(view.templates[path].New(path).Parse(viewTmpl))
	}

	return view
}

type View struct {
	templates map[string]*template.Template
}

func (v View) Render(name string, out io.Writer, req *http.Request) error {
	return v.templates[name].Execute(out, map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(req),
	})
}
