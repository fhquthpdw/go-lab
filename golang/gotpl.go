package main

import (
	"os"
	"text/template"
)

const t = `Look at this: {{"{{"}} $name {{"}}"}} this is a test 
{{range .}} {{.Name}} {{.Count }} 
{{.JoinString}}
{{end}}`

type Record struct {
	Name       string
	Count      int
	JoinString string
}

func main() {
	tmpl, err := template.New("test").Parse(t)
	if err != nil {
		panic(err)
	}
	var r []Record
	r = append(r, Record{
		Name:       "Suzanne",
		Count:      236,
		JoinString: "rs",
	})
	f, _ := os.Create("test.txt")
	err = tmpl.Execute(f, r)

	//err = tmpl.Execute(os.Stdout, r)

	if err != nil {
		panic(err)
	}
}
