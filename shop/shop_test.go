package shop

import (
	"testing"

	"github.com/keploy/go-sdk/v3/keploytest"
)

func TestUserName(t *testing.T) {
	defer keploytest.T(t)()
	name, err := UserName(1)
	if err != nil {
		t.Fatalf("UserName(1): %v", err)
	}
	if name != "ada" {
		t.Fatalf("UserName(1) = %q, want %q", name, "ada")
	}
}

func TestPrice(t *testing.T) {
	defer keploytest.T(t)()
	cents, err := Price("book")
	if err != nil {
		t.Fatalf("Price(book): %v", err)
	}
	if cents != 4200 {
		t.Fatalf("Price(book) = %d cents, want 4200", cents)
	}
}
