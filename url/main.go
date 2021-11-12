package main

import (
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("http://paas.127-0-0-1.nip.io:30080/api/edge")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(u.Host, u.Path)
}
