package db

import (
	"github.com/gocql/gocql"
	"github.com/jackei1989/bookstore_oauth-api/src/clients/cassandra"
	"github.com/jackei1989/bookstore_oauth-api/src/domain/access_token"
	"github.com/jackei1989/bookstore_oauth-api/src/utils/errors"
)

const (
	QUERYGETACCESSTOKEN    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?"
	QUERYCREATEACCESSTOKEN = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES(?,?,?,?)"
	QUERYUPDATEEXPIRES     = "UPDATE access_tokens SET expires=? WHERE access_token=?"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(token access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(QUERYGETACCESSTOKEN, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found with given id")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}

func (r *dbRepository) Create(token access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(QUERYCREATEACCESSTOKEN,
		token.AccessToken,
		token.UserId, token.ClientId,
		token.Expires,
	).Exec(); err != nil {
		return errors.NewBadRequestError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(QUERYUPDATEEXPIRES,
		token.Expires,
		token.AccessToken,
	).Exec(); err != nil {
		return errors.NewBadRequestError(err.Error())
	}
	return nil
}
