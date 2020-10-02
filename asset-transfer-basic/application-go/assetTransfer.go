/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"asset-transfer-basic/logger"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger()
	defer logger.Sync()

	logger.Info("============ application-golang starts ============")

	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		logger.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable", err)
	}

	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		logger.Fatalf("Failed to create wallet", err)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			logger.Fatalf("Failed to populate wallet contents", err)
		}
	}

	ccpPath := filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		logger.Fatalf("Failed to connect to gateway", err)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		logger.Fatalf("Failed to get network", err)
	}

	contract := network.GetContract("basic")

	logger.Info("Submit Transaction: InitLedger")
	result, err := contract.SubmitTransaction("InitLedger")
	if err != nil {
		logger.Errorf("Failed to Submit transaction", err)
	}
	logger.Info(string(result))

	logger.Info("Evaluate Transaction: GetAllAssets")
	result, err = contract.EvaluateTransaction("GetAllAssets")
	if err != nil {
		logger.Fatalf("Failed to evaluate transaction", err)
	}
	logger.Info(string(result))

	logger.Info("Submit Transaction: CreateAsset")
	result, err = contract.SubmitTransaction("CreateAsset", "asset13", "yellow", "5", "Tom", "1300")
	if err != nil {
		logger.Fatalf("Failed to Submit transaction", err)
	}
	logger.Info(string(result))

	logger.Info("Evaluate Transaction: ReadAsset")
	result, err = contract.EvaluateTransaction("ReadAsset", "asset13")
	if err != nil {
		logger.Fatalf("Failed to evaluate transaction", err)
	}
	logger.Info(string(result))

	logger.Info("Evaluate Transaction: AssetExists", zap.String("detail", ", function returns 'true' if an asset with given assetID exist"))
	result, err = contract.EvaluateTransaction("AssetExists", "asset1")
	if err != nil {
		logger.Fatalf("Failed to evaluate transaction", err)
	}
	logger.Info(string(result))

	logger.Info("Submit Transaction: TransferAsset asset1", zap.String("detail", "transfer to new owner of Tom"))
	_, err = contract.SubmitTransaction("TransferAsset", "asset1", "Tom")
	if err != nil {
		logger.Fatalf("Failed to Submit transaction", err)
	}

	logger.Info("Evaluate Transaction: ReadAsset", zap.String("detail", "function returns 'asset1' attributes"))
	result, err = contract.EvaluateTransaction("ReadAsset", "asset1")
	if err != nil {
		logger.Fatalf("Failed to evaluate transaction", err)
	}
	logger.Info(string(result))
	logger.Info("============ application-golang ends ============")
}

func populateWallet(wallet *gateway.Wallet) error {
	logger.Info("============ Populating wallet ============")
	credPath := filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	return wallet.Put("appUser", identity)
}
