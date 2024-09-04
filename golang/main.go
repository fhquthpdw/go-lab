package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

func main() {
	// 解析模板
	tmpl, err := template.New("app-configmap.txt.tpl").Funcs(template.FuncMap{
		"hasPrefix": strings.HasPrefix,
	}).ParseFiles("gin/app-configmap.txt.tpl")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	// 准备数据
	data := struct {
		Name string
	}{
		Name: "World",
	}

	// 渲染模板
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal("Execute Files error: ", err)
	}
}
