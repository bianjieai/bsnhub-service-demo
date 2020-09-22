# Integration

Irita is built base on cosmos-sdk and tendermint, which supports interoperability between various heterogeneous chains. This is an integration example showing it.

There are three services in this integration example:

1. Oracle service: provided by chainlink.
2. Mintbase service: provided by Ethereum.
3. Record service: provided by irishub.

Consumer can call services or cross-chain services through iritahub or other chains which connect to iritahub with [service relayer](https://github.com/bianjieai/bsnhub-service-relayer). 

In this example, we initiate all service invocation from Ethereum(or Fabric...), and then all service invocation will be relayed to iritahub, and the on-chain or off-chain service provider process responds the invocation. Eventually, the response will be relayed back to Ethereum(or Fabric...) and the consumer will get the response.

### Install

* Mintbase service provider daemon

  ```bash
  make install-nftservice
  ```

* Record service provider daemon

  ```bash
  make install-recordservice
  ```

* chainlink

  https://github.com/bianjieai/chainlink/tree/v0.8.13-irita

* relayer

  https://github.com/bianjieai/bsnhub-service-relayer

### Precondition

1. Create Mintbase contract on rinkeby(ethereum testnet)

   https://rinkeby.mintbase.io/my-market/0x6a38659237feedd73c333a1e5ce8535542f14ab8 

2. Register services on iritahub

   > The service definitions and bindings should be automatically published on Ethereum by the relayer.

   * register Mintbase service

     ```bash
     service_name=mintbase_service
     
     iritacli tx service define --chain-id $chain_id --from provider --name $service_name --description="provide token price" --tags=price --schemas=/examples/nft/service/service_definition.json -b block -y --keyring-backend file
     
     iritacli tx service bind --chain-id $chain_id --from provider --service-name $service_name --deposit=100000point --qos=50 --pricing /examples/nft/service/service_pricing.json -b block -y --keyring-backend file
     ```

   * register record service

     ```bash
     service_name=record_service
     
     iritacli tx service define --chain-id $chain_id --from provider --name $service_name --description="provide token price" --tags=price --schemas=/examples/record/service/service_definition.json -b block -y --keyring-backend file
     
     iritacli tx service bind --chain-id $chain_id --from provider --service-name $service_name --deposit=100000point --qos=50 --pricing /examples/record/service/service_pricing.json -b block -y --keyring-backend file
     ```

### RUN

1. create `eth_usdt` feed and start chainlink node 
2. start mintbase service daemon & record service daemon

3. call Oracle service, get the `ETH/USDT` price
4. call Mintbase service  to mint NFT
5. call Record service



### TODO

1. update irita using cosmos-sdk v0.40 (WIP)
2. update irishub-go-sdk
3. update irita-go-sdk
4. update chainlink-node using latest irita-sdk
5. consumer process