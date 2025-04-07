package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	InfuraProjectID        string
	WETHContractAddress    string
	ChainlinkETHUSDAddress string
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetInfuraProjectID() string {
	return os.Getenv("INFURA_PROJECT_ID")
}

func GetWETHContractAddress() string {
	return os.Getenv("WETH_CONTRACT_ADDRESS")
}

func GetChainlinkETHUSDAddress() string {
	return os.Getenv("CHAINLINK_ETH_USD_ADDRESS")
}
