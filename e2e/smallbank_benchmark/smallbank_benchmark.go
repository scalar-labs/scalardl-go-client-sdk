package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/scalar-labs/dl/v3"
	client_config "github.com/scalar-labs/dl/v3/client/config"
	client_error "github.com/scalar-labs/dl/v3/client/error"
	client_service "github.com/scalar-labs/dl/v3/client/service"
	"github.com/scalar-labs/dl/v3/ledger/statuscode"
)

var (
	propertiesFile = flag.String("properties", "client.properties", "the properties file")
	accountNum     = flag.Int("num-accounts", 10000, "the number of target accounts")
	concurrencyNum = flag.Int("num-concurrencies", 1, "the number of concurrencies to run")
	duration       = flag.Int("duration", 200, "the duration of benchmark in seconds")
	rampUp         = flag.Int("ramp-up-time", 30, "the ramp up time in seconds")
)

func main() {
	flag.Parse()

	var (
		service client_service.ClientService
		err     error
	)

	if service, err = createClientService(*propertiesFile); err != nil {
		printError(err)
		os.Exit(1)
	}
	defer service.Close()

	if err = service.RegisterCertificate(); err != nil {
		if clientErr, ok := err.(client_error.ClientError); ok {
			if clientErr.StatusCode() != statuscode.CertificateAlreadyRegistered {
				printError(err)
				os.Exit(1)
			}
		} else {
			printError(err)
			os.Exit(1)
		}
	}

	if err = registerContracts(service); err != nil {
		if clientErr, ok := err.(client_error.ClientError); ok {
			if clientErr.StatusCode() != statuscode.ContractAlreadyRegistered {
				printError(err)
				os.Exit(1)
			}
		} else {
			printError(err)
			os.Exit(1)
		}
	}

	createAccounts(service)

	var (
		start = time.Now().UnixMilli()
		end   = start + int64(*duration*1000) + int64(*rampUp*1000)

		ctx, cancel               = context.WithCancel(context.Background())
		transactions      uint32  = 0
		totalTransactions uint32  = 0
		totalLatency      int64   = 0
		errorCounts       uint32  = 0
		tps               float64 = 0
		avgLatency        float64 = 0
	)

	for i := 0; i < *concurrencyNum; i++ {
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					operation, argument := createRequest()
					eachStart := time.Now().UnixMilli()

					atomic.AddUint32(&transactions, 1)

					if _, err := service.ExecuteContract(operation, argument, nil); err == nil {
						if time.Now().UnixMilli() >= start+int64(*rampUp*1000) {
							atomic.AddUint32(&totalTransactions, 1)
							atomic.AddInt64(&totalLatency, time.Now().UnixMilli()-eachStart)
						}
					} else {
						fmt.Println(err)
						atomic.AddUint32(&errorCounts, 1)
					}
				}
			}
		}(ctx)
	}

	var from = start
	for {
		var to = time.Now().UnixMilli()

		if to >= end {
			cancel()
			time.Sleep(time.Second)
			break
		}

		if to-from > 0 {
			fmt.Printf("%.2f tps\n", float32(transactions)/float32(to-from)*1000)
		}

		atomic.StoreUint32(&transactions, 0)
		from = time.Now().UnixMilli()

		time.Sleep(time.Second)
	}

	tps = float64(totalTransactions) / float64(*duration)

	if totalTransactions != 0 {
		avgLatency = float64(totalLatency) / float64(totalTransactions)
	}

	fmt.Printf("TPS: %.2f\n", tps)
	fmt.Printf("Average-Latency(ms): %.2f\n", avgLatency)
	fmt.Printf("Error-Counts: %d\n", errorCounts)
}

func createRequest() (operation string, argument dl.JSONObject) {
	argument = dl.JSONObject{}

	var operations []string = []string{
		"transact_savings",
		"deposit_checking",
		"write_check",
		"send_payment",
		"amalgamate",
	}

	var account1, account2 int = rand.Intn(*accountNum), rand.Intn(*accountNum)
	var amount int = rand.Intn(100) + 1

	if account1 == account2 {
		account1 = (account1 + 1) % *accountNum
	}

	operation = operations[rand.Intn(len(operations))]

	switch operation {
	case "transact_savings", "deposit_checking", "write_check":
		argument["customer_id"] = account1
		argument["amount"] = amount

	case "send_payment":
		argument["source_customer_id"] = account1
		argument["dest_customer_id"] = account2
		argument["amount"] = amount

	case "amalgamate":
		argument["source_customer_id"] = account1
		argument["dest_customer_id"] = account2
	}

	return
}

func printError(err error) {
	if clientErr, ok := err.(client_error.ClientError); ok {
		fmt.Println(clientErr)
	} else {
		fmt.Println(err)
	}
}

func createClientService(file string) (service client_service.ClientService, err error) {
	var properties []byte
	var config client_config.ClientConfig

	if properties, err = ioutil.ReadFile(*propertiesFile); err != nil {
		return
	}

	if config, err = client_config.NewClientConfigFromJavaProperties(string(properties)); err != nil {
		return
	}

	return client_service.NewClientService(config)
}

func registerContracts(service client_service.ClientService) (err error) {
	var (
		contractBytes      []byte
		contractProperties dl.JSONObject
	)

	if contractBytes, err = os.ReadFile("Amalgamate.class"); err != nil {
		return
	}
	if err = service.RegisterContract(
		"amalgamate",
		"com.example.contract.smallbank.Amalgamate",
		contractBytes,
		contractProperties,
	); err != nil {
		return
	}

	if contractBytes, err = os.ReadFile("CreateAccount.class"); err != nil {
		return
	}
	if err = service.RegisterContract(
		"create_account",
		"com.example.contract.smallbank.CreateAccount",
		contractBytes,
		contractProperties,
	); err != nil {
		return
	}

	if contractBytes, err = os.ReadFile("DepositChecking.class"); err != nil {
		return
	}
	if err = service.RegisterContract(
		"deposit_checking",
		"com.example.contract.smallbank.DepositChecking",
		contractBytes,
		contractProperties,
	); err != nil {
		return
	}

	if contractBytes, err = os.ReadFile("SendPayment.class"); err != nil {
		return
	}
	if err = service.RegisterContract(
		"send_payment",
		"com.example.contract.smallbank.SendPayment",
		contractBytes,
		contractProperties,
	); err != nil {
		return
	}

	if contractBytes, err = os.ReadFile("TransactSavings.class"); err != nil {
		return
	}
	if err = service.RegisterContract(
		"transact_savings",
		"com.example.contract.smallbank.TransactSavings",
		contractBytes,
		contractProperties,
	); err != nil {
		return
	}

	if contractBytes, err = os.ReadFile("WriteCheck.class"); err != nil {
		return
	}
	if err = service.RegisterContract(
		"write_check",
		"com.example.contract.smallbank.WriteCheck",
		contractBytes,
		contractProperties,
	); err != nil {
		return
	}

	return
}

func createAccounts(s client_service.ClientService) (err error) {
	if *concurrencyNum > *accountNum {
		*concurrencyNum = *accountNum
	}

	var (
		reminder          = *accountNum % *concurrencyNum
		chunkSize         = (*accountNum - reminder) / *concurrencyNum
		chunks    [][]int = make([][]int, 0)
		chunk     []int   = make([]int, 0)
	)

	for i := 0; i < *accountNum; i++ {
		chunk = append(chunk, i)

		if len(chunk) >= chunkSize {
			chunks = append(chunks, chunk)
			chunk = make([]int, 0)
		}
	}

	var i = 0
	for _, a := range chunk {
		chunks[i] = append(chunks[i], a)
		i = (i + 1) % *concurrencyNum
	}

	var wg sync.WaitGroup

	for i = 0; i < *concurrencyNum; i++ {
		var accounts []int = chunks[i]

		wg.Add(1)

		go func(s client_service.ClientService, accounts []int) {
			for _, a := range accounts {
				if _, err = s.ExecuteContract(
					"create_account",
					dl.JSONObject{
						"customer_id":              a,
						"customer_name":            fmt.Sprintf("Number %d", a),
						"initial_checking_balance": 100000,
						"initial_savings_balance":  100000,
					},
					nil,
				); err == nil {
					fmt.Printf("created: %d\n", a)
				}
			}

			wg.Done()
		}(s, accounts)
	}

	wg.Wait()

	return
}
