package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	const master = `Names:{{"\n"}}{{block "list" .}}{{range .}}{{println "-" .}}{{end}}{{end}}`
	const overlay = `{{define "list"}}{{range .}}{{println "+" .}}{{end}}{{end}}`

	var guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}

	masterTmpl, _ := template.New("master").Parse(master)
	// Clone master template, add new template that redefines "list" // HL
	overlayTmpl, _ := template.Must(masterTmpl.Clone()).Parse(overlay)

	masterTmpl.Execute(os.Stdout, guardians)
	fmt.Println("-----")
	overlayTmpl.Execute(os.Stdout, guardians)
}
