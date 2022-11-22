package hashtable

import (
	"testing"

	"github.com/andey-robins/bigdata/similarity/hash"
)

func TestHashtable(t *testing.T) {
	// create a hash table
	// note we use a tiny size to guarantee there are collisions and verify
	// they are handled properly. The proof there are collisions is by the
	// pigeon-hole principle. We insert 3 unique keys into a hashtable with
	// capacity of 2.
	h := New(2, hash.Sha256Wrapper)

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

	if h.Collisions() != 2 {
		t.Errorf("Incorrect number of collisions. exp=%v got=%v", 2, h.Collisions())
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
		t.Errorf("key1 reported as existing, even though it doesn't")
	}

	// replace key1 and double check Exists works for buckets with collisions
	h.Insert("key1", 1)
	if !h.Exists("key3") {
		t.Errorf("key3 reported as not existing, even though it does")
	}
	if !h.Exists("key1") {
		t.Errorf("key1 reported as not existing, even though it does")
	}

	keys := h.Keys()
	for i, key := range keys {
		if i == 0 && key != "key2" {
			t.Errorf("key2 not present in Keys()")
		}
		if i == 1 && key != "key3" {
			t.Errorf("key3 not present in Keys()")
		}
		if i == 2 && key != "key1" {
			t.Errorf("key1 not present in Keys()")
		}
	}
}
