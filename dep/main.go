// Command dep is the sample's one real dependency: a small HTTP service the
// client in ../shop calls. It listens on CONVERT_DEP (default
// 127.0.0.1:18095) and prints "dep ready" once it does, so the proof waits
// for that line rather than for a time.
package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	addr := os.Getenv("CONVERT_DEP")
	if addr == "" {
		addr = "127.0.0.1:18095"
	}
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("dep failed:", err)
		os.Exit(1)
	}
	fmt.Println("dep ready")
	users := map[string]string{"1": "ada", "2": "grace"}
	prices := map[string]int{"book": 4200, "pen": 150}
	mux := http.NewServeMux()
	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		name, ok := users[strings.TrimPrefix(r.URL.Path, "/users/")]
		if !ok {
			http.NotFound(w, r)
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]string{"name": name})
	})
	mux.HandleFunc("/prices/", func(w http.ResponseWriter, r *http.Request) {
		cents, ok := prices[strings.TrimPrefix(r.URL.Path, "/prices/")]
		if !ok {
			http.NotFound(w, r)
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]int{"cents": cents})
	})
	if err := http.Serve(l, mux); err != nil {
		fmt.Fprintln(os.Stderr, "dep:", err)
		os.Exit(1)
	}
}
