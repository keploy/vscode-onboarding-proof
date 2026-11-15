package shop

import "testing"

func TestUserName(t *testing.T) {
	name, err := UserName(1)
	if err != nil {
		t.Fatalf("UserName(1): %v", err)
	}
	if name != "ada" {
		t.Fatalf("UserName(1) = %q, want %q", name, "ada")
	}
}

func TestPrice(t *testing.T) {
	cents, err := Price("book")
	if err != nil {
		t.Fatalf("Price(book): %v", err)
	}
	if cents != 4200 {
		t.Fatalf("Price(book) = %d cents, want 4200", cents)
	}
}
