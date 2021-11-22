package main

import "time"

func main() {
	NewServer("url", SetF1(1), SetF2("f2"), SetF3(time.Now()))
}

type Server struct {
	url string
	f1  int
	f2  string
	f3  time.Time
}

type OptionFunc func(*Server)

func NewServer(url string, options ...OptionFunc) *Server {
	s := &Server{
		url: url,
	}
	for _, option := range options {
		option(s)
	}
	return s
}

func SetF1(i int) OptionFunc {
	return func(s *Server) {
		s.f1 = i
	}
}

func SetF2(str string) OptionFunc {
	return func(s *Server) {
		s.f2 = str
	}
}

func SetF3(t time.Time) OptionFunc {
	return func(s *Server) {
		s.f3 = t
	}
}
