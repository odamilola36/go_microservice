package security

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestNewToken(t *testing.T) {
	id := primitive.NewObjectID()
	token, err := NewToken(id.Hex())
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestNewTokenPayload(t *testing.T) {
	id := primitive.NewObjectID()
	token, err := NewToken(id.Hex())
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	payload, err := NewTokenPayload(token)
	assert.NoError(t, err)
	assert.Equal(t, payload.UserId, id.Hex())
}
