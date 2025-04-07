package ethclient

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetERC20Balance(client *ethclient.Client, tokenAddress common.Address, userAddress common.Address) (*big.Int, error) {
	erc20ABI := `[{"constant":true,"inputs":[{"name":"owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`
	parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		return nil, err
	}

	callData, err := parsedABI.Pack("balanceOf", userAddress)
	if err != nil {
		return nil, err
	}

	msg := ethereum.CallMsg{
		To:   &tokenAddress,
		Data: callData,
	}

	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, err
	}

	var balance *big.Int
	err = parsedABI.UnpackIntoInterface(&balance, "balanceOf", result)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func GetChainlinkPrice(client *ethclient.Client, contractAddress common.Address) (*big.Int, error) {
	chainlinkETHUSDABI := `[{"constant":true,"inputs":[],"name":"latestAnswer","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"view","type":"function"}]`
	parsedABI, err := abi.JSON(strings.NewReader(chainlinkETHUSDABI))
	if err != nil {
		return nil, err
	}

	callData, err := parsedABI.Pack("latestAnswer")
	if err != nil {
		return nil, err
	}

	msg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: callData,
	}

	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, err
	}

	var price *big.Int
	err = parsedABI.UnpackIntoInterface(&price, "latestAnswer", result)
	if err != nil {
		return nil, err
	}

	return price, nil
}
