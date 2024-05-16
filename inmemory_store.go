package main

type KeyValueStore struct {
	store map[string]MetaData
}

func NewKeyValueStore(store map[string]MetaData) *KeyValueStore {
	return &KeyValueStore{store: store}
}

func (k *KeyValueStore) Put(key string, meta MetaData) {
	k.store[key] = meta
}

func (k *KeyValueStore) Get(key string) MetaData {
	return k.store[key]
}
