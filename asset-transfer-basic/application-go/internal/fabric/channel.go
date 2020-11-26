package fabric

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"github.com/nkmr-jp/go-logger-scaffold/logger"
)

type Channel struct {
	channelID string
	network   *gateway.Network
}

func NewChannel(channelID string) (*Channel, error) {
	network, err := fabric.GetNetwork(channelID)
	if err != nil {
		return nil, err
	}
	return &Channel{
		channelID: channelID,
		network:   network,
	}, nil
}

func (c *Channel) GetContract(chaincodeID string) *gateway.Contract {
	logger.Info("GET_CONTRACT")
	return c.network.GetContract(chaincodeID)
}
