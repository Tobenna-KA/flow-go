// Package sctest implements a sample BPL contract and example testing
// code using the emulator test blockchain.
package sctest

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/dapperlabs/flow-go/pkg/types"
	"github.com/dapperlabs/flow-go/sdk/emulator"
	"github.com/stretchr/testify/assert"
)

const (
	greatTokenContractFile = "./contracts/great-token.bpl"
)

func readFile(path string) []byte {
	contents, _ := ioutil.ReadFile(path)
	return contents
}

// Taken from sdk/emulator/emulator_test.go
// TODO make this more reusable
func bytesToString(b []byte) string {
	if b == nil {
		return "nil"
	}
	return strings.Join(strings.Fields(fmt.Sprintf("%d", b)), ",")
}

// Creates a script that mints an NFT and put it into storage.
func generateMintScript(nftCodeAddr types.Address) []byte {
	template := `
		import GreatNFTMinter, GreatNFT from 0x%s

		fun main(acct: Account) {
			var minter = GreatNFTMinter()
			var nft = minter.maybeMint()

			acct.storage["my_nft"] = nft
		}`

	filledTemplate := fmt.Sprintf(template, nftCodeAddr.String())
	return []byte(filledTemplate)
}

func newEmulator() *emulator.EmulatedBlockchain {
	return emulator.NewEmulatedBlockchain(&emulator.EmulatedBlockchainOptions{
		RuntimeLogger: func(msg string) { fmt.Println(msg) },
	})
}

func TestDeployment(t *testing.T) {
	b := newEmulator()

	// Should be able to deploy a contract as a new account with no keys.
	nftCode := readFile(greatTokenContractFile)
	_ = nftCode
	_, err := b.CreateAccount(nil, nil)
	assert.Nil(t, err)
	b.CommitBlock()
}

func TestMinting(t *testing.T) {
	b := newEmulator()

	// First, deploy the contract
	nftCode := readFile(greatTokenContractFile)
	contractAddr, err := b.CreateAccount(nil, nftCode)
	assert.Nil(t, err)

	// Create a transaction that mints a GreatToken by calling the minter.
	mintScript := generateMintScript(contractAddr)
	fmt.Println(string(mintScript))

	tx := types.Transaction{
		Script:         mintScript,
		ComputeLimit:   10,
		PayerAccount:   b.RootAccountAddress(),
		ScriptAccounts: []types.Address{b.RootAccountAddress()},
	}

	tx.AddSignature(b.RootAccountAddress(), b.RootKey())
	err = b.SubmitTransaction(&tx)
	assert.Nil(t, err)
	b.CommitBlock()
}
