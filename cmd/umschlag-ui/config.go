package main

type server struct {
	Host          string
	Addr          string
	Cert          string
	Key           string
	Endpoint      string
	Static        string
	Storage       string
	LetsEncrypt   bool
	StrictCurves  bool
	StrictCiphers bool
	Pprof         bool
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
