package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

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
		fmt.Println(err)
		os.Exit(1)
	}

	if clientConfig, err = client_config.NewClientConfigFromJavaProperties(string(properties)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var service client_service.ClientService
	if service, err = client_service.NewClientService(clientConfig); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer service.Close()

	var (
		argument []interface{} = []interface{}{*assetID, *startAge, *endAge}
		result   model.LedgerValidationResult
	)

	if result, err = service.ValidateLedger(argument...); err != nil {
		if clientError, ok := err.(client_error.ClientError); ok {
			fmt.Printf(
				"%d %s\n",
				clientError.StatusCode(),
				clientError.Error(),
			)
		} else {
			fmt.Println(err)
		}

		os.Exit(1)
	}

	fmt.Printf("Status Code: %d\n", result.Code)
	fmt.Printf("Proof of Ledger: %v\n", result.Proof)
	if !result.AuditorProof.Equal(asset.Proof{}) {
		fmt.Printf("Proof of Auditor: %v\n", result.AuditorProof)

	}
}
