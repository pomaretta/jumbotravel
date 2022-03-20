package adminbuilders

import (
	"errors"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type PutAccessLoggingQueryBuilder struct {
	builders.MySQLQueryBuilder

	requestId    string
	tokenId      string
	tokenName    string
	ip           string
	method       string
	path         string
	query        string
	errorMessage string
	status       int
}

func (qb *PutAccessLoggingQueryBuilder) SetRequestId(requestId string) {
	qb.requestId = requestId
}

func (qb *PutAccessLoggingQueryBuilder) SetTokenId(tokenId string) {
	qb.tokenId = tokenId
}

func (qb *PutAccessLoggingQueryBuilder) SetTokenName(tokenName string) {
	qb.tokenName = tokenName
}

func (qb *PutAccessLoggingQueryBuilder) SetIp(ip string) {
	qb.ip = ip
}

func (qb *PutAccessLoggingQueryBuilder) SetMethod(method string) {
	qb.method = method
}

func (qb *PutAccessLoggingQueryBuilder) SetPath(path string) {
	qb.path = path
}

func (qb *PutAccessLoggingQueryBuilder) SetQuery(query string) {
	qb.query = query
}

func (qb *PutAccessLoggingQueryBuilder) SetErrorMessage(errorMessage string) {
	qb.errorMessage = errorMessage
}

func (qb *PutAccessLoggingQueryBuilder) SetStatus(status int) {
	qb.status = status
}

func (qb *PutAccessLoggingQueryBuilder) createArguments() ([]interface{}, error) {

	args := []interface{}{}

	if qb.requestId == "" {
		return nil, errors.New("requestId is required")
	}
	args = append(args, qb.requestId)

	if qb.tokenId == "" {
		return nil, errors.New("tokenId is required")
	}
	args = append(args, qb.tokenId)
	args = append(args, qb.tokenName)

	if qb.ip == "" {
		return nil, errors.New("ip is required")
	}
	args = append(args, qb.ip)

	if qb.method == "" {
		return nil, errors.New("method is required")
	}
	args = append(args, qb.method)

	if qb.path == "" {
		return nil, errors.New("path is required")
	}
	args = append(args, qb.path)
	args = append(args, qb.query)
	args = append(args, qb.errorMessage)

	if qb.status == 0 {
		return nil, errors.New("status is required")
	}
	args = append(args, qb.status)

	return args, nil
}

func (qb *PutAccessLoggingQueryBuilder) BuildQuery() (string, []interface{}, error) {

	args, err := qb.createArguments()
	if err != nil {
		return "", nil, err
	}

	partialQuery := `
		INSERT INTO api_access (
			requestid,
			token_id,
			token_name,
			ip,
			method,
			path,
			query,
			error_message,
			status
		) VALUES (
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?
		);
	`

	return partialQuery, args, nil
}
