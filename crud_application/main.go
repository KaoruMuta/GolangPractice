package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

// reference: https://qiita.com/Syoitu/items/8e7e3215fb7ac9dabc3a
func main() {
	mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, q *http.Request) {
        message := map[string]string {
            "message": "Hello World",
        }
        jsonMessage, err := json.Marshal(message)
        if err != nil {
            panic(err.Error())
        }
        w.Write(jsonMessage)
    })
	http.ListenAndServe("127.0.0.1:8080", mux)
	fmt.Println("Hello World!")
}