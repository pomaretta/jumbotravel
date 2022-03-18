# **JumboTravel API**

Provide access to the JumboTravel Databases.

## Run API

In order to run the API, you need to provide the following configurations:

API Database Configurations
```toml
[Database]
Host = ""
Port = ""
DatabaseName = ""
User = ""
Password = ""
Environment = "PROD"

[Database]
Host = ""
Port = ""
DatabaseName = ""
User = ""
Password = ""
Environment = "DEV"
```

Execute Pipeline (Docs, Build, Run)
```bash
make ENV=<DEV|PROD> CONF=<path-to-file>
```

## **Authorizer**

Generate Private and Public Keys
```bash
openssl genrsa -out rsa.private 2048
openssl rsa -in new_rsa.private -out rsa.public -pubout -outform PEM
```

Generate JWT Token
```bash
go run cmd/generate-token.go -tokenName $TOKEN_NAME -tokenType api -endpoint "*/master/*"
```
