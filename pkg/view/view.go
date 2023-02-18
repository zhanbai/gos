package view

import (
	"gos/pkg/logger"
	"io"
	"strings"
	"text/template"
)

// Render 渲染视图
func Render(w io.Writer, name string, data interface{}) {
	dir := "resources/views/"
	name = strings.Replace(name, ".", "/", -1)
	tmpl, err := template.ParseFiles(dir + name + ".tmpl")
	logger.LogError(err)

	err = tmpl.Execute(w, nil)
	logger.LogError(err)
}
