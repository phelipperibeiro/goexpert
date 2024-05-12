package main

import (
	"os"
	"text/template"
)

// documentation
// https://pkg.go.dev/html/template

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 40}
	tmp := template.New("CursoTemplate")
	tmp, err := tmp.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}\n")
	if err != nil {
		panic(err)
	}
	err = tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}

// go run 4-important-packages/6-templates/1/main.go
