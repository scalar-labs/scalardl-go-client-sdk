package main

import (
	"flag"
	"io/ioutil"
	"log"

	client_config "github.com/scalar-labs/scalardl-go-client-sdk/v3/client/config"
	client_error "github.com/scalar-labs/scalardl-go-client-sdk/v3/client/error"
	client_service "github.com/scalar-labs/scalardl-go-client-sdk/v3/client/service"
	"github.com/scalar-labs/scalardl-go-client-sdk/v3/ledger/asset"
	"github.com/scalar-labs/scalardl-go-client-sdk/v3/ledger/model"
)

var propertiesFile = flag.String("properties", "client.properties", "the properties file")
var assetID = flag.String("asset_id", "", "the asset ID")
var startAge = flag.Int("start_age", 0, "the start age of the asset")
var endAge = flag.Int("end_age", client_service.JavaMaxIntValue, "the end age of the asset")

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

	var (
		argument []interface{} = []interface{}{*assetID, *startAge, *endAge}
		result   model.LedgerValidationResult
	)

	if result, err = service.ValidateLedger(argument...); err != nil {
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

	log.Printf("Status Code: %d\n", result.Code)
	log.Printf("Proof of Ledger: %v\n", result.Proof)
	if !result.AuditorProof.Equal(asset.Proof{}) {
		log.Printf("Proof of Auditor: %v\n", result.AuditorProof)
	}
}
