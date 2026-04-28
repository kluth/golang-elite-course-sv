package assignment

type Server struct { Port int }
type Option func(*Server)

func WithPort(p int) Option {
	panic("implement me")
}

func NewServer(opts ...Option) *Server {
	s := &Server{}
	for _, o := range opts { o(s) }
	return s
}
