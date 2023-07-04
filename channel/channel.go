package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// forLoop()
	signalNotify()
}
// hahahahahahahhahahahahahhahhahahahahahahhahahahahahhahhahahahahahahhahahahahahhahhahahahahahahhahahahahahhahhahahahahahahhahahahahahhahhahahahahahahhahahahahahhahhahahahahahahhahahahahahhahhahahahahahahhahahahahahhahhahahahahahahhahahahahahhahhahahahahahahhahahahahahhah
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
