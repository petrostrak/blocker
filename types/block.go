package types

import (
	"crypto/sha256"

	"github.com/petrostrak/blocker/proto"
	pb "google.golang.org/protobuf/proto"
)

// HashBlock returns a SHA256 of the Header
func HashBlock(block *proto.Block) []byte {
	b, err := pb.Marshal(block)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(b)

	return hash[:]
}
