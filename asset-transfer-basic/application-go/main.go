package main

import (
	"asset-transfer-basic/internal/fabric"

	"github.com/nkmr-jp/go-logger-scaffold/logger"
)

func main() {
	logger.InitLogger()
	defer logger.Sync()
	fabric.InitFabric()
	defer fabric.Close()
}

func run() {
	c := fabric.GetContract()
	result, err := c.SubmitTransaction("InitLedger")
	if err != nil {
		logger.Fatalf("Failed to Submit transaction", err)
	}
	logger.Info(string(result))

	result, err = c.EvaluateTransaction("GetAllAssets")
	if err != nil {
		logger.Fatalf("Failed to evaluate transaction", err)
	}
	logger.Info(string(result))
}
