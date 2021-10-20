package main

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db, err := leveldb.OpenFile("./debug/db", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Put([]byte("foo"), []byte("bar"), nil); err != nil {
		log.Fatal(err)
	}
	ds, err := db.Get([]byte("foo"), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", ds)
	it := db.NewIterator(nil, nil)
	for it.Next() {
		log.Printf("[key] %s, [val] %s", it.Key(), it.Value())
	}
	prop, err := db.GetProperty("leveldb.stats")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(prop)
	if err = db.Delete([]byte("foo"), nil); err != nil {
		log.Fatal(err)
	}
	ds, err = db.Get([]byte("foo"), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", ds)
	prop, err = db.GetProperty("leveldb.stats")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(prop)
}
