package main

type StorageEngine interface {
	Put([]byte, string, []byte) (int64, error)
	Get(data MetaData) ([]byte, error)
	GetActiveFileID() string
	checkIfThresholdCrossed() bool
	CreateNewFile(fileHandlerStrategy func(string, int) (FileRespository, error)) error
}
