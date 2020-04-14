package views

import (
	"html/template"
	"path/filepath"
)

func NewView(layout string, files ...string) *View {
	files = append(
		files,
		"views/layouts/bootstrap.gohtml",
		"views/layouts/alert.gohtml",
		"views/layouts/navbar.gohtml",
		"views/layouts/footer.gohtml",
		)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout: layout,
	}
}

type View struct {
	Template *template.Template
	Layout string
}

func layoutFiles() []string {
	files, err := filepath.Glob("views/layouts/*.gohtml")
	if err != nil {
		panic(err)
	}
}