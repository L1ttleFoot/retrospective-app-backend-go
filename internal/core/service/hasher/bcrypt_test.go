package core_service_hasher

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHasher_Hash(t *testing.T) {
	h := NewHasher(bcrypt.DefaultCost)

	hash, err := h.Hash("password123")
	if err != nil {
		t.Fatalf("Hash() error = %v", err)
	}
	if hash == "" {
		t.Fatal("Hash() returned empty string")
	}
	if hash == "password123" {
		t.Fatal("Hash() returned plain password")
	}
}

func TestHasher_Compare_Success(t *testing.T) {
	h := NewHasher(bcrypt.DefaultCost)

	hash, err := h.Hash("password123")
	if err != nil {
		t.Fatalf("Hash() error = %v", err)
	}

	err = h.Compare(hash, "password123")
	if err != nil {
		t.Fatalf("Compare() error = %v", err)
	}
}

func TestHasher_Compare_Fail(t *testing.T) {
	h := NewHasher(bcrypt.DefaultCost)

	hash, err := h.Hash("password123")
	if err != nil {
		t.Fatalf("Hash() error = %v", err)
	}

	err = h.Compare(hash, "wrongpassword")
	if err == nil {
		t.Fatal("Compare() expected error for wrong password, got nil")
	}
}

func TestHasher_Compare_InvalidHash(t *testing.T) {
	h := NewHasher(bcrypt.DefaultCost)

	err := h.Compare("invalidhash", "password123")
	if err == nil {
		t.Fatal("Compare() expected error for invalid hash, got nil")
	}
}

func TestHasher_DifferentCosts(t *testing.T) {
	for _, cost := range []int{bcrypt.MinCost, bcrypt.DefaultCost} {
		h := NewHasher(cost)

		hash, err := h.Hash("password123")
		if err != nil {
			t.Fatalf("Hash() error with cost %d = %v", cost, err)
		}

		err = h.Compare(hash, "password123")
		if err != nil {
			t.Fatalf("Compare() error with cost %d = %v", cost, err)
		}
	}
}

func TestHasher_HashProducesDifferentHashes(t *testing.T) {
	h := NewHasher(bcrypt.DefaultCost)

	hash1, err := h.Hash("password123")
	if err != nil {
		t.Fatalf("Hash() error = %v", err)
	}

	hash2, err := h.Hash("password123")
	if err != nil {
		t.Fatalf("Hash() error = %v", err)
	}

	if hash1 == hash2 {
		t.Fatal("Hash() produced same hash for same password (bcrypt should use salt)")
	}

	err = h.Compare(hash1, "password123")
	if err != nil {
		t.Fatalf("Compare() failed for hash1 = %v", err)
	}

	err = h.Compare(hash2, "password123")
	if err != nil {
		t.Fatalf("Compare() failed for hash2 = %v", err)
	}
}
