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
	iterAndModify(db)
}

func easy(db *leveldb.DB) {
	if err := db.Put([]byte("foo"), []byte("bar"), nil); err != nil {
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

func iterAndModify(db *leveldb.DB) {
	data := map[string]string{
		"hello": "world",
		"foo":   "bar",
		"fo0":   "bax",
	}
	for k, v := range data {
		if err := db.Put([]byte(k), []byte(v), nil); err != nil {
			log.Fatal(err)
		}
	}
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		log.Printf("[%s]: %s", iter.Key(), iter.Value())
		if err := db.Delete(iter.Key(), nil); err != nil {
			log.Fatal(err)
		}
	}
	log.Print("Iter again")
	iter = db.NewIterator(nil, nil)
	for iter.Next() {
		log.Printf("[%s]: %s", iter.Key(), iter.Value())
	}
}
