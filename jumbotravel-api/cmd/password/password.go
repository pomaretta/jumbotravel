package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"log"
)

func main() {

	var password string

	flag.StringVar(&password, "password", "", "Password")
	flag.Parse()

	if password == "" {
		log.Fatal(" -password is required")
	}

	hash := md5.Sum([]byte(password))
	fmt.Printf("%x", hash)
}
