// Command app is the sample's HTTP app, the one T1 records and tests: each
// request calls the dependency through ../shop. It listens on CONVERT_APP
// (default 127.0.0.1:18096).
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"convertsample/shop"
)

func main() {
	addr := os.Getenv("CONVERT_APP")
	if addr == "" {
		addr = "127.0.0.1:18096"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/user/"))
		if err != nil {
			http.Error(w, "bad id", http.StatusBadRequest)
			return
		}
		name, err := shop.UserName(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]any{"id": id, "name": name})
	})
	mux.HandleFunc("/price/", func(w http.ResponseWriter, r *http.Request) {
		item := strings.TrimPrefix(r.URL.Path, "/price/")
		cents, err := shop.Price(item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]any{"item": item, "cents": cents})
	})
	fmt.Println("app listening on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Fprintln(os.Stderr, "app:", err)
		os.Exit(1)
	}
}
