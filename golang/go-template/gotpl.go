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

type RecordFile struct {
	Name        string
	Project     string
	Environment string
}

func main() {
	// parse text
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

	// parse template file
	tmplFile, err := template.ParseFiles("tpl.yaml.tpl")
	var rFile []RecordFile
	rFile = append(rFile, RecordFile{
		Name:        "touch-point",
		Project:     "tpp",
		Environment: "dev",
	})
	fFile, _ := os.Create("testFile.txt")
	err = tmplFile.Execute(fFile, rFile)
	if err != nil {
		panic(err)
	}
}
