package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

type description struct {
	name string
	info string
}

func main() {
	syncMap := sync.Map{}

	wg := sync.WaitGroup{}
	wg.Add(4)

	for i := range []int{1, 2, 3, 4} {
		log.Print("working ", i)
		go func() {
			log.Print("storing ", i)
			syncMap.Store(strconv.Itoa(i), description{
				name: strconv.Itoa(i),
				info: fmt.Sprintf("This is %d", i),
			})
			wg.Done()
		}()
	}

	wg.Wait()

	syncMap.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %s, value: %#v\n", key, value)
		return true
	})
}
