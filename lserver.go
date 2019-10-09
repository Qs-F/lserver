package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/Qs-F/rewrite"
	"github.com/sirupsen/logrus"
)

type Option struct {
	Port       *string // which port server use
	Directory  *string // which directory server publish
	RewriteOld *string // Rewrite old string
	RewriteNew *string // Rewrite new string
	IsPublic   *bool   // internallly public, or globally public
	IsNotCORS  *bool   // availability of CORS
}

func handleFlag() *Option {
	return &Option{
		Port:       flag.String("p", "8080", "Server exposing port (default is 8080)"),
		Directory:  flag.String("d", "./", "Server exposing directory (default is current directory)"),
		RewriteOld: flag.String("old", "", "Rewrite old"),
		RewriteNew: flag.String("new", "", "Rewrite new"),
		IsPublic:   flag.Bool("pub", false, "internal server or public server (default is internal)"),
		IsNotCORS:  flag.Bool("cors", false, "Use or not CORS (default is using, adding this flag means forbidden CORS)"),
	}
}

type HTTPHeader struct {
	Key   string
	Value string
}

type Server struct {
	Addr      string
	Directory string
	CORS      bool
	Rewrite   *rewrite.Rule
}

func (o *Option) newServer() *Server {
	addr := ""
	if *o.IsPublic {
		addr = "localhost"
	} else {
		addr = "0.0.0.0"
	}
	return &Server{
		Addr:      addr + ":" + *o.Port,
		Directory: *o.Directory,
		CORS:      !*o.IsNotCORS,
		Rewrite:   &rewrite.Rule{{Old: *o.RewriteOld, New: *o.RewriteNew}},
	}
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	})
}

func log(s string) {
	logrus.Infoln("Time: " + time.Now().Format(time.StampMilli) + "    Req: " + s)
}

func connLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log(r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

func (s *Server) getHandler() http.Handler {
	if s.CORS {
		return cors(s.Rewrite.Map(http.FileServer(http.Dir(s.Directory))))
	} else {
		return s.Rewrite.Map(http.FileServer(http.Dir(s.Directory)))
	}
}

func main() {
	o := handleFlag()
	flag.Parse()
	s := o.newServer()
	handler := connLog(s.getHandler())
	log("Starting Server at… [ " + s.Addr + " ]")
	http.Handle("/", handler)
	err := http.ListenAndServe(s.Addr, nil)
	if err != nil {
		logrus.Fatalln(err.Error())
	}
}
