package types

import (
	"crypto/sha256"

	pb "github.com/golang/protobuf/proto"
	"github.com/ssssunat/blocker/crypto"
	"github.com/ssssunat/blocker/proto"
)

// Hash block returns SHA 256 of the header
func HashBlock(block *proto.Block) []byte {
	b, err := pb.Marshal(block)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(b)
	return hash[:]
}

func SignBlock(pk *crypto.PrivateKey, b *proto.Block) *crypto.Signature {
	return pk.Sign(HashBlock(b))
}

