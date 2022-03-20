package application

func (app *Application) PutAccessLogging(requestId, tokenId, tokenName, ip, method, path, query, errorMessage string, status int) error {
	return app.MySQLFetcher.PutAccessLogging(requestId, tokenId, tokenName, ip, method, path, query, errorMessage, status)
}
