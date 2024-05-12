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
	tmp := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}

// go run 4-important-packages/6-templates/2-must/main.go
