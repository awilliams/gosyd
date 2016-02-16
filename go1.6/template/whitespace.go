package main

import "text/template"
import "os"

func main() {
	const tmpl = `{{ range . }}
{{ . }}
{{ end }}`
	t := template.Must(template.New("").Parse(tmpl))
	t.Execute(os.Stdout, []string{"mars", "mercury", "pluto", "neptune"})
}
