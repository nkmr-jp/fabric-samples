package main

import (
	"asset-transfer-basic/internal/fabric"

	"github.com/nkmr-jp/go-logger-scaffold/logger"
)

const (
	channelID   = "mychannel"
	chaincodeID = "basic"
)

func main() {
	logger.InitLogger()
	defer logger.Sync()
	fabric.InitFabric()
	defer fabric.Close()
	run()
}

func run() {
	newChannel, err := fabric.NewChannel(channelID)
	if err != nil {
		logger.Errorf("NEW_CHANNEL_ERROR", err)
	}
	contract := newChannel.GetContract(chaincodeID)

	result, err := contract.SubmitTransaction("InitLedger")
	if err != nil {
		logger.Fatalf("INIT_LEDGER_ERROR", err)
	}
	logger.Info(string(result))

	result, err = contract.EvaluateTransaction("GetAllAssets")
	if err != nil {
		logger.Fatalf("GET_ALL_ASSETS_ERROR", err)
	}
	logger.Info(string(result))
}
