package types

import (
	"fmt"
	"testing"

	"github.com/ssssunat/blocker/crypto"
	"github.com/ssssunat/blocker/proto"
	"github.com/ssssunat/blocker/util"
	"github.com/stretchr/testify/assert"
)

// my baalance 100 coins

// want to send 5 coints to "AAA"

// 2outputs
// 5 	to the dude we wanna send
// 95 back to our address

func TestNewTransaction(t *testing.T) {
	fromPrivKey := crypto.GeneratePrivateKey()
	fromAddress := fromPrivKey.Public().Address().Bytes()

	toPrivKey := crypto.GeneratePrivateKey()
	toAddress := toPrivKey.Public().Address().Bytes()

	in := &proto.TxInput{
		PrevTxHash:   util.RandomHash(),
		PrevOutIndex: 0,
		PublicKey:    fromPrivKey.Public().Bytes(),
	}

	output1 := &proto.TxOutput{
		Amount:  5,
		Address: toAddress,
	}
	output2 := &proto.TxOutput{
		Amount:  95,
		Address: fromAddress,
	}

	tx := &proto.Transaction{
		Version: 1,
		Inputs:  []*proto.TxInput{in},
		Output:  []*proto.TxOutput{output1, output2},
	}
	sig := SignTransaction(fromPrivKey, tx)
	in.Signature = sig.Bytes()

	assert.True(t, VerifyTransaction(tx))

	fmt.Printf("%+v\n", tx)
}
