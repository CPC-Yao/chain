// Copyright 2018 The cpchain authors
// This file is part of the cpchain library.
//
// The cpchain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The cpchain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the cpchain library. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"bitbucket.org/cpchain/chain/accounts/abi/bind"
	"bitbucket.org/cpchain/chain/accounts/keystore"
	"bitbucket.org/cpchain/chain/api/cpclient"
	"bitbucket.org/cpchain/chain/commons/log"
	dpor "bitbucket.org/cpchain/chain/contracts/dpor/contracts"
	campaign "bitbucket.org/cpchain/chain/contracts/dpor/contracts/campaign"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type keystorePair struct {
	keystorePath string
	passphrase   string
}

var (
	endPoint  = "http://localhost:8501"
	dataDir   = "data/"
	keystores = []keystorePair{
		{
			"data1/keystore/",
			"password",
		},
		{
			"data2/keystore/",
			"password",
		},
		{
			"data3/keystore/",
			"pwdnode1",
		},
		{
			"data4/keystore/",
			"pwdnode2",
		},
		{
			"data5/keystore/",
			"password",
		},
		{
			"data6/keystore/",
			"password",
		},
		{
			"data7/keystore/",
			"password",
		},
		{
			"data8/keystore/",
			"password",
		},
		{
			"data9/keystore/",
			"password",
		},
		{
			"data10/keystore/",
			"password",
		},
	}
)

func getAccount(keyStoreFilePath string, passphrase string) (*ecdsa.PrivateKey, *ecdsa.PublicKey, common.Address, []byte) {
	// Load account.
	ff, err := filepath.Abs(".")
	file, err := os.Open(ff + "/" + dataDir + keyStoreFilePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	keyPath, err := filepath.Abs(filepath.Dir(file.Name()))
	if err != nil {
		log.Fatal(err.Error())
	}

	kst := keystore.NewKeyStore(keyPath, 2, 1)

	// Get account.
	account := kst.Accounts()[0]
	account, key, err := kst.GetDecryptedKey(account, passphrase)
	if err != nil {
		log.Fatal(err.Error())
	}

	privateKey := key.PrivateKey
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return privateKey, publicKeyECDSA, fromAddress, key.RsaKey.PublicKey.RsaPublicKeyBytes
}

func claimCampaign(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, address common.Address, contractAddress common.Address) {
	// Create client.
	client, err := cpclient.Dial(endPoint)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("from address:", address.Hex()) // 0xe94b7b6c5a0e526a4d97f9768ad6097bde25c62a

	// Check balance.
	bal, err := client.BalanceAt(context.Background(), address, nil)
	fmt.Println("balance:", bal)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("gasPrice:", gasPrice)

	startTime := time.Now()
	fmt.Printf("transaction start at: %s\n", time.Now())

	ctx := context.Background()

	instance, err := campaign.NewCampaign(contractAddress, client)

	baseDeposit := 50
	gasLimit := 3000000
	numOfCampaign := 10
	cpuNonce := 345678
	cpuBlockNum := 100
	memNonce := 345678
	memBlockNum := 101

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Value = big.NewInt(int64(baseDeposit * numOfCampaign))
	auth.GasLimit = uint64(gasLimit)
	auth.GasPrice = gasPrice

	tx, err := instance.ClaimCampaign(auth, big.NewInt(int64(numOfCampaign)), uint64(cpuNonce), big.NewInt(int64(cpuBlockNum)),
		uint64(memNonce), big.NewInt(int64(memBlockNum)))
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("transaction hash: ", tx.Hash().Hex())

	startTime = time.Now()
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		log.Fatalf("failed to deploy contact when mining :%v", err)
	}

	fmt.Printf("tx mining take time:%s\n", time.Since(startTime))
	fmt.Println("receipt.Status:", receipt.Status)

	noc, deposit, startView, stopView, err := instance.CandidateInfoOf(nil, address)
	fmt.Println("candidate info of", address.Hex(), ":", noc, deposit, startView, stopView)
}

func claimSigner(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, address common.Address, contractAddress common.Address, rsaPubkey []byte) {
	// Create client.
	client, err := cpclient.Dial(endPoint)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("from address:", address.Hex()) // 0xe94b7b6c5a0e526a4d97f9768ad6097bde25c62a

	// Check balance.
	bal, err := client.BalanceAt(context.Background(), address, nil)
	fmt.Println("balance:", bal)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("gasPrice:", gasPrice)

	startTime := time.Now()
	fmt.Printf("transaction start at: %s\n", time.Now())

	ctx := context.Background()

	instance, err := dpor.NewProposerRegister(contractAddress, client)

	gasLimit := 3000000

	auth := bind.NewKeyedTransactor(privateKey)
	auth.GasLimit = uint64(gasLimit)
	auth.GasPrice = gasPrice
	auth.Value = big.NewInt(500)

	fmt.Println("rsaPubkey", hex.Dump(rsaPubkey))

	tx, err := instance.RegisterPublicKey(auth, rsaPubkey)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("transaction hash: ", tx.Hash().Hex())

	startTime = time.Now()
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		log.Fatalf("failed to deploy contact when mining :%v", err)
	}

	fmt.Printf("tx mining take time:%s\n", time.Since(startTime))
	fmt.Println("receipt.Status:", receipt.Status)

	pubkey, err := instance.GetPublicKey(nil, address)
	fmt.Println("pubkey of", address.Hex(), ":", pubkey)
}

func main() {

	campaignAddress := common.HexToAddress("0x1a9fAE75908752d0ABf4DCa45ebcaC311C376290")
	signerAddress := common.HexToAddress("0x4CE687F9dDd42F26ad580f435acD0dE39e8f9c9C")

	for i, kPair := range keystores {
		fmt.Println(i)
		keystoreFile, passphrase := kPair.keystorePath, kPair.passphrase
		privKey, pubKey, addr, rsaPubKey := getAccount(keystoreFile, passphrase)
		claimCampaign(privKey, pubKey, addr, campaignAddress)
		claimSigner(privKey, pubKey, addr, signerAddress, rsaPubKey)
	}
}
