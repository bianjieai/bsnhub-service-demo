# bsnhub-service-demo
bsnhub service demo - marketservice daemon for service provider

## Install
```bash
make install
```

## Run

### Precondition

Make sure `iritacli` is installed and `BSNHub` is accessible.

### Create key pair for service provider and consumer
```bash
# set environment variable
chain_id=test
service_name=price_service
root_account=<your root account>

# generate provider address
iritacli keys add provider --keyring-backend file

# generate consumer address
iritacli keys add consumer --keyring-backend file

# send some token to provider address
iritacli tx send $root_account $(iritacli keys show provider --keyring-backend file -o json | jq -r '.address') 1000000point --chain-id $chain_id -b block -y

# send some token to consumer address
iritacli tx send $root_account $(iritacli keys show consumer --keyring-backend file -o json | jq -r '.address') 1000000point --chain-id $chain_id -b block -y
```

### Create service definition
```bash
# set provider address as a PowerUser
iritacli tx admin add-roles $(iritacli keys show provider --keyring-backend file -o json | jq -r '.address') PowerUser --chain-id $chain_id -b block -y --from $root_account

# define service
iritacli tx service define --chain-id $chain_id --from provider --name $service_name --description="provide token price" --tags=price --schemas=marketservice/service/service_definition.json -b block -y --keyring-backend file
```

### Create service binding
```bash
# bind service
iritacli tx service bind --chain-id $chain_id --from provider --service-name $service_name --deposit=100000point --qos=50 --pricing marketservice/service/service_pricing.json -b block -y --keyring-backend file

# qury bindings
iritacli query service bindings $service_name --chain-id $chain_id
```

### Start marketservice daemon
```bash
marketservice start marketservice start [chain-id] [node-uri] provider [password] binance
```

### Call service
```bash
# call service
iritacli tx service call --chain-id $chain_id --from consumer --service-name $service_name --data "{\"base\":\"iris\",\"quote\":\"usdt\"}" --providers $(iritacli keys show provider --keyring-backend file -o json | jq -r '.address') --service-fee-cap 1point --timeout 50 --frequency 5 -b block -y --keyring-backend file
```

### Query request & response
```bash
request_id = <your_request_id>
iritacli query service request $request_id --chain-id $chain_id
iritacli query service response $request_id --chain-id $chain_id
```
