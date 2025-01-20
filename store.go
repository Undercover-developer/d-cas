package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
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

func (p PathKey) RootPath() string {
	return strings.Split(p.FullPath(), "/")[0]
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

func (s *Store) Has(key string) bool {
	pathKey := s.PathTransformFunc(key)
	_, err := os.Stat(pathKey.FullPath())
	if err == fs.ErrNotExist {
		return false
	}

	return true
}

func (s *Store) Delete(key string) error {
	pathKey := s.PathTransformFunc(key)
	defer func() {
		log.Printf("deleted [%s] from disk", pathKey.Filename)
	}()
	return os.RemoveAll(pathKey.RootPath())
}

func (s *Store) Read(key string) (io.Reader, error) {
	f, err := s.readStream(key)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, f)

	return buf, err
}

func (s *Store) readStream(key string) (io.ReadCloser, error) {
	pathKey := s.PathTransformFunc(key)
	return os.Open(pathKey.FullPath())
}

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
