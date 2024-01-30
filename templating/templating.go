package templating

import (
	"embed"
	"go-with-tests/files"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	templ *template.Template
}

func NewPostRendered() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")

	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p files.Post) error {
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}

	return nil

}
