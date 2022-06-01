package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// forLoop()
	// signalNotify()
	closeChannel()
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

func signalNotify() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	log.Println("catch signal", sig)
}

func closeChannel() {
	ch := make(chan struct{ name string })
	ctx := context.Background()

	go func() {
		log.Print("input data")
		ch <- struct{ name string }{name: "c"}
		log.Print("try to close")
		close(ch)
		log.Print("closed")
		<-ctx.Done()
		log.Print("over")
	}()

	time.Sleep(time.Second)
	for {
		select {
		case <-ctx.Done():
			log.Print("context done")
		case c, ok := <-ch:
			log.Printf("%v, %v", c, ok)
			if !ok {
				log.Print("channel has been closed")
				return
			}
		}
	}
}
