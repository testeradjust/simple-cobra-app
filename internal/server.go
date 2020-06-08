package internal

import "fmt"

type Server struct {
	IsActive bool
	Version int
	Name string
}

func NewServer(version int, name string) *Server {
	server := &Server{
		IsActive: false,
		Version: version,
		Name: name,
	}
	return server
}

func (s *Server) IncrementVersion() {
	s.Version++
}

func (s *Server) SetActive() error {
	if s.Version <= 1 {
		return fmt.Errorf("server version invalid, cannot activate")
	}
	s.IsActive = true
	return nil
}

