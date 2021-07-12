package config

import (
	"errors"
	"flag"
	"log"
	"net/url"
	"os"
)

type conf struct {
	Port string
}

func IsUrl(str string) error {
	_, err := url.ParseRequestURI(str)
	if err == nil {
		err = errors.New("Passed")
	}
	return err
}

func ReadConfig() conf {
	configPort, err := os.ReadFile("server.cfg")
	if err != nil {
		log.Fatalln(err)
	}

	var port = flag.String("port", string(configPort), "Port number")
	flag.Parse()

	config := conf{Port: *port}

	return config
}
