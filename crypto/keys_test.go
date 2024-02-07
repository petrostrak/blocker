package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	assert.Equal(t, len(privKey.Bytes()), privKeyLen)

	pubKey := privKey.Public()
	assert.Equal(t, len(pubKey.Bytes()), pubKeyLen)
}

func TestPrivateKeySign(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	msg := []byte("a test message")

	sig := privKey.Sign(msg)
	assert.True(t, sig.Verify(pubKey, msg))

	assert.False(t, sig.Verify(pubKey, []byte("a wrong test message")))

	invalidPrivKey := GeneratePrivateKey()
	assert.False(t, sig.Verify(invalidPrivKey.Public(), msg))
}

func TestNewPrivateKeyFromString(t *testing.T) {
	var (
		seed       = "eb8dae323c78eba00bd2260246476e7e169fbcb60bfceb2a48f9b12029e3d3ac"
		privKey    = NewPrivateKeyFromString(seed)
		addressStr = "bfb51671644577add9f84b8501303fba90f9f0b7"
	)

	assert.Equal(t, privKeyLen, len(privKey.Bytes()))
	address := privKey.Public().Address()
	assert.Equal(t, addressStr, address.String())
}

func TestPublicKeyToAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	publicKey := privKey.Public()
	address := publicKey.Address()

	assert.Equal(t, addressLen, len(address.Bytes()))
}
