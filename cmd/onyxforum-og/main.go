package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/igorsaux/onyxforum-og/pkg/browser"
	"github.com/igorsaux/onyxforum-og/pkg/page"
)

var (
	host = flag.String("host", "127.0.0.1", "serving host")
	port = flag.Int("port", 8080, "serving port")
)

func init() {
	flag.Parse()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		target := query.Get("target")

		if target == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("`target` in query is missing."))

			return
		}

		post := page.ParsePost(target)
		html := page.RenderPost(post)
		image := browser.HtmlToImage(html, "#og")

		w.Header().Add("Content-Type", "image/png")
		w.WriteHeader(http.StatusOK)
		w.Write(image)

		return
	})

	fmt.Printf("Runs on http://%v:%v\n", *host, *port)
	http.ListenAndServe(fmt.Sprintf("%v:%v", *host, *port), nil)
}
