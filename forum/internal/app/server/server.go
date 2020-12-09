package server

import (
	"log"
	"net/http"

	"github.com/61lf0yl3/div-01/forum/internal/app/handlers"
)

// Server ...
type Server struct {
	connect *handlers.Connect
}

// Start ...
func Start() error {

	s := Server{}

	if err := s.connect.Database.InitDB("forum.db"); err != nil {
		log.Println(err)
		return err
	}

	s.configureRouter()

	log.Println("server started succefully")
	return http.ListenAndServe(":8001", nil)
}

func (s *Server) configureRouter() {
	http.HandleFunc("/", s.connect.MainHandler)
	http.HandleFunc("/signup", s.connect.SignupHandler)
	http.HandleFunc("/usercreate", s.connect.UserCreateHandler)
}
