package main

import "time"

type Server struct {
	url string
	f1  int
	f2  string
	f3  time.Time
}

type Option func(*Server)

func NewServer(url string, options ...Option) *Server {
	s := &Server{}
	for _, option := range options {
		option(s)
	}
	return s
}

func SetF1(i int) Option {
	return func(s *Server) {
		s.f1 = 1
	}
}

func SetF2(str string) Option {
	return func(s *Server) {
		s.f2 = str
	}
}

func SetF3(t time.Time) Option {
	return func(s *Server) {
		s.f3 = time.Time{}
	}
}
