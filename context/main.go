package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	go func() {
		defer log.Println("hello I'm sub")
		<-ctx.Done()
	}()
	log.Print("start sleeping")
	<-time.After(time.Second)
}
