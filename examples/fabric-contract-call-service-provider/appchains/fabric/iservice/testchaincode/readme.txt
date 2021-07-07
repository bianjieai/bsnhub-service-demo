

peer chaincode package -n cc_cross -p relayer/appchains/fabric/iservice/chaincode/cmd/ -v 0.1 cc_cross.0.1.pak

peer chaincode install cc_cross.0.1.pak

peer chaincode instantiate -o order1.ordernode.bsnbase.com:17051 -C netchannel -n cc_cross -v 0.1 -c '{"Args":["init"]}' --tls true --cafile /etc/hyperledger/fabric/certs/ordererOrganizations/ordernode.bsnbase.com/orderers/order1.ordernode.bsnbase.com/tls/tlsintermediatecerts/tls-ca-ordernode-bsnbase-com-15901-2.pem

peer chaincode upgrade -o order1.ordernode.bsnbase.com:17051 -C netchannel -n cc_cross -v 0.2 -c '{"Args":["init"]}' --tls true --cafile /etc/hyperledger/fabric/certs/ordererOrganizations/ordernode.bsnbase.com/orderers/order1.ordernode.bsnbase.com/tls/tlsintermediatecerts/tls-ca-ordernode-bsnbase-com-15901-2.pem


peer chaincode invoke -o order1.ordernode.bsnbase.com:17051 -C netchannel -n cc_cross -c '{"Args":["CallService","{\"serviceName\":\"test\",\"input\":\"abc\",\"timeout\":100}"]}' --tls true --cafile /etc/hyperledger/fabric/certs/ordererOrganizations/ordernode.bsnbase.com/orderers/order1.ordernode.bsnbase.com/tls/tlsintermediatecerts/tls-ca-ordernode-bsnbase-com-15901-2.pem

peer chaincode invoke -o order1.ordernode.bsnbase.com:17051 -C netchannel -n cc_cross -c '{"Args":["setResponse","{\"requestID\":\"9e019874c892468edd018acd62ed1063419b849c275f9204e7859d6e4d206bfb\",\"output\":\"aabbcc\",\"icRequestID\":\"5622\"}"]}' --tls true --cafile /etc/hyperledger/fabric/certs/ordererOrganizations/ordernode.bsnbase.com/orderers/order1.ordernode.bsnbase.com/tls/tlsintermediatecerts/tls-ca-ordernode-bsnbase-com-15901-2.pem


peer chaincode query -C netchannel -n cc_cross -c '{"Args":["query","9e019874c892468edd018acd62ed1063419b849c275f9204e7859d6e4d206bfb"]}'

peer chaincode query -C channel202010310000001 -n cc_cross -c '{"Args":["getResponse","fe2267255a2245fee2ce5be1c64cfb7705dd363a0fe8e6347120885250a3d03e"]}'


peer chaincode invoke -o order1.ordernode.bsnbase.com:17051 -C netchannel -n cc_cross -c '{"Args":["addServiceBinding","{\"name\":\"test1\",\"description\":\"abc\",\"schemas\":\"a-z\",\"provider\":\"fabric\",\"serviceFee\":\"100pont\",\"qos\":100}"]}' --tls true --cafile /etc/hyperledger/fabric/certs/ordererOrganizations/ordernode.bsnbase.com/orderers/order1.ordernode.bsnbase.com/tls/tlsintermediatecerts/tls-ca-ordernode-bsnbase-com-15901-2.pem

peer chaincode invoke -o order1.ordernode.bsnbase.com:17051 -C netchannel -n cc_cross -c '{"Args":["updateservicebinding","{\"serviceName\":\"test1\",\"provider\":\"fabric\",\"serviceFee\":\"1000pont\",\"qos\":200}"]}' --tls true --cafile /etc/hyperledger/fabric/certs/ordererOrganizations/ordernode.bsnbase.com/orderers/order1.ordernode.bsnbase.com/tls/tlsintermediatecerts/tls-ca-ordernode-bsnbase-com-15901-2.pem


peer chaincode query -C netchannel -n cc_cross -c '{"Args":["getServiceBinding","test1"]}'

peer chaincode query -C netchannel -n cc_cross -c '{"Args":["getServiceBindings"]}'


peer chaincode invoke -o order1.ordernode.bsnbase.com:17051 -C netchannel -n cc_cross -c '{"Args":["setResponse","{\"requestID\":\"55c5b5571573e29c9de181ff84af9f3886a6b26944374d53c945970ed317c177\",\"output\":\"aabbcc\",\"icRequestID\":\"5622\"}"]}' --tls true --cafile /etc/hyperledger/fabric/certs/ordererOrganizations/ordernode.bsnbase.com/orderers/order1.ordernode.bsnbase.com/tls/tlsintermediatecerts/tls-ca-ordernode-bsnbase-com-15901-2.pem
