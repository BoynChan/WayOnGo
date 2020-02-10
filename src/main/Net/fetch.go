package Net

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func fetch(args []string) {
	for _, url := range args[:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
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
			_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
