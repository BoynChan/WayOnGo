package controller

import (
	"html/template"
	"io/ioutil"
	"os"
)

// Author:Boyn
// Date:2020/2/23

// 这个是全局变量中templates这个map的初始化函数
// 它会扫描templates包下的模板文件并将解析好的模板文件放在map中
func populateTemplates() map[string]*template.Template {
	const basePath = ".\\src\\web\\go-mega\\templates"
	result := make(map[string]*template.Template)

	layout := template.Must(template.ParseFiles(basePath + string(os.PathSeparator) + "_base.html"))
	dir, err := os.Open(basePath + string(os.PathSeparator) + "content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + string(os.PathSeparator) + "content" + string(os.PathSeparator) + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}
