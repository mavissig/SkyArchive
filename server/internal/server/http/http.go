package http

import (
	"log"
	"net/http"
)

type UseCase interface {
	Mock()
	Auth()
}

type Http struct {
	uc UseCase
}

func (s *Http) Run() {
	http.HandleFunc("/auth", s.Auth)
	http.HandleFunc("/user", s.GetUser)
	http.HandleFunc("/upload", s.UploadFile)
	http.HandleFunc("download", s.DownloadFile)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("error run server: ", err)
	}
}

func (s *Http) Auth(w http.ResponseWriter, r *http.Request) {
	s.uc.Auth()
}

func (s *Http) GetUser(w http.ResponseWriter, r *http.Request) {
	s.uc.Mock()
}

func (s *Http) UploadFile(w http.ResponseWriter, r *http.Request) {
	s.uc.Mock()
}

func (s *Http) DownloadFile(w http.ResponseWriter, r *http.Request) {
	s.uc.Mock()
}

func New(uc UseCase) *Http {
	return &Http{
		uc: uc,
	}
}
