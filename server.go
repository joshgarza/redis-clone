package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Payload struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method Not Allowed")
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed reading request body: %s", err)
		return
	}

	var payload Payload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed decoding JSON data: %s", err)
		return
	}

	fmt.Fprintf(w, "Received key: %s and value: %s", payload.Key, payload.Value)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method Not Allowed")
		return
	}

	fmt.Println("GET params were:", r.URL.Query())

	key := r.URL.Query().Get("key")
	// Fetch value for the given key from your cache (once you implement it)
	fmt.Fprintf(w, "Value for key: %s is Value", key)  // Using string interpolation to display the key.
}

func main() {
	http.HandleFunc("/cache/set/", setHandler)  // This will match both POST and GET, but inside the handlers, you've differentiated them.
	http.HandleFunc("/cache/get/", getHandler)  // This is okay since both handlers first check the HTTP method.
	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", nil)
}
