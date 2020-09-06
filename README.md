# bsnhub-service-demo
bsnhub service demo - iservice daemon for service provider

## BUILD
```bash
make install
```

## RUNNING

### Precondition

Make sure `iritacli` is installed.

### Create key pair for service providers
```bash
# generate provider address
iritacli keys add provider --keyring-backend file

# generate consumer address
iritacli keys add consumer --keyring-backend file
```

### Create service definition
```bash
# set chain_id
chain_id=test

# set service name
service_name=price_service

# set provider address as a PowerUser
iritacli tx admin add-roles $(iritacli keys show provider --keyring-backend file -o json | jq -r '.address') PowerUser --chain-id $chain_id -b block --from <RootAdmin_address>

# define service use PowerUser
iritacli tx service define --chain-id $chain_id test --from <PowerUser_account> --name $service_name --description="provide token price" --tags=price --schemas=iservice/service/service_definition.json -b block --keyring-backend file
```

### Create service binding
```bash
# send 1000000point to provider address
iritacli tx send <your_account> $(iritacli keys show provider --keyring-backend file -o json | jq -r '.address') 1000000point --chain-id test -b block

# bind service
iritacli tx service bind --chain-id $chain_id --from provider --service-name $service_name --deposit=100000point --qos=50 --pricing iservice/service/service_pricing.json -b block --keyring-backend file

# qury bindings
iritacli query service bindings $service_name --chain-id $chain_id
```

### Start iservice daemon
```bash
iservice start provider huobi
```

### Call service
```bash
# send 1000000point to consumer address
iritacli tx send <your_account> $(iritacli keys show consumer --keyring-backend file -o json | jq -r '.address') 1000000point --chain-id test -b block

# call service
iritacli tx service call --chain-id $chain_id --from consumer --service-name $service_name --data "{\"base\":\"link\",\"quote\":\"usdt\"}" --providers $(iritacli keys show provider --keyring-backend file -o json | jq -r '.address') --service-fee-cap 1point --timeout 2 --frequency 5 -b block --keyring-backend file
```

### Query request & response
```bash
iritacli query service request B75DF797D0A3A39A9FD15E6E3CDD9EA3F86154E35CE86C66861C50527F32FBD60000000000000000000000000000000100000000000008A00000 --chain-id $chain_id
iritacli query service response B75DF797D0A3A39A9FD15E6E3CDD9EA3F86154E35CE86C66861C50527F32FBD60000000000000000000000000000000100000000000008A00000 --chain-id $chain_id
```
