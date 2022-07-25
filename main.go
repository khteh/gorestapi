package main
import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"net/http"
	"html/template"
	"github.com/khteh/fibonacci"
	"github.com/khteh/greetings"
)
var templatePath = "templates/"
var templates = template.Must(template.ParseGlob(templatePath + "*.html"))
var validPath = regexp.MustCompile("^/(fibonacci|home)/([a-zA-Z0-9]+)$")
type Index struct {
	Title string
	Greeting string
}
type Fibonacci struct {
	Title string
	FibInput   uint32
	FibResult uint64
	Error string
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	greeting, _ := greetings.Greeting("")
	result := Index {
		Title: "GoLang RESTful API",
		Greeting: greeting,
	}
	e := templates.ExecuteTemplate(w, "index.html", result)
	if e != nil {
		//log.Fatal(e)
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}
}
func fibonacciHandler(w http.ResponseWriter, r *http.Request) {
	result := Fibonacci {
		Title: "Fibonacci",
	}
	if r.Method == http.MethodPost {
		//fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
		n, err := strconv.ParseUint(r.FormValue("n"), 10, 32)
		if err != nil {
			if numError, ok := err.(*strconv.NumError); ok {
				if numError.Err == strconv.ErrRange {
					result.Error = fmt.Sprintf("Invalid input %s out of range!", r.FormValue("n"))
				}
			} else {
				//log.Fatal(err) This stops the http server!
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		} else {
			result.FibInput = uint32(n)
			result.FibResult = fibonacci.Fibonacci(uint32(n))
		}
	}
	e := templates.ExecuteTemplate(w, "fibonacci.html", result)
	if e != nil {
		//log.Fatal(e)
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			fmt.Printf("Invalid path! %s\n", r.URL.Path)
			//return "", errors.New("Invalid Page Title! " + r.URL.Path)
		} else {
			//for i := 0; i < len(m); i++ {
			//	fmt.Printf("%d: %s\n", i, m[i])
			//}
			fn(w, r, m[2]) // 0: r.URL.Path, 1: First group - (edit|save|view), 2: Second group - ([a-zA-Z0-9]+) = Title
		}
	}
}
// healthz is a liveness probe.
func healthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func main() {
	defer func() {
		fmt.Println("Bye!")
	}()
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/fibonacci", fibonacciHandler)
	//http.HandleFunc("/edit/", makeHandler(editHandler))
	//http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil))
}