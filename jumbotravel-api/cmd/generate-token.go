package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pomaretta/jumbotravel/jumbotravel-api/lib/rsajwt"
)

type endpoints []string

func (e *endpoints) String() string {
	return fmt.Sprintf("%v#", *e)
}

func (e *endpoints) Set(value string) error {
	*e = append(*e, value)
	return nil
}

func main() {
	var tokenName, tokenType string
	var endpoints endpoints

	flag.StringVar(&tokenName, "tokenName", "", "Token name")
	flag.StringVar(&tokenType, "tokenType", "", "Token type (api, user)")

	flag.Var(&endpoints, "endpoint", `Endpoint to generate the token for.
Structure: [http method]/[endpoint]`)

	flag.Parse()

	if tokenName == "" {
		log.Fatal(" -tokenName is required")
	}

	if tokenType != "api" && tokenType != "user" {
		log.Fatal(" -tokenType is required and should be api or user")
	}

	var err error

	// Get private key from file and return a Reader
	privateKey, err := os.Open("rsa.private")
	if err != nil {
		log.Fatal(err)
	}

	//signer, err := rsajwt.ShorterSignerFromPrivateKey("rsa.private")
	signer, err := rsajwt.ShorterSignerFromPrivateKey(privateKey)
	if err != nil {
		log.Fatalf("Cannot create signer: %s", err)
	}

	allowPolicy := rsajwt.AllowPolicy{
		Resources: endpoints,
	}

	signed, err := signer.Sign(allowPolicy, tokenName, tokenType)
	if err != nil {
		log.Fatalf("Cannot sign message: %s", err)
	}

	fmt.Println(signed.Token)
}
