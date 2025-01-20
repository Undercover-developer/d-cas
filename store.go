package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type PathKey struct {
	Pathname string
	Filename string
}

func (p PathKey) FullPath() string {
	return fmt.Sprintf("%s/%s", p.Pathname, p.Filename)
}

type PathTransformFunc func(string) PathKey

var DefaultPathTransformFunc = func(key string) PathKey {
	return PathKey{key, key}
}

func CASPathTransformFunc(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	hasStr := hex.EncodeToString(hash[:])

	blockSize := 10
	sliceLen := len(hasStr) / blockSize
	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
		from, to := i*blockSize, (i*blockSize)+blockSize
		paths[i] = hasStr[from:to]
	}
	return PathKey{
		Pathname: strings.Join(paths, "/"),
		Filename: hasStr,
	}
}

type StoreOpts struct {
	PathTransformFunc PathTransformFunc
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

// func (s *Store) readStream(key string) (io.Reader, error) {
// 	pathKey := s.PathTransformFunc(key)

// 	f, err := os.Open(p)
// }

func (s *Store) writeStream(key string, r io.Reader) error {
	pathKey := s.PathTransformFunc(key)
	if err := os.MkdirAll(pathKey.Pathname, os.ModePerm); err != nil {
		return err
	}

	pathAndFilename := pathKey.FullPath()

	f, err := os.Create(pathAndFilename)
	if err != nil {
		return err
	}
	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}
	log.Printf("written (%d) bytes to disk: %s", n, pathAndFilename)
	return nil
}
