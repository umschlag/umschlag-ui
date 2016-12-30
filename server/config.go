package main

type server struct {
	Host        string
	Addr        string
	Cert        string
	Key         string
	Root        string
	Endpoint    string
	Storage     string
	LetsEncrypt bool
}

type config struct {
	Debug  bool
	Server *server
}

var (
	// Config represents the general server configuration.
	Config = &config{
		Server: &server{},
	}
)
