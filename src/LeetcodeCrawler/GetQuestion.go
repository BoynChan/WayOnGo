package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Author:Boyn
// Date:2020/3/13

func main() {
	//csrftoken=JUm0RtjN8H6E77O7E4Rs3FdKRRl5BBOy39qxUwMt04HjBrwnsyviBWdDiMWP0ZId
	r := strings.NewReader("{'operationName':'questionData','variables':{'titleSlug':'majority-element'},'query':'query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    boundTopicId\n    title\n    titleSlug\n    content\n    translatedTitle\n    translatedContent\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    isLiked\n    similarQuestions\n    contributors {\n      username\n      profileUrl\n      avatarUrl\n      __typename\n    }\n    langToValidPlayground\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    companyTagStats\n    codeSnippets {\n      lang\n      langSlug\n      code\n      __typename\n    }\n    stats\n    hints\n    solution {\n      id\n      canSeeDetail\n      __typename\n    }\n    status\n    sampleTestCase\n    metaData\n    judgerAvailable\n    judgeType\n    mysqlSchemas\n    enableRunCode\n    envInfo\n    book {\n      id\n      bookName\n      pressName\n      source\n      shortDescription\n      fullDescription\n      bookImgUrl\n      pressImgUrl\n      productUrl\n      __typename\n    }\n    isSubscribed\n    isDailyQuestion\n    dailyRecordStatus\n    editorType\n    ugcQuestionId\n    __typename\n  }\n}\n'}")
	request, _ := http.NewRequest("POST", "https://leetcode-cn.com/graphql/", r)
	request.AddCookie(&http.Cookie{
		Name:  "csrftoken",
		Value: "JUm0RtjN8H6E77O7E4Rs3FdKRRl5BBOy39qxUwMt04HjBrwnsyviBWdDiMWP0ZId",
	})
	request.Header["x-csrftoken"] = []string{"JUm0RtjN8H6E77O7E4Rs3FdKRRl5BBOy39qxUwMt04HjBrwnsyviBWdDiMWP0ZId"}
	// sec-fetch-dest: empty
	//sec-fetch-mode: cors
	//sec-fetch-site: same-origin
	request.Header["sec-fetch-dest"] = []string{"empty"}
	request.Header["sec-fetch-mode"] = []string{"cors"}
	request.Header["User-Agent"] = []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"}
	request.Header["sec-fetch-site"] = []string{"same-origin"}

	c := http.DefaultClient
	res, _ := c.Do(request)
	fmt.Println(res.StatusCode)
	content, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(content))
}
