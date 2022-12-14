package access_token

import (
	"github.com/jackei1989/bookstore_oauth-api/src/utils/errors"
	"strings"
	"time"
)

const EXPIRATIONTIME = 24

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id!")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id!")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid clinet id!")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time!")
	}
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(EXPIRATIONTIME * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
