package main

// Simple webserver running system utilities. And more.
import (
	"fmt"
	//	"github.com/HarleyWaagmeester/execPing"
	"execPing"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "echo...")
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		flusher.Flush()
		name, _ := execPing.Ping(w)
		println(name)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

} //main
