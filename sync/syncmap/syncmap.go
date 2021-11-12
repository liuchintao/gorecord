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
		x := i
		go func() {
			log.Print("storing ", x)
			syncMap.Store(strconv.Itoa(x), description{
				name: strconv.Itoa(x),
				info: fmt.Sprintf("This is %d", x),
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
