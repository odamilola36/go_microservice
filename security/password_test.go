package security

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptPassword(t *testing.T) {
	password := "password"
	encryptedPassword, err := EncryptPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, encryptedPassword)
	i := len(encryptedPassword)
	assert.Equal(t, i, 60)
}

func TestVerifyPassword(t *testing.T) {
	password := "password"
	encryptedPassword, err := EncryptPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, encryptedPassword)

	isVerified := VerifyPassword(password, encryptedPassword)
	assert.True(t, isVerified)
	assert.NotEqual(t, password, encryptedPassword)

}
