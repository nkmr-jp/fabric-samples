## Running the test network

You can use the `./network.sh` script to stand up a simple Fabric test network. The test network has two peer organizations with one peer each and a single node raft ordering service. You can also use the `./network.sh` script to create channels and deploy chaincode. For more information, see [Using the Fabric test network](https://hyperledger-fabric.readthedocs.io/en/latest/test_network.html). The test network is being introduced in Fabric v2.0 as the long term replacement for the `first-network` sample.

Before you can deploy the test network, you need to follow the instructions to [Install the Samples, Binaries and Docker Images](https://hyperledger-fabric.readthedocs.io/en/latest/install.html) in the Hyperledger Fabric documentation.


## 追加メモ

```sh
# CONFIGのPATHを追加する
source cfgenv.sh

# ネットワークを再構築して、チャネルを作って、CC(basic)をデプロイする
make restart-with-cc
# 台帳を初期化
make invoke CTOR='{"function":"InitLedger","Args":[]}'
# 所有者を変更
make invoke CTOR='{"function":"TransferAsset","Args":["asset6","Christopher"]}'
# 全てのAssetを取得
make query CTOR='{"Args":["GetAllAssets"]}' | jq -c .[]
# Assetを指定して取得
make query CTOR='{"Args":["ReadAsset","asset6"]}' | jq .
# net_testのコンテナの全てのログを表示
make logs
# fabric関連のコンテナを表示
make ps

# peerを切り替える
source org1env.sh
source org2env.sh
```