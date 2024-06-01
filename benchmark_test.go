package main

import "testing"

func BenchmarkTestAppendToFile(b *testing.B) {
	ds := DiskStore{
		directoryName:     "/Users/pallavagarwal/Documents/CaskDB/data_store",
		currentActiveFile: 0,
		activeFileID:      "/Users/pallavagarwal/Documents/CaskDB/data_store/file-0",
		threshold:         600,
	}
	tc := []struct {
		header []byte
		key    string
		val    []byte
	}{
		{
			header: []byte("header"),
			key:    "key1",
			val:    []byte("value1"),
		},
	}
	for i := 0; i < b.N; i++ {
		_, err := ds.Put(tc[0].header, tc[0].key, tc[0].val)
		if err != nil {
			return
		}
	}
}

func BenchmarkTestAppendToFileWithBufio(b *testing.B) {
	fileDao, _ := NewFileBufioDao("/Users/pallavagarwal/Documents/CaskDB/data_store/file-0", 0)

	ds := SegmentDB{
		directoryName:  "/Users/pallavagarwal/Documents/CaskDB/data_store",
		currentFileNum: 0,
		activeFileID:   "/Users/pallavagarwal/Documents/CaskDB/data_store/file-0",
		threshold:      600,
		fileDao:        fileDao,
	}
	tc := []struct {
		header []byte
		key    string
		val    []byte
	}{
		{
			header: []byte("header"),
			key:    "key1",
			val:    []byte("value1"),
		},
	}
	for i := 0; i < b.N; i++ {
		_, err := ds.Put(tc[0].header, tc[0].key, tc[0].val)
		if err != nil {
			return
		}
	}
}
