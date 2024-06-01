package main

import (
	"errors"
	"fmt"
	"os"
)

type SegmentDB struct {
	directoryName  string
	threshold      int64
	keyDir         *KeyValueStore // per file or single for whole caskDB?
	fileDao        FileRespository
	activeFileID   string
	currentFileNum int
}

func NewSegmentDB(directoryName string, threshold int64, keyValueStore *KeyValueStore) (*SegmentDB, error) {
	err := os.Mkdir(directoryName, 0755)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			fmt.Println("Directory already exists")
		} else {
			fmt.Println("Error while creating directory: ", err)
			panic(err)
		}
	}
	ds := &SegmentDB{
		directoryName:  directoryName,
		threshold:      threshold,
		activeFileID:   "",
		keyDir:         keyValueStore,
		currentFileNum: -1,
	}

	return ds, nil
}

func (ds *SegmentDB) CreateNewFile(fileHandlerStrategy func(string, int) (FileRespository, error)) error {
	fileRepostiory, err := fileHandlerStrategy(ds.directoryName+"/file-", ds.currentFileNum+1)
	if err != nil {
		fmt.Println("Error while creating new file: ", err)
		return err
	}
	ds.currentFileNum = ds.currentFileNum + 1
	ds.activeFileID = fileRepostiory.GetFileName()
	fmt.Println("New File ID: ", ds.activeFileID)
	ds.fileDao = fileRepostiory
	return nil
}

func (ds *SegmentDB) GetActiveFileID() string {
	return ds.activeFileID
}

func (ds *SegmentDB) Put(data []byte, key string, value []byte) (int64, error) {
	_, err := ds.fileDao.Save(data)
	if err != nil {
		fmt.Printf("error while writing to file: %s", err)
		return 0, err
	}

	_, err = ds.fileDao.SaveString(key)
	if err != nil {
		fmt.Printf("error while writing key to file: %s", err)
		return 0, err
	}

	offset := ds.fileDao.GetOffset()

	_, err = ds.fileDao.Save(value)
	if err != nil {
		fmt.Printf("error while writing value to file: %s", err)
		return 0, err
	}

	return offset, nil
}

func (ds *SegmentDB) Get(metadata MetaData) ([]byte, error) {
	offset := metadata.ValuePos
	valueSize := metadata.ValueSize
	resp, err := ds.fileDao.Get(valueSize, offset)
	if err != nil {
		fmt.Println("error while reading from offset: ", err)
		return nil, err
	}
	return resp, err
}

func (ds *SegmentDB) checkIfThresholdCrossed() bool {
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
