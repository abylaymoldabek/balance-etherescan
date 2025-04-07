package ethclient

import (
	"fmt"
	"log"
	"net/url"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// EthereumMainnetURL is the URL for the Ethereum mainnet
	EthereumMainnetURL = "https://mainnet.infura.io/v3/"
)

func ConnectToEthereumClient(infuraProjectID string) (*ethclient.Client, error) {
	url, err := url.JoinPath(EthereumMainnetURL, infuraProjectID)
	if err != nil {
		return nil, fmt.Errorf("failed join path with url %v: %w", EthereumMainnetURL, err)
	}
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		return nil, err
	}
	return client, nil
}
