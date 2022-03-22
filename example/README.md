In this folder, we can see four main packages in different sub-folders, register_certificate, register_contract, execute_contract and validate_ledger. They demonstrate how to use Scalar DL Go Client SDK to send basic operation requests to Scalar DL networks.

To build all of them
```
cd <sub-folder>
go build
```

## Local Scalar DL Network
We need a local Auditor-enabled Scalar DL network before we run these examples.

Reference [scalardl-samples
](https://github.com/scalar-labs/scalardl-samples) to start up local Scalar DL networks.

## client.properties
The Scalar DL Client Properties example can be found in [client.properties](client.properties).

This properties example is configured to connect to a local and auditor-enabled Scalar DL network.

### register_certificate
Run
```
./register_certificate/register_certificate -properties client.properties
```
to register the client certificate of the example.

### register_contract
Run
```
./register_contract/register_contract -properties client.properties -contract StateReader.class -id state-reader -name com.org1.contract.StateReader
```

```
./register_contract/register_contract -properties client.properties -contract StateUpdater.class -id state-updater -name com.org1.contract.StateUpdater
```

```
./register_contract/register_contract -properties client.properties -contract ValidateLedger.class -id validate-ledger -name com.scalar.dl.client.contract.ValidateLedger
```

to register three contracts of the example.

The implementation of these contracts can be found [here](https://github.com/scalar-labs/scalardl-java-client-sdk/tree/master/src/main/java/com).

### execute_contract
Run
```
./execute_contract/execute_contract -properties client.properties -id state-updater -argument '{"asset_id":"foo","state":1}'
```

to update the asset `foo`'s state to be `1`.

Run
```
./execute_contract/execute_contract -properties client.properties -id state-reader -argument '{"asset_id":"foo"}'
```
to check the asset's detail.

### validate_ledger
Run
```
./validate_ledger/validate_ledger -properties client.properties -asset_id foo
```
to validate if the asset foo has been tampered with.
