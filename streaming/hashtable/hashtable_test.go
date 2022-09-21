package hashtable

import (
	"crypto/sha256"
	"testing"
)

func TestHashtable(t *testing.T) {
	// create a hash table
	// note we use a tiny size to guarantee there are collisions and verify
	// they are handled properly. The proof there are collisions is by the
	// pigeon-hole principle. We insert 3 unique keys into a hashtable with
	// capacity of 2.
	h := New(2, sha256.Sum256)

	// check multiple insert happy paths
	if err := h.Insert("key1", 1); err != nil {
		t.Errorf("Unable to insert into hash table. err=%v", err)
	}
	if err := h.Insert("key2", 2); err != nil {
		t.Errorf("Unable to insert into hash table. err=%v", err)
	}
	if err := h.Insert("key3", 3); err != nil {
		t.Errorf("Unable to insert into hash table. err=%v", err)
	}

	// Check insert error paths
	if err := h.Insert("key3", 30); err == nil {
		t.Errorf("Tried to insert with an existing key, expected an error")
	}

	// check multiple Get happy paths
	if val, err := h.Get("key1"); err != nil && val != 1 {
		t.Errorf("Unable to retrieve key1. got=%v exp=1", val)
	}
	if val, err := h.Get("key2"); err != nil && val != 2 {
		t.Errorf("Unable to retrieve key1. got=%v exp=2", val)
	}
	if val, err := h.Get("key3"); err != nil && val != 3 {
		t.Errorf("Unable to retrieve key1. got=%v exp=3", val)
	}

	// Check get error paths
	if _, err := h.Get("Nonexistent"); err == nil {
		t.Errorf("retreived value for nonexistent value. Expected an error")
	}

	// Check update happy path
	if err := h.Update("key1", 10); err != nil {
		t.Errorf("Unable to update key. err=%v", err)
	}
	if val, err := h.Get("key1"); err != nil && val != 10 {
		t.Errorf("Unable to retrieve updated key1. got=%v exp=10", val)
	}

	// Check update error paths
	if err := h.Update("Nonexistent", -1); err == nil {
		t.Errorf("updated value for nonexistent value. Expected an error")
	}

	// Check delete happy path
	if err := h.Delete("key1"); err != nil {
		t.Errorf("Unable to delete key1. err=%v", err)
	}

	// Check delete error paths
	if err := h.Delete("Nonexistent"); err == nil {
		t.Errorf("deleted value for nonexistent value. Expected an error")
	}

	// Check Exists possible results
	if !h.Exists("key2") {
		t.Errorf("key2 reported as not existing, even though it does")
	}
	if h.Exists("key1") {
		t.Errorf("key2 reported as existing, even though it doesn't")
	}
}
