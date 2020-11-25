package fabric

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"github.com/nkmr-jp/go-logger-scaffold/logger"
)

var (
	credPath = filepath.Join(
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
	certPath = filepath.Join(credPath, "signcerts", "cert.pem")
	keyDir   = filepath.Join(credPath, "keystore")
)

func createWallet() *gateway.Wallet {
	wallet, err := gateway.NewFileSystemWallet(walletPath)
	if err != nil {
		logger.Fatalf("Failed to create wallet", err)
	}

	if !wallet.Exists(walletLabel) {
		err = populateWallet(wallet)
		if err != nil {
			logger.Fatalf("Failed to populate wallet contents", err)
		}
	}
	return wallet
}

func populateWallet(wallet *gateway.Wallet) error {
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}
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

	identity := gateway.NewX509Identity(mspID, string(cert), string(key))

	return wallet.Put(walletLabel, identity)
}
