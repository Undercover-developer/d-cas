package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "testfile"
	pathKey := CASPathTransformFunc(key)
	expectedOriginalKey := "e05fcb614ab36fdee72ee1f2754ed85e2bd0e8d0"
	if pathKey.Original != expectedOriginalKey {
		t.Errorf("have %s expected %s", pathKey.Original, expectedOriginalKey)
	}
	fmt.Printf("file path: %s\n", pathKey.Pathname)
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	data := bytes.NewReader([]byte("jpg bytes"))
	if err := s.WriteStream("samplepicture", data); err != nil {
		t.Error(err)
	}
}
