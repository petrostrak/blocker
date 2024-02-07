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
