package main

import (
	"fmt"
	"os"
	"testing"
)

//func BenchmarkTestAppendToFile(b *testing.B) {
//	ds := DiskStore{
//		directoryName:     "/Users/pallavagarwal/Documents/CaskDB/data_store",
//		currentActiveFile: 0,
//		activeFileID:      "/Users/pallavagarwal/Documents/CaskDB/data_store/file-0",
//		threshold:         600,
//	}
//	tc := []struct {
//		header []byte
//		key    string
//		val    []byte
//	}{
//		{
//			header: []byte("header"),
//			key:    "key1",
//			val:    []byte(dummy),
//		},
//	}
//	for i := 0; i < b.N; i++ {
//		_, err := ds.Put(tc[0].header, tc[0].key, tc[0].val)
//		if err != nil {
//			return
//		}
//	}
//	err := os.Remove("/Users/pallavagarwal/Documents/CaskDB/data_store/file-0")
//	if err != nil {
//		fmt.Println(err)
//	}
//}

// this performs better when we have made a single write operation for key value pair
// to be flushed at the end only while writing the final value into file
func BenchmarkTestAppendToFileWithBufio(b *testing.B) {
	fileBufioDao, _ := NewFileBufioDao("/Users/pallavagarwal/Documents/CaskDB/data_store/file-", 1)

	ds := SegmentDB{
		directoryName:  "/Users/pallavagarwal/Documents/CaskDB/data_store",
		currentFileNum: 1,
		activeFileID:   "/Users/pallavagarwal/Documents/CaskDB/data_store/file-1",
		threshold:      600,
		fileDao:        fileBufioDao,
	}
	tc := []struct {
		header []byte
		key    string
		val    []byte
	}{
		{
			header: []byte("header"),
			key:    dummy,
			val:    []byte(dummy),
		},
	}
	for i := 0; i < b.N; i++ {
		_, err := ds.Put(tc[0].header, tc[0].key, tc[0].val)
		if err != nil {
			return
		}
	}
	err := os.Remove("/Users/pallavagarwal/Documents/CaskDB/data_store/file-1")
	if err != nil {
		fmt.Println(err)
	}
}
func BenchmarkTestAppendToFileWithFileDao(b *testing.B) {
	fileDao, err := NewFileDao("/Users/pallavagarwal/Documents/CaskDB/data_store/file-", 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	ds := SegmentDB{
		directoryName:  "/Users/pallavagarwal/Documents/CaskDB/data_store",
		currentFileNum: 2,
		activeFileID:   "/Users/pallavagarwal/Documents/CaskDB/data_store/file-2",
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
			key:    dummy,
			val:    []byte(dummy),
		},
	}
	for i := 0; i < b.N; i++ {
		_, err := ds.Put(tc[0].header, tc[0].key, tc[0].val)
		if err != nil {
			return
		}
	}
	err = os.Remove("/Users/pallavagarwal/Documents/CaskDB/data_store/file-2")
	if err != nil {
		fmt.Println(err)
	}
}
