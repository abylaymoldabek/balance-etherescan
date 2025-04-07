package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/abylaymoldabek/balance-etherescan/cmd/config"
	"github.com/abylaymoldabek/balance-etherescan/pkg/ethclient"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Ethereum address is required as the first argument")
	}

	address := os.Args[1]

	config.LoadConfig()

	infuraProjectID := config.GetInfuraProjectID()

	// Подключение к Ethereum
	client, err := ethclient.ConnectToEthereumClient(infuraProjectID)
	if err != nil {
		log.Fatalf("Error connecting to Ethereum: %v", err)
	}

	// Получение баланса ETH
	ethBalance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		log.Fatalf("Failed to retrieve ETH balance: %v", err)
	}

	// Получение баланса WETH
	wethAddress := common.HexToAddress(config.GetWETHContractAddress())
	wethBalance, err := ethclient.GetERC20Balance(client, wethAddress, common.HexToAddress(address))
	if err != nil {
		log.Fatalf("Failed to retrieve WETH balance: %v", err)
	}

	// Получение цены ETH в USD через Chainlink
	chainlinkAddress := common.HexToAddress(config.GetChainlinkETHUSDAddress())
	ethUSDPrice, err := ethclient.GetChainlinkPrice(client, chainlinkAddress)
	if err != nil {
		log.Fatalf("Failed to get ETH/USD price: %v", err)
	}

	// Пересчет баланса в USD
	ethBalanceUSD := new(big.Float).Mul(new(big.Float).SetInt(ethBalance), new(big.Float).SetInt(ethUSDPrice))
	wethBalanceUSD := new(big.Float).Mul(new(big.Float).SetInt(wethBalance), new(big.Float).SetInt(ethUSDPrice))

	// Вывод результата
	fmt.Printf("ETH balance: %s\n", ethBalance.String())
	fmt.Printf("ETH balance in USD: %s\n", ethBalanceUSD.String())
	fmt.Printf("WETH balance: %s\n", wethBalance.String())
	fmt.Printf("WETH balance in USD: %s\n", wethBalanceUSD.String())
}
