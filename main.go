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
	tmp := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := tmp.Execute(os.Stdout, []Curso{
		{"Go", 40},
		{"Java", 45},
		{"Pyton", 50},
	})
	if err != nil {
		panic(err)
	}
}
