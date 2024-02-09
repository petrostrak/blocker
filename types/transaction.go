package types

import (
	"crypto/sha256"

	"github.com/petrostrak/blocker/crypto"
	"github.com/petrostrak/blocker/proto"

	pb "google.golang.org/protobuf/proto"
)

func SignTransaction(pk *crypto.PrivateKey, tx *proto.Transcation) *crypto.Signature {
	return pk.Sign(HashTransaction(tx))
}

func HashTransaction(tx *proto.Transcation) []byte {
	b, err := pb.Marshal(tx)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(b)
	return hash[:]
}

func VerifyTransaction(tx *proto.Transcation) bool {
	for _, input := range tx.Inputs {
		sig := crypto.SignatureFromBytes(input.Signarure)
		pubKey := crypto.PublicKeyFromBytes(input.PublicKey)
		if !sig.Verify(pubKey, HashTransaction(tx)) {
			return false
		}
	}
	return true
}
