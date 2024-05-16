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

func (k *KeyValueStore) Delete(key string) {
	if _, ok := k.store[key]; ok {
		k.Delete(key)
	}
}

func (k *KeyValueStore) GetAll() []string {
	var keys []string
	for key, _ := range k.store {
		keys = append(keys, key)
	}
	return keys
}
