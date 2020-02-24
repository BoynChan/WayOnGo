package controller

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"web/go-mega/vm"
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
	username, ok := val.(string)
	if !ok {
		return "", errors.New("can not get session user")
	}
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

// 检查长度,字段名,字段域和最大最小长度都由调用者指定
func checkLen(fieldName, fieldValue string, minLen, maxLen int) string {
	lenField := len(fieldValue)
	if lenField < minLen {
		return fmt.Sprintf("%s field is too short, less than %d", fieldName, minLen)
	}
	if lenField > maxLen {
		return fmt.Sprintf("%s field is too long, more than %d", fieldName, maxLen)
	}
	return ""
}

// 检查用户名长度
func checkUsername(username string) string {
	return checkLen("Username", username, 3, 20)
}

// 检查密码长度
func checkPassword(password string) string {
	return checkLen("Password", password, 6, 50)
}

// 正则匹配查看邮箱地址是否正确
func checkEmail(email string) string {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		return fmt.Sprintf("Email field not a valid email")
	}
	return ""
}

// 检查用户用户名与密码是否正确
func checkUserPassword(username, password string) string {
	if !vm.CheckLogin(username, password) {
		return fmt.Sprintf("Username and password is not correct.")
	}
	return ""
}

// 检查用户是否存在
func checkUserExist(username string) string {
	if vm.CheckUserExist(username) {
		return fmt.Sprintf("Username already exist, please choose another username")
	}
	return ""
}

// 检查登录时的参数
func checkLogin(username, password string) []string {
	var errs []string
	if errCheck := checkUsername(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkPassword(password); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkUserPassword(username, password); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	return errs
}

// 检查注册时的参数
func checkRegister(username, email, password string) []string {
	var errs []string
	if errCheck := checkUsername(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkPassword(password); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkEmail(email); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkUserExist(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	return errs
}

// 添加新用户
func addUser(username, password, email string) error {
	return vm.AddUser(username, password, email)
}

func setFlash(w http.ResponseWriter, r *http.Request, message string) {
	session, _ := store.Get(r, sessionName)
	session.AddFlash(message, flashName)
	session.Save(r, w)
}

func getFlash(w http.ResponseWriter, r *http.Request) string {
	session, _ := store.Get(r, sessionName)
	fm := session.Flashes(flashName)
	if fm == nil {
		return ""
	}
	session.Save(r, w)
	return fmt.Sprintf("%v", fm[0])
}

// 获取请求中的分页参数
func getPage(r *http.Request) int {
	url := r.URL
	query := url.Query()
	q := query.Get("page")
	if q == "" {
		return 1
	}
	page, err := strconv.Atoi(q)
	if err != nil {
		return 1
	}
	return page
}
