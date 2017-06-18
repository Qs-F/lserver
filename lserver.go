package main

import (
	"flag"
	"fmt"
	"net/http"
	"text/template"
)

type View struct {
	Port string
}

var Gport string

var Body string = `
var host = localhost:{{.Port}}
`

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	})
}

func NewServer(dir, addr, port string) {
	Gport = port
	http.Handle("/", cors(http.FileServer(http.Dir(dir))))
	http.HandleFunc("/go/portinfo.js", viewJS)
	err := http.ListenAndServe(addr+":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func viewJS(w http.ResponseWriter, r *http.Request) {
	js := View{Gport}
	tmpl, err := template.New("new").Parse(Body)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, js)
	if err != nil {
		panic(err)
	}
}

func main() {
	var p = flag.String("p", "8080", "Set port")
	var d = flag.String("d", "./", "Set directory")
	var pub = flag.Bool("pub", false, "Public server(default is false)")
	addr := ""
	flag.Parse()
	if !*pub {
		addr = "0.0.0.0"
		fmt.Println("Start server @ 0.0.0.0:" + *p + ":" + *d)
	} else {
		addr = ""
		fmt.Println("Start server @ localhost:" + *p + ":" + *d)
	}
	fmt.Println("To stop server, pls ctrl-c")
	NewServer(*d, addr, *p)
}
