package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	forLoop()
}

func forLoop() {
	c := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for s := range c {
			log.Println("get", s, "[len]", len(c), "[cap]", cap(c))
		}
	}()
	go func() {
		for _, s := range []string{"a", "b", "c", "d", "e"} {
			c <- s
			time.Sleep(time.Second)
		}
		wg.Done()
	}()
	wg.Wait()
}
