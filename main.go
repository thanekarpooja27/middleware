package main

import (
    "fmt"
    "net/http"

    "github.com/bradfitz/gomemcache/memcache"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        mc := memcache.New("127.0.0.1:11211") // Replace with your Memcache server address
        key := "mykey"
        item, err := mc.Get(key)
        if err == nil {
            fmt.Fprintf(w, "Value for key %s: %s", key, item.Value)
        } else if err == memcache.ErrCacheMiss {
            fmt.Fprintf(w, "Key %s not found in cache", key)
        } else {
            fmt.Fprintf(w, "Error: %v", err)
        }
    })

    http.ListenAndServe(":8086", nil)
}
