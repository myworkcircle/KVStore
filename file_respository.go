package main

type FileRespository interface {
	Save([]byte) (int64, error)
	SaveString(string) (int64, error)
	Get(int, int64) ([]byte, error)
	GetOffset() int64
	GetFileName() string
}
