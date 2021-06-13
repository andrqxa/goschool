package apiserver

// APIServer ...
type APIServer struct {
	config *Config
}

// New ...
func New(newconfig *Config) *APIServer {
	return &APIServer{
		config: newconfig,
	}
}

// Start ...
func (s *APIServer) Start() error {
	return nil
}
