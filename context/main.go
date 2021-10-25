package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-ctx.Done()
		log.Println("hello I'm sub")
		wg.Done()
	}()
	log.Print("start sleeping")
	<-time.After(time.Second)
	cancel()
	wg.Wait()
}
