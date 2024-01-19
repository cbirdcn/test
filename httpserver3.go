package main

import (
    "fmt"
    "net/http"
    "github.com/boltdb/bolt"
    "log"
    "strconv"
)

var db *bolt.DB

func aHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "a\n")
    db.Update(func(tx *bolt.Tx) error {
    b, err := tx.CreateBucketIfNotExists([]byte("posts"))
    if err != nil {
        return err
    }
    return b.Put([]byte("2015-01-01"), []byte("My New Year post"))
})
}

func bHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "b\n")
    db.Update(func(tx *bolt.Tx) error {
    b, err := tx.CreateBucketIfNotExists([]byte("posts"))
    if err != nil {
        return err
    }
    return b.Put([]byte("2016-01-01"), []byte("My New Year post 2016"))
})
}

func cHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "c\n")
    for i:=2017;i<=2030;i++ {
	go hand(i)
    }
}

func hand(i int){
    db.Update(func(tx *bolt.Tx) error {
    b, err := tx.CreateBucketIfNotExists([]byte("posts"))
    if err != nil {
        return err
    }
    return b.Put([]byte(strconv.Itoa(i)), []byte("My New Year post"))
})

}

func main() {
    var err error
    db, err = bolt.Open("blog.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/a", aHandler)
    http.HandleFunc("/b", bHandler)
    http.HandleFunc("/c", cHandler)
    http.ListenAndServe(":8000", nil)

    defer db.Close()
}
