package main

import (
    "log"

    "github.com/boltdb/bolt"
)

func main() {
    db, err := bolt.Open("blog.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }

db.Update(func(tx *bolt.Tx) error {
    b, err := tx.CreateBucketIfNotExists([]byte("posts"))
    if err != nil {
        return err
    }
    return b.Put([]byte("2015-01-01"), []byte("My New Year post"))
}) 


    defer db.Close()

    // ...
}
