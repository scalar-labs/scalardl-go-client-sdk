package main

import (
	"flag"
	"io/ioutil"
	"log"

	client_config "github.com/scalar-labs/scalardl-go-client-sdk/v3/client/config"
	client_error "github.com/scalar-labs/scalardl-go-client-sdk/v3/client/error"
	client_service "github.com/scalar-labs/scalardl-go-client-sdk/v3/client/service"
)

var propertiesFile = flag.String("properties", "client.properties", "the properties file")

func main() {
	flag.Parse()

	var (
		clientConfig client_config.ClientConfig
		properties   []byte
		err          error
	)

	if properties, err = ioutil.ReadFile(*propertiesFile); err != nil {
		log.Panicln(err)
	}

	if clientConfig, err = client_config.NewClientConfigFromJavaProperties(string(properties)); err != nil {
		log.Panicln(err)
	}

	var service client_service.ClientService
	if service, err = client_service.NewClientService(clientConfig); err != nil {
		log.Panicln(err)
	}
	defer service.Close()

	if err = service.RegisterCertificate(); err != nil {
		if clientError, ok := err.(client_error.ClientError); ok {
			log.Panicf(
				"%d %s\n",
				clientError.StatusCode(),
				clientError.Error(),
			)
		} else {
			log.Panicln(err)
		}
	}
}
