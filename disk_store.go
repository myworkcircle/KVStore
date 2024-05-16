package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type DiskStore struct {
	directoryName     string
	currentActiveFile int
	activeFileID      string
	threshold         int64
}

func NewDiskStore(directoryName string, currentActiveFile int, activeFileID string, threshold int64) *DiskStore {
	return &DiskStore{directoryName: directoryName, currentActiveFile: currentActiveFile, activeFileID: activeFileID, threshold: threshold}
}

func (ds *DiskStore) Open(name string) {
	err := os.Mkdir("/Users/pallavagarwal/Documents/CaskDB/data_store", 0755)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			fmt.Println("Directory already exists")
		} else {
			fmt.Println("Error while creating directory: ", err)
			panic(err)
		}
	}
	ds.directoryName = "/Users/pallavagarwal/Documents/CaskDB/data_store"
	ds.currentActiveFile = -1
	ds.threshold = 20
	file, err := ds.CreateNewFile()
	if err != nil {
		fmt.Println(err)
	} else {
		ds.activeFileID = file.Name()
	}
	//defer os.RemoveAll(name)
}

func (ds *DiskStore) CreateNewFile() (*os.File, error) {
	ds.currentActiveFile += 1
	filePath := ds.directoryName + "/file-" + strconv.FormatInt(int64(ds.currentActiveFile), 16)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		if errors.Is(err, &os.PathError{}) {
			fmt.Println("file already exists")
		} else {
			fmt.Println(err)
			return nil, err
		}
	}
	ds.activeFileID = file.Name()
	return file, nil
}

func (ds *DiskStore) GetActiveFileID() string {
	return ds.activeFileID
}

func (ds *DiskStore) AppendToFile(data []byte, key string, value []byte) (int64, error) {
	fmt.Println("Writing data to file name: ", ds.activeFileID)
	file, err := os.OpenFile(ds.activeFileID, os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Printf("error while opening file for write: %s", err)
		return 0, err
	}

	_, err = file.Write(data)
	if err != nil {
		fmt.Printf("error while writing to file: %s", err)
		return 0, err
	}

	_, err = file.WriteString(key)
	if err != nil {
		fmt.Printf("error while writing key to file: %s", err)
		return 0, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("error while checking file stat: %s", err)
		return 0, err
	}
	offset := fileInfo.Size()

	_, err = file.Write(value)
	if err != nil {
		fmt.Printf("error while writing value to file: %s", err)
		return 0, err
	}
	//err = os.WriteFile(ds.GetActiveFileID(), data, os.ModeAppend)
	fileInfo, _ = file.Stat()
	fileSize := fileInfo.Size()
	fmt.Println("file size after write: ", fileSize)
	return offset, nil
}

func (ds *DiskStore) GetValue(metadata MetaData) ([]byte, error) {
	fmt.Println("Reading data from file name: ", metadata.FileID)
	file, err := os.Open(metadata.FileID)
	if err != nil {
		fmt.Println("error while reading file: ", err)
		return nil, err
	}
	offset := metadata.ValuePos
	valueSize := metadata.ValueSize
	ret, err := file.Seek(offset, io.SeekStart)
	if err != nil {
		fmt.Println("error while seeking to file: ", err)
		return nil, err
	}
	readValue := make([]byte, valueSize)
	_, err = file.ReadAt(readValue, ret)
	if err != nil {
		fmt.Println("error while reading from offset: ", err)
		return nil, err
	}
	return readValue, err
}

func (ds *DiskStore) checkIfThresholdCrossed() bool {
	file, err := os.Open(ds.activeFileID)
	if err != nil {
		fmt.Println("error while checking file size: ", err)
		return false
	}
	fileInfo, _ := file.Stat()
	if fileInfo.Size() > ds.threshold {
		return true
	}
	return false
}

func (ds *DiskStore) Compaction() error {
	entries, err := os.ReadDir(ds.directoryName)
	if err != nil {
		return err
	}

	filesInfo := make([]os.FileInfo, len(entries))
	for _, entry := range entries {
		file, err := entry.Info()
		if err != nil {
			fmt.Println("error while reading file: ", err)
		}
		filesInfo = append(filesInfo, file)
	}

	return nil
}
