package internal

type Server struct {
	Version int
	Name string
}

func NewServer(version int, name string) *Server {
	server := &Server{
		Version: version,
		Name: name,
	}
	return server
}

func (s *Server) IncrementVersion() {
	s.Version++
}
