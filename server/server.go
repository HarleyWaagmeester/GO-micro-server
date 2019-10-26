package main
// Simple webserver running system utilities. And more.
import (
        "fmt"
        "html"
        "log"
        "net/http"
	"github.com/HarleyWaagmeester/execPing"
)

func main() {


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
        })

        http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
                fmt.Fprintf(w, "Hi")
        })
        http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request){
		flusher, ok := w.(http.Flusher)
		if !ok {
		 	http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		 	return
		}
		w.Header().Set("Content-Type", "text/html")
		flusher.Flush()
		execPing.Ping(w);
        })

        log.Fatal(http.ListenAndServe(":8081", nil))

} //main
