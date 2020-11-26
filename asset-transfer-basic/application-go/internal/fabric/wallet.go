package fabric

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"github.com/nkmr-jp/go-logger-scaffold/logger"
)

// TODO: get from config
const (
	walletPath  = "wallet"
	walletLabel = "appUser"
	mspID       = "Org1MSP"
)

type Wallet struct {
	wallet *gateway.Wallet
	label  string
	mspID  string
}

func newWallet() *Wallet {
	wallet, err := gateway.NewFileSystemWallet(walletPath)
	if err != nil {
		logger.Fatalf("Failed to create wallet", err)
	}
	return &Wallet{
		wallet: wallet,
		label:  walletLabel,
		mspID:  mspID,
	}
}

func (w *Wallet) build() error {
	if w.wallet == nil {
		logger.Error("WALLET_IS_NOT_INITIALIZED")
	}
	if w.wallet.Exists(w.label) {
		logger.Info("ALREADY_BUILT")
		return nil
	}

	// get certificate
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	// get private key
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

	// creates an X509 identity
	identity := gateway.NewX509Identity(w.mspID, string(cert), string(key))
	if err := w.wallet.Put(w.label, identity); err != nil {
		return err
	}

	return nil
}
