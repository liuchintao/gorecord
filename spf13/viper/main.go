package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	path := "./debug/config.yaml"
	if envPath := os.Getenv("CONFIG_FILE"); envPath != "" {
		path = envPath
	}
	viper.SetConfigFile(path)
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Print("config has been changed")
		for _, key := range viper.AllKeys() {
			log.Printf("[%s]: %v", key, viper.Get(key))
		}
	})
	viper.WatchConfig()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	log.Print("exit with ", sig.String())
}
