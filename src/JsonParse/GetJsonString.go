package JsonParse

import (
	"io/ioutil"
	"net/http"
)

// Author:Boyn
// Date:2020/3/15

func GetString() string {
	response, _ := http.Get("http://news-at.zhihu.com/api/4/news/latest")
	jsonBytes, _ := ioutil.ReadAll(response.Body)
	jsonString := string(jsonBytes)
	return jsonString
}
