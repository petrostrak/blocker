package types

import (
	"fmt"
	"testing"

	"github.com/petrostrak/blocker/crypto"
	"github.com/petrostrak/blocker/proto"
	"github.com/petrostrak/blocker/util"
)

// My balance 100 coins
// I want to send 5 coins to address "AAA"
//
// I need to make an input, to specify the output
// of the previous transaction where we received
// coins.

func TestNewTransaction(t *testing.T) {
	fromPrivKey := crypto.GeneratePrivateKey()
	fromAddress := fromPrivKey.Public().Address().Bytes()

	toPrivKey := crypto.GeneratePrivateKey()
	toAddr := toPrivKey.Public().Address().Bytes()

	input := &proto.TxInput{
		PrevHash:     util.RandomHash(),
		PrevOutIndex: 0,
		PublicKey:    fromPrivKey.Public().Bytes(),
	}

	output1 := &proto.TxOutput{
		Amount:  5,
		Address: toAddr,
	}

	output2 := &proto.TxOutput{
		Amount:  95,
		Address: fromAddress,
	}

	tx := &proto.Transcation{
		Version: 1,
		Inputs:  []*proto.TxInput{input},
		Outputs: []*proto.TxOutput{output1, output2},
	}

	sig := SignTransaction(fromPrivKey, tx)
	input.Signarure = sig.Bytes()

	fmt.Printf("%+v\n", tx)
}
