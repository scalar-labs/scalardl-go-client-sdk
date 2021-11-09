# Re-generate gRPC protobuf files

Check [prerequsisites](https://grpc.io/docs/languages/go/quickstart/#prerequisites) to install necessary tools.
We also need to add `option go_package = "./rpc";` to the *.proto files to specify the Golang package.
Then, use the following command to generate related files.
```
protoc --go_out=<path-to-this-repository> --go-grpc_out=<path-to-this-repository> scalar.proto
```

