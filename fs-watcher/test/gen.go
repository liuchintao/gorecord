package main

import (
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	writeIOUtil()
}

func writeWithTempFile() {
	f, err := ioutil.TempFile("./", "gen-go-")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for i := range [100]byte{} {
		f.Write([]byte(strconv.Itoa(i) + "\n"))
	}
}

func writeIOUtil() {
	if err := ioutil.WriteFile("ioutil.test", []byte("hello"), 0o640); err != nil {
		log.Fatal(err)
	}
}
