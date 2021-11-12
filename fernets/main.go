package main

import (
	"log"

	"github.com/fernet/fernet-go"
)

func main() {
	s := "HW9d1jaB3mHh30Of32e2djJeBXRXIx4FmK3OidRxCX4="
	f, err := fernet.DecodeKey(s)
	if err != nil {
		panic(err)
	}
	es := f.Encode()
	log.Println(s)
	log.Println(es)
}
