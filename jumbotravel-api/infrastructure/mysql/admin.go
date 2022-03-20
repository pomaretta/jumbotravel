package mysql

import "github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders/adminbuilders"

func (db *MySQL) PutAccessLogging(requestId, tokenId, tokenName, ip, method, path, query, errorMessage string, status int) error {

	qb := &adminbuilders.PutAccessLoggingQueryBuilder{}
	qb.SetRequestId(requestId)
	qb.SetTokenId(tokenId)
	qb.SetTokenName(tokenName)
	qb.SetIp(ip)
	qb.SetMethod(method)
	qb.SetPath(path)
	qb.SetQuery(query)
	qb.SetErrorMessage(errorMessage)
	qb.SetStatus(status)

	_, err := db.Put(qb)
	if err != nil {
		return err
	}

	return nil
}
