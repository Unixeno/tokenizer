// Package vocab provides datastructure and methods to work on word token dictionary
package vocab

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	// "strings"
)

type Idx int32

// Int32 returns int32 representation of an Idx
func (idx Idx) Int32() int32 {
	return int32(idx)
}

// Dict is a container for word token dictionary.
// It map token to key and autogenerated index as value
// Dict will expose methods to access its internal vocab
type Dict struct {
	tokens map[string]Idx
}

// New returns a vocab dict from the given tokens
func New(tokens []string) Dict {
	v := make(map[string]Idx, len(tokens))
	for i, t := range tokens {
		v[t] = Idx(i)
	}
	return Dict{tokens: v}
}

// FromFile will read a newline delimited file into a Dict
// It expects single token each line
func FromFile(path string) (Dict, error) {
	f, err := os.Open(path)
	if err != nil {
		return Dict{}, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	voc := Dict{tokens: map[string]Idx{}}
	for scanner.Scan() {
		voc.Add(scanner.Text())
	}

	return voc, nil
}

// Add adds an item to the voc dict if not existing
func (v *Dict) Add(token string) error {
	// If exist do nothing
	if v.HasToken(token) {
		err := errors.New("Token string already existed.")
		return err
	}

	v.tokens[token] = Idx(v.Size())
	return nil
}

// Index returns index of word token in vocab dict.
// It Will be negative if token doesn't exist.
func (v *Dict) Index(token string) (Idx, error) {
	i, ok := v.tokens[token]
	if !ok {
		err := errors.New("Token string does not exist.")
		return Idx(-1), err
	}
	return Idx(i), nil
}

// Token get a token by given index, returns the empty string if not existing
func (v *Dict) Token(idx Idx) (token string) {
	for k, val := range v.tokens {
		if Idx(val) == idx {
			token = k
			break
		}
	}
	// not found. Return empty
	return fmt.Sprintf(token)
}

// HasIdx returns true if the vocab contains a token for given index
func (v *Dict) HasIdx(idx Idx) bool {
	found := false
	for _, val := range v.tokens {
		if Idx(val) == idx {
			found = true
			break
		}
	}

	return found
}

// HasToken returns true if vocab contains token
func (v *Dict) HasToken(token string) bool {

	_, ok := v.tokens[token]

	return ok

}

// Size returns the size of the vocabulary
func (v *Dict) Size() int {
	return len(v.tokens)
}
