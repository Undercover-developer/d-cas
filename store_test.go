package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "testfile"
	pathKey := CASPathTransformFunc(key)
	expectedOriginalKey := "e05fcb614ab36fdee72ee1f2754ed85e2bd0e8d0"
	if pathKey.Filename != expectedOriginalKey {
		t.Errorf("have %s expected %s", pathKey.Filename, expectedOriginalKey)
	}
	fmt.Printf("file path: %s\n", pathKey.Pathname)
}

func TestStore(t *testing.T) {
	//test write stream
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "samplepicture"
	data := []byte("jpg bytes")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	//test read stream
	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := io.ReadAll(r)
	if err != nil {
		t.Error(err)
	}

	if string(b) != string(data) {
		t.Errorf("wants %s have %s\n", data, b)
	}
}
