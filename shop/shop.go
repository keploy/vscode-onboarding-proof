// Package shop is the sample's client of its dependency (../dep).
package shop

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

var client = &http.Client{Timeout: 10 * time.Second}

func base() string {
	if a := os.Getenv("CONVERT_DEP"); a != "" {
		return "http://" + a
	}
	return "http://127.0.0.1:18095"
}

func get(path string, into any) error {
	res, err := client.Get(base() + path)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("GET %s: %s", path, res.Status)
	}
	return json.NewDecoder(res.Body).Decode(into)
}

// UserName is the name of the user with this id.
func UserName(id int) (string, error) {
	var u struct {
		Name string `json:"name"`
	}
	if err := get(fmt.Sprintf("/users/%d", id), &u); err != nil {
		return "", err
	}
	return u.Name, nil
}

// Price is an item's price, in cents.
func Price(item string) (int, error) {
	var p struct {
		Cents int `json:"cents"`
	}
	if err := get("/prices/"+item, &p); err != nil {
		return 0, err
	}
	return p.Cents, nil
}
