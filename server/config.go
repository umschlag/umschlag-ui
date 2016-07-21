package main

type server struct {
	Addr     string
	Cert     string
	Key      string
	Root     string
	Endpoint string
}

type config struct {
	Update bool
	Debug  bool
	Server *server
}

var (
	// Config represents the general server configuration.
	Config = &config{
		Server: &server{},
	}
)
