package dphx

import (
	"log"

	"github.com/MOZGIII/go-socks5"
	"github.com/kelseyhightower/envconfig"
)

var appConfig Config

// Run initializes config and starts SOCKS server.
func Run() {
	envconfig.MustProcess("DPHX", &appConfig)
	ListenAndServe()
}

// ListenAndServe starts the SOCKS server.
func ListenAndServe() {
	// Create a SOCKS5 server
	conf := &socks5.Config{
		Dial:     SSHDial,
		Resolver: DummyResolver{},
	}
	server, err := socks5.New(conf)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("SOCKS5 server is starting at %s", appConfig.LocalAddr)
	if err := server.ListenAndServe("tcp", appConfig.LocalAddr); err != nil {
		log.Fatalf(err.Error())
	}
}
