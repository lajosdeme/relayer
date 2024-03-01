package core

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lajosdeme/transaction-relayer/config"
	"github.com/lajosdeme/transaction-relayer/contracts"
	"github.com/lajosdeme/transaction-relayer/types"
)

func ExecuteRelayCall(execPayload types.ExecutePayload, remainingQuota int) (string, int, error) {
	client, err := ethclient.Dial(config.Get().NodeUrl)
	if err != nil {
		return "", 0, err
	}

	universalProfile, err := contracts.NewUniversalProfile(common.HexToAddress(execPayload.Address), client)
	if err != nil {
		return "", 0, err
	}

	keyManagerAddr, err := universalProfile.Owner(nil)
	fmt.Println("key manager address: ", keyManagerAddr)
	if err != nil {
		return "", 0, err
	}

	keyManager, err := contracts.NewLSP6(keyManagerAddr, client)
	if err != nil {
		return "", 0, err
	}

	auth, err := getAuth()
	if err != nil {
		return "", 0, err
	}

	fmt.Println(execPayload.Tx)

	sig, err := hexutil.Decode(execPayload.Tx.Signature)
	if err != nil {
		return "", 0, err
	}

	abiPayload, err := hexutil.Decode(execPayload.Tx.Abi)
	if err != nil {
		return "", 0, err
	}

	var nonce big.Int
	nonce.SetString(execPayload.Tx.Nonce, 10)

	// this timestamp means valid until 1 Jan 2025
	validityTstamp := big.NewInt(1735689600)

	estimatedGas, err := estimateGas(&keyManagerAddr, sig, &nonce, validityTstamp, abiPayload)
	if err != nil {
		return "", 0, err
	}

	if estimatedGas > uint64(remainingQuota) {
		return "", 0, errors.New("not enough quota")
	}

	tx, err := keyManager.ExecuteRelayCall(auth, sig, &nonce, validityTstamp, abiPayload)
	if err != nil {
		fmt.Println("transaction error: ", err)
		return "", 0, err
	}

	fmt.Println("Relay transaction executed: ", tx.Hash().String())
	fmt.Println("GAS USED: ", tx.Gas())

	return tx.Hash().String(), int(tx.Gas()), nil
}

func getAuth() (*bind.TransactOpts, error) {
	privKey, err := crypto.HexToECDSA(config.Get().Key)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(int64(4201)))
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func estimateGas(address *common.Address, signature []byte, nonce *big.Int, validityTimestamps *big.Int, payload []byte) (uint64, error) {

	contractABI, err := abi.JSON(strings.NewReader(contracts.LSP6ABI))
	if err != nil {
		return 0, err
	}

	encodedData, err := contractABI.Pack("executeRelayCall", signature, nonce, validityTimestamps, payload)
	if err != nil {
		return 0, err
	}

	client, err := client()
	if err != nil {
		return 0, err
	}
	msg := ethereum.CallMsg{
		To:   address,
		Data: encodedData,
	}

	gas, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return 0, err
	}
	return gas, nil
}
