package fabric

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"github.com/nkmr-jp/go-logger-scaffold/logger"
)

var (
	once    sync.Once
	fabric  *Fabric
	ccpPath = filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)
)

const (
	channelID   = "mychannel"
	chaincodeID = "basic"
	walletPath  = "wallet"
	walletLabel = "appUser"
	mspID       = "Org1MSP"
)

type Fabric struct {
	gateway  *gateway.Gateway
	network  *gateway.Network
	contract *gateway.Contract
}

func InitFabric() *Fabric {
	once.Do(func() {
		if err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true"); err != nil {
			logger.Fatalf("SET_ENV_ERROR", err)
		}
		gw, err := gateway.Connect(
			gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
			gateway.WithIdentity(createWallet(), "appUser"),
		)
		if err != nil {
			logger.Fatalf("CONNECT_TO_GATEWAY_ERROR", err)
		}
		network, err := gw.GetNetwork(channelID)
		if err != nil {
			logger.Fatalf("GET_NETWORK_ERROR", err)
		}

		fabric = &Fabric{
			gateway:  gw,
			network:  network,
			contract: network.GetContract(chaincodeID),
		}
		logger.Info("FABRIC_CLIENT_INITIALIZED")
	})
	return fabric
}

func Close() {
	logger.Info("CLOSE_GATEWAY")
	fabric.gateway.Close()
}

func GetContract() *gateway.Contract {
	return fabric.contract
}

func checkInit() {
	if fabric == nil {
		logger.Fatalf("", fmt.Errorf("the fabric is not initialized. InitFabric() must be called"))
	}
}
