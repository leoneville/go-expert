package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 40}
	templ := template.New("CursoTemplate")
	templ, _ = templ.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	if err := templ.Execute(os.Stdout, curso); err != nil {
		panic(err)
	}
}
