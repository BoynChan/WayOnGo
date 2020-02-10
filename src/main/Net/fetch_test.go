package Net

import "testing"

func TestFetch(t *testing.T) {
	Fetch([]string{"http://gopl.io"})
}

func TestFetchAllFunc(t *testing.T) {
	FetchAll([]string{"http://gopl.io", "http://baidu.com", "https://golang.org", "https://godoc.org"})
}
