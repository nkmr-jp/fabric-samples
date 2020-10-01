# See https://hyperledger-fabric.readthedocs.io/en/release-2.2/test_network.html
# See https://hyperledger-fabric.readthedocs.io/en/release-2.2/deploy_chaincode.html

CTOR=""
CC_PACKAGE_ID=""

restart:
	./network.sh down
	./network.sh up createChannel
	./network.sh deployCC

invoke:
	@peer chaincode invoke -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.example.com --tls \
        --cafile $$PWD/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
        -C mychannel -n basic --peerAddresses localhost:7051 \
        --tlsRootCertFiles $$PWD/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt \
        --peerAddresses localhost:9051 \
        --tlsRootCertFiles $$PWD/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt \
        -c '$(CTOR)'
query:
	@peer chaincode query -C mychannel -n basic -c '$(CTOR)'

logs:
	@./monitordocker.sh net_test

ps:
	@docker ps -a | grep -e fabric -e logspout

approve:
	peer lifecycle chaincode approveformyorg -o localhost:7050 \
		--ordererTLSHostnameOverride orderer.example.com \
		--channelID mychannel --name basic --version 1.0 \
		--package-id $(CC_PACKAGE_ID) --sequence 1 --tls \
		--cafile $$PWD/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem