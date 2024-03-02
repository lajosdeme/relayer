package core

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lajosdeme/transaction-relayer/config"
	"github.com/lajosdeme/transaction-relayer/contracts"
)

const chainId = 4201

func HandleSubscriptionEvents() {
	address := common.HexToAddress(config.Get().SubscriptionContractAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}

	c, err := client()
	if err != nil {
		panic(err)
	}

	logs := make(chan types.Log)

	sub, err := c.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		panic(err)
	}

	fmt.Println("subscribed to blockchain events")

	for {
		select {
		case err := <-sub.Err():
			fmt.Println("failed to parse event: ", err)
		case vLog := <-logs:
			go handleEvent(vLog)
		}
	}
}

func client() (*ethclient.Client, error) {
	client, err := ethclient.Dial(config.Get().NodeUrl)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func BlockchainAuth() (*bind.TransactOpts, error) {
	privKey, err := crypto.HexToECDSA(config.Get().Key)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(chainId))
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func SubscriptionContract(client *ethclient.Client) (*contracts.Subscription, error) {
	sub, err := contracts.NewSubscription(common.HexToAddress(config.Get().SubscriptionContractAddress), client)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func handleEvent(l types.Log) {
	newWorkEventSig := crypto.Keccak256Hash([]byte("NewSubscription(bytes32,uint256)")).Hex()
	eventSig := l.Topics[0].Hex()

	switch eventSig {
	case newWorkEventSig:
		var id [32]byte
		copy(id[:], l.Topics[1].Bytes()[:32])

		client, err := client()
		if err != nil {
			fmt.Println(err)
			return
		}
		contract, err := SubscriptionContract(client)
		if err != nil {
			fmt.Println(err)
			return
		}

		userId, err := contract.UserIds(nil, id)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("user id: ", userId)

		expiry := new(big.Int).SetBytes(l.Topics[2].Bytes())

		fmt.Println("Expiry:", expiry.Int64())

		if err := DB().SubscribeUser(userId, int(expiry.Int64())); err != nil {
			fmt.Println(err)
			return
		}
	}
}
