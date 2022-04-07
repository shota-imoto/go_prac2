package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("substitute.html"))
	t.Execute(w, "Hello World")
}

func random(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("random.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func iterate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("iterate.html")
	//	monsters := []string{"pikachu", "ev", "camex", "mu2"}
	monsters := []string{}
	t.Execute(w, monsters)
}

func include(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("include1.html", "include2.html")
	t.Execute(w, "Hello World")
}

// カスタム変数
func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func custom(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("custom.html").Funcs(funcMap)
	t, _ = t.ParseFiles("custom.html")
	t.Execute(w, time.Now())
}

func context(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("context.html")
	content := `I asked: <i>"What's up?</i>`
	t.Execute(w, content)
}

func xss_server(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("xss_server.html")
	t.Execute(w, r.FormValue("comment"))
}

func xss_form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("xss_form.html")
	t.Execute(w, nil)
}

func layout(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("layout.html", "layout_red.html")
	} else {
		t, _ = template.ParseFiles("layout.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/random", random)
	http.HandleFunc("/iterate", iterate)
	http.HandleFunc("/include", include)
	http.HandleFunc("/custom", custom)
	http.HandleFunc("/context", context)
	http.HandleFunc("/xss_form", xss_form)
	http.HandleFunc("/xss_server", xss_server)
	http.HandleFunc("/layout", layout)

	server.ListenAndServe()
}
