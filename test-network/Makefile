# See https://hyperledger-fabric.readthedocs.io/en/release-2.2/test_network.html
# See https://hyperledger-fabric.readthedocs.io/en/release-2.2/deploy_chaincode.html

CTOR=""
CC_PACKAGE_ID=""

env:
	@echo "[\"$$CORE_PEER_LOCALMSPID\"]"

# networkを再構築
restart: down up

# networkを再構築してchaincode(basic)をデプロイ
restart-with-cc: restart
	@echo ""
	@echo "-------- deployCC --------"
	@echo ""
	./network.sh deployCC

# networkを構築してチャネルを設定(mychannel)
up:
	@echo ""
	@echo "-------- network up createChannel --------"
	@echo ""
	./network.sh up createChannel

# networkを削除
down:
	@echo ""
	@echo "-------- network down. --------"
	@echo ""
	if docker ps -a | grep -e logspout; then docker stop logspout; docker rm logspout; fi;
	./network.sh down

# chaincodeを実行(invoke)
invoke: env
	@peer chaincode invoke -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.example.com --tls \
        --cafile $$PWD/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
        -C mychannel -n basic --peerAddresses localhost:7051 \
        --tlsRootCertFiles $$PWD/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt \
        --peerAddresses localhost:9051 \
        --tlsRootCertFiles $$PWD/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt \
        -c '$(CTOR)'

# chaincodeを実行(query)
query: env
	@peer chaincode query -C mychannel -n basic -c '$(CTOR)'

# 全てのdockerコンテナのログを見る
logs:
	@./monitordocker.sh net_test

# 起動しているコンテナの確認
ps:
	@docker ps -a | grep -e fabric -e logspout -e chaincode


# Deploy Chaincode

# chaincodeをパッケージング
lc-pkg: env
	peer lifecycle chaincode package basic.tar.gz \
		--path ../asset-transfer-basic/chaincode-typescript/ \
		--lang node --label basic_1.0

# chaincodeのパッケージをPeerにインストール
lc-install: env
	peer lifecycle chaincode install basic.tar.gz

# インストールされているchaincodeを確認
lc-query: env
	peer lifecycle chaincode queryinstalled

# chaincode定義を承認
lc-approve: env
	peer lifecycle chaincode approveformyorg -o localhost:7050 \
		--ordererTLSHostnameOverride orderer.example.com \
		--channelID mychannel --name basic --version 1.0 \
		--package-id $(CC_PACKAGE_ID) --sequence 1 --tls \
		--cafile $$PWD/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

# chaincodeをチャネルにコミットする準備ができてるか確認
lc-readiness: env
	@peer lifecycle chaincode checkcommitreadiness \
		--channelID mychannel --name basic --version 1.0 --sequence 1 --tls \
		--cafile $$PWD/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
		--output json

# chaincodeをチャネルにコミット
lc-commit: env
	peer lifecycle chaincode commit -o localhost:7050 \
		--ordererTLSHostnameOverride orderer.example.com \
		--channelID mychannel --name basic --version 1.0 --sequence 1 --tls \
		--cafile $$PWD/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
		--peerAddresses localhost:7051 \
		--tlsRootCertFiles $$PWD/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt \
		--peerAddresses localhost:9051 \
		--tlsRootCertFiles $$PWD/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
