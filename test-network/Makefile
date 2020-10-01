# See https://hyperledger-fabric.readthedocs.io/en/release-2.2/test_network.html
# See https://hyperledger-fabric.readthedocs.io/en/release-2.2/deploy_chaincode.html

CTOR=""
CC_PACKAGE_ID=""

env:
	@echo "[\"$$CORE_PEER_LOCALMSPID\"]"

restart: down up

restart-with-cc: restart
	@echo ""
	@echo "-------- deployCC --------"
	@echo ""
	./network.sh deployCC

up:
	@echo ""
	@echo "-------- network up createChannel --------"
	@echo ""
	./network.sh up createChannel

down:
	@echo ""
	@echo "-------- network down. --------"
	@echo ""
	if docker ps -a | grep -e logspout; then docker stop logspout; docker rm logspout; fi;
	./network.sh down

invoke: env
	@peer chaincode invoke -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.example.com --tls \
        --cafile $$PWD/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
        -C mychannel -n basic --peerAddresses localhost:7051 \
        --tlsRootCertFiles $$PWD/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt \
        --peerAddresses localhost:9051 \
        --tlsRootCertFiles $$PWD/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt \
        -c '$(CTOR)'

query: env
	@peer chaincode query -C mychannel -n basic -c '$(CTOR)'

logs:
	@./monitordocker.sh net_test

ps:
	@docker ps -a | grep -e fabric -e logspout -e chaincode

lc-pkg: env
	peer lifecycle chaincode package basic.tar.gz \
		--path ../asset-transfer-basic/chaincode-typescript/ \
		--lang node --label basic_1.0

lc-install: env
	peer lifecycle chaincode install basic.tar.gz

lc-query: env
	peer lifecycle chaincode queryinstalled

lc-approve: env
	peer lifecycle chaincode approveformyorg -o localhost:7050 \
		--ordererTLSHostnameOverride orderer.example.com \
		--channelID mychannel --name basic --version 1.0 \
		--package-id $(CC_PACKAGE_ID) --sequence 1 --tls \
		--cafile $$PWD/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

lc-commit: env
	peer lifecycle chaincode checkcommitreadiness \
		--channelID mychannel --name basic --version 1.0 --sequence 1 --tls \
		--cafile $$PWD/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
		--output json
