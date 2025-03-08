package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()

	assert.Equal(t, len(privKey.Bytes()), privKeyLen)

	pubKey := privKey.Public()
	assert.Equal(t, len(pubKey.Bytes()), pubKeyLen)

}

func TestNewPrivateKeyFromString(t *testing.T) {
	var (
		seed       = "6326479873f027c234f3b110ad869b7f4aceca2392c99f242650d0a9bf7a628c"
		addressStr = "1c0db28d463b11bcba77d44617fc05653b61a78f"
		privKey = NewPrivateKeyFromString(seed)
	)

	
	assert.Equal(t, privKeyLen, len(privKey.Bytes()))
	address := privKey.Public().Address()
	fmt.Println(address)
	assert.Equal(t, address.String(), addressStr)
}

func TestPrivKeySign(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()

	msg := []byte("foo bar bazz")

	sig := privKey.Sign(msg)
	assert.True(t, sig.Verify(pubKey, msg))
	assert.False(t, sig.Verify(pubKey, []byte("foo")))

	// Test with invalid pubKey
	fakePriveKey := GeneratePrivateKey()
	fakePubKey := fakePriveKey.Public()
	assert.False(t, sig.Verify(fakePubKey, msg))
}

func TestPublicKeyToAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	address := pubKey.Address()
	assert.Equal(t, addressLen, len(address.Bytes()))
	fmt.Println(address)
}
