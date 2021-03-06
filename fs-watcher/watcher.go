package main

import (
	"context"
	"log"

	"github.com/syncthing/notify"
)

func main() {
	watch(context.Background())
}

func watch(ctx context.Context) {
	c := make(chan notify.EventInfo, 1024)
	if err := notify.Watch("./test/...", c, notify.Write); err != nil {
		log.Fatalln(err)
	}
	if err := notify.Watch(".", c, notify.Write); err != nil {
		log.Fatalln(err)
	}
	defer notify.Stop(c)
	for {
		e := <-c
		log.Println(e.Event().String())
	}
}
