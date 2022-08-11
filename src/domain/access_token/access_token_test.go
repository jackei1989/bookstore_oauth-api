package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessToken_Constants(t *testing.T) {
	assert.EqualValues(t, 24, EXPIRATIONTIME, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined access token id")
	assert.Truef(t, at.UserId == 0, "new access token should not have associated user id")
}

func TestAccessToken_IsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiring three hours from new should NOT be expired")
}
