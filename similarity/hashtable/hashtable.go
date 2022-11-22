package hashtable

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"

	"github.com/andey-robins/bigdata/similarity/ngram"
)

type Hashtable struct {
	table       []tableRow
	hashFunc    func(gram *ngram.Ngram) [32]byte
	truncLength int
	keyMap      map[string]bool
}

type tableRow struct {
	row []tableElement
}

type tableElement struct {
	key   string
	value int
}

// New will create a new hashtable with a capacity approximately the size given as an
// argument. The true size is the smallest power of 2 that is greater than the requested size
func New(size int, hashFunc func(gram *ngram.Ngram) [32]byte) *Hashtable {
	// find the smallest power of two range that nicely fits
	// the requested size
	for i := 3; i < 64; i++ {
		power := int(math.Pow(2, float64(i)))
		if power >= size {
			h := &Hashtable{}
			h.truncLength = power
			h.table = make([]tableRow, power)
			h.hashFunc = hashFunc
			return h
		}
	}

	// default should only happen if we try to create a hashtable with size
	// greater than uint64 max size. in that case, default to uint64 size
	h := &Hashtable{}
	h.table = make([]tableRow, int(math.Pow(2, 64.0)))
	// h.keyMap = make(map[string]bool)
	h.truncLength = 64
	h.hashFunc = hashFunc
	return h
}

// Insert inserts a new key/value pair into the hashtable.
// Should return an error if the key already exists.
func (h *Hashtable) Insert(key string, value int) error {
	addr := getHashedKey(key, h.truncLength, h.hashFunc)
	if h.Exists(key) {
		return errors.New("key already exists. use hashtable.Update to modify the value")
	}
	if len(h.table[addr].row) != 0 {
		h.table[addr].row = append(h.table[addr].row, tableElement{key, value})
		return nil
	}
	h.table[addr].row = append(make([]tableElement, 0), tableElement{key, value})
	// h.keyMap[key] = true
	return nil
}

// Update updates an existing key to be associated with a different value.
// Should return an error if the key doesn't already exist.
func (h *Hashtable) Update(key string, value int) error {
	addr := getHashedKey(key, h.truncLength, h.hashFunc)
	// fmt.Printf("for key=%v, hash=%v, val=%v\n", key, addr, value)
	if !keyExists(&h.table[addr], key) {
		return errors.New("key does not exist")
	}
	updateKey(key, value, &h.table[addr])
	return nil
}

// Delete deletes a key/value pair from the hashtable.
// Should return an error if the given key doesn't exist.
func (h *Hashtable) Delete(key string) error {
	addr := getHashedKey(key, h.truncLength, h.hashFunc)
	if !keyExists(&h.table[addr], key) {
		return errors.New("key does not exist")
	}
	removeKey(key, &h.table[addr])
	return nil
}

// Exists returns true if the key exists in the hashtable, false otherwise.
func (h *Hashtable) Exists(key string) bool {
	addr := getHashedKey(key, h.truncLength, h.hashFunc)
	return keyExists(&h.table[addr], key)
}

// Get returns the value associated with the given key.
// Should return an error if value doesn't exist.
func (h *Hashtable) Get(key string) (int, error) {
	addr := getHashedKey(key, h.truncLength, h.hashFunc)
	if v, err := getKey(key, &h.table[addr]); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

func (h *Hashtable) GetSimilarSentences() [][]string {
	allStrings := make([][]string, 0)
	for _, row := range h.table {
		rowStrings := make([]string, 0)
		if len(row.row) > 1 {
			for _, key := range row.row {
				rowStrings = append(rowStrings, key.key)
			}
		}
		allStrings = append(allStrings, rowStrings)
	}
	return allStrings
}

// Keys returns a list of the keys in the hashtable
func (h *Hashtable) Keys() []string {
	keys := make([]string, 0)

	for _, row := range h.table {
		for _, element := range row.row {
			keys = append(keys, element.key)
		}
	}

	return keys
}

func (h *Hashtable) Collisions() int {
	collisions := 0
	for _, row := range h.table {
		if len(row.row) > 1 {
			collisions += len(row.row) - 1
		}
	}
	return collisions
}

// Print will output the key and value for every existing entry in the table
func (h *Hashtable) Print() string {
	outString := ""
	for _, row := range h.table {
		// each row is a linked list
		for _, element := range row.row {
			outString += fmt.Sprintf("%v = %v\n", element.key, element.value)
		}
	}
	return outString
}

// generate the hashed key and parse it based on the current size of
// the hash table
func getHashedKey(key string, len int, hashFunc func(gram *ngram.Ngram) [32]byte) int {
	ngram := ngram.New(4, key)
	hash := hashFunc(ngram)
	switch len {
	case 8:
		return int(math.Abs(float64(int(hash[0]) % len)))
	case 16:
		return int(math.Abs(float64(int(binary.BigEndian.Uint16(hash[0:2])) % len)))
	case 32:
		return int(math.Abs(float64(int(binary.BigEndian.Uint32(hash[0:4])) % len)))
	default:
		return int(math.Abs(float64(int(binary.BigEndian.Uint64(hash[0:8])) % len)))
	}
}

// returns true if the key exists and false otherwise
func keyExists(r *tableRow, key string) bool {
	for _, e := range r.row {
		if e.key == key {
			return true
		}
	}
	return false
}

// returns the value of a key if the key is present
// and a "key does not exist" error if the key isn't there
func getKey(key string, r *tableRow) (int, error) {
	for _, e := range r.row {
		if e.key == key {
			return e.value, nil
		}
	}

	return 0, errors.New("key does not exist")
}

// updates the value of a key if it exists and does nothing if
// it doesn't exist
func updateKey(key string, value int, r *tableRow) {
	for i, e := range r.row {
		if e.key == key {
			r.row[i].value = value
		}
	}
}

// check a slice for a key, if it is present, remove it from the slice
func removeKey(key string, r *tableRow) {
	index := 0

	for i, e := range r.row {
		if e.key == key {
			index = i
		}
	}

	r.row = append(r.row[:index], r.row[index+1:]...)
}
