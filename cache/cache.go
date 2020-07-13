package cache

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// CacheDir represents the path on disk to the cache location
var CacheDir string

// Init initializes the named cache
func Init(name string) (err error) {
	CacheDir = filepath.Join(os.Getenv("HOME"), name)

	// Make sure CacheDir exists, relative to $HOME
	err = os.Mkdir(CacheDir, 0766)

	// Ignore error relating to dir existence
	if err != nil && !os.IsExist(err) {
		return err
	}

	return nil
}

// Get looks up cached key and returns its value
func Get(key string) (val string, err error) {
	// Get key path
	keyPath, err := getKeyPath(key)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadFile(keyPath)
	if os.IsNotExist(err) {
		return "", nil
	} else if err != nil {
		return "", err
	}

	return string(b), nil
}

// Set adds a key value pair to the cache
func Set(key string, value string) (err error) {
	// Get key path
	keyPath, err := getKeyPath(key)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(keyPath, []byte(value), 0644)
	if err != nil {
		return err
	}

	return nil
}

func getKeyPath(key string) (path string, err error) {
	// Get MD5 hash of key
	h := md5.New()
	_, err = io.WriteString(h, key)
	if err != nil {
		return "", err
	}

	hashedKey := hex.EncodeToString(h.Sum(nil))
	return filepath.Join(CacheDir, hashedKey), nil
}
