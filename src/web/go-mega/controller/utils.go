package controller

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
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

// 封装Session的操作
func getSessionUser(r *http.Request) (string, error) {
	var username string
	session, err := store.Get(r, sessionName)
	if err != nil {
		return "", err
	}
	val := session.Values["user"]
	fmt.Println("val:", val)
	username, ok := val.(string)
	if !ok {
		return "", errors.New("can not get session user")
	}

	fmt.Println("username:", username)
	return username, nil
}

func setSessionUser(w http.ResponseWriter, r *http.Request, username string) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Values["user"] = username
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

func clearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}
