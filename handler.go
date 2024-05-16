package main

import (
	"fmt"
	"time"
)

type Handler struct {
	diskStore     *DiskStore
	inMemoryStore *KeyValueStore
}

func (h *Handler) put(key, value string) error {
	valueSize := len([]byte(value))
	keySize := len([]byte(key))
	header := Header{
		TimeStamp: time.Now(),
		KeySize:   keySize,
		ValueSize: valueSize,
	}
	//fileValue := Record{
	//	Header: header,
	//	Value:  value,
	//	Key:    key,
	//}
	serialisedRecord, err := ConvertToBytes(header)
	if err != nil {
		return err
	}

	pos, err := h.diskStore.AppendToFile(serialisedRecord, key, []byte(value))
	if err != nil {
		return err
	}
	metadata := MetaData{
		FileID:    h.diskStore.activeFileID,
		ValueSize: valueSize,
		ValuePos:  pos,
		TimeStamp: header.TimeStamp,
	}
	h.inMemoryStore.Put(key, metadata)
	if h.diskStore.checkIfThresholdCrossed() {
		_, err = h.diskStore.CreateNewFile()
		if err != nil {
			return err
		}
	}
	fmt.Println("Write complete")
	return nil
}

func (j *Handler) get(key string) (string, error) {
	metaData := j.inMemoryStore.Get(key)
	byteValue, err := j.diskStore.GetValue(metaData)
	if err != nil {
		return "", err
	}
	actualValue := string(byteValue)
	fmt.Println("read complete")
	return actualValue, nil
}
