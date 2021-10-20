package main

import (
	"encoding/json"
	"log"
)

type s struct {
	M map[string]interface{} `json:"m,omitempty"`
}

func main() {
	s1 := s{
		M: map[string]interface{}{},
	}
	bs, _ := json.Marshal(s1)
	log.Println(string(bs))
	var s2 s
	json.Unmarshal(bs, &s2)
	log.Println(s2.M == nil)
}
