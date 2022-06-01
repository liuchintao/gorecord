package main

import (
	"encoding/json"
	"fmt"
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

	test := func(x ...*s) {
		fmt.Println(x)
		for _, val := range x {
			fmt.Print(*val)
		}
	}

	fmt.Println("dodododo")
	test(nil)
	test(&s{}, &s{M: map[string]interface{}{"aa": ""}})
}
