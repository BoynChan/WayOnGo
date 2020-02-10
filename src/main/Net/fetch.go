package Net

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func FetchAll(urls []string) {
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go fetchByRoutine(url, ch)
	}
	for range urls {
		//获取管道中输出
		fmt.Println(<-ch)
	}
	secs := time.Since(start).Milliseconds()
	fmt.Printf("Fetch Finished. Used Time:%dms", secs)
	close(ch)
}

func fetchByRoutine(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Fetching: %s", err)
		return
	}
	//暂时先将网页进行丢弃
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("Reading: %s", err)
		return
	}
	secs := time.Since(start).Milliseconds()
	ch <- fmt.Sprintf("Read(%s): %s | Used Time: %dms | Readed Bytes: %d", url, resp.Status, secs, nbytes)

}

func Fetch(args []string) {
	for _, url := range args {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
			os.Exit(1)
		}
		var prefix string
		if strings.HasPrefix(url, "https") {
			prefix = "HTTPS"
		} else {
			prefix = "HTTP"
		}
		_, _ = fmt.Fprintf(os.Stdout, "%s %s\n", prefix, resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
