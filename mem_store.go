package hap

import (
	"fmt"
	"strings"
)

type memStore map[string][]byte

func NewMemStore() Store {
	return memStore{}
}

func (fs memStore) Set(key string, value []byte) error {
	fs[key] = value

	return nil
}

func (fs memStore) Get(key string) ([]byte, error) {
	if v, ok := fs[key]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("no entry for key %s", key)
}

func (fs memStore) Delete(key string) error {
	delete(fs, key)

	return nil
}

func (fs memStore) KeysWithSuffix(s string) (keys []string, err error) {
	for k, _ := range fs {
		if strings.HasSuffix(k, s) {
			keys = append(keys, k)
		}
	}

	return
}
