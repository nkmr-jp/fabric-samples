package fabric

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"github.com/nkmr-jp/go-logger-scaffold/logger"
)

var (
	once   sync.Once
	fabric *gateway.Gateway
)

// InitFabric is initialize the entry point to a Fabric network
func InitFabric() *gateway.Gateway {
	once.Do(func() {
		if err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true"); err != nil {
			logger.Fatalf("SET_ENV_ERROR", err)
		}

		wallet := newWallet()
		if err := wallet.build(); err != nil {
			logger.Errorf("WALLET_BUILD_ERROR", err)
		}

		gw, err := gateway.Connect(
			gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
			gateway.WithIdentity(wallet.wallet, wallet.label),
		)
		if err != nil {
			logger.Fatalf("CONNECT_TO_GATEWAY_ERROR", err)
		}

		fabric = gw
		logger.Info("FABRIC_GATEWAY_INITIALIZED")
	})
	return fabric
}

// Close the connection and all associated resources
func Close() {
	if fabric == nil {
		return
	}
	logger.Info("CLOSE_FABRIC_GATEWAY")
	fabric.Close() // future use ( Close() is not implemented yet. )
}
