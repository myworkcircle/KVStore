package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type FileDao struct {
	reader *os.File
	writer *os.File
	offset int64
}

func NewFileDao(path string, fileId int) (FileRespository, error) {
	filePath := path + strconv.FormatInt(int64(fileId), 16)
	writer, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		if errors.Is(err, &os.PathError{}) {
			fmt.Println("file already exists")
		} else {
			fmt.Println(err)
			return nil, err
		}
	}

	reader, err := os.OpenFile(filePath, os.O_RDONLY, 0755)
	if err != nil {
		if errors.Is(err, &os.PathError{}) {
			fmt.Println("file already exists")
		} else {
			fmt.Println(err)
			return nil, err
		}
	}

	fileDao := &FileDao{
		writer: writer,
		reader: reader,
		offset: 0,
	}
	return fileDao, nil
}

func (f *FileDao) Save(data []byte, _ bool) (int64, error) {
	write, err := f.writer.Write(data)
	if err != nil {
		return 0, err
	}
	offset := f.offset
	f.offset += int64(write)
	return offset, nil
}

func (f *FileDao) SaveString(data string, _ bool) (int64, error) {
	write, err := f.writer.WriteString(data)
	if err != nil {
		return 0, err
	}
	offset := f.offset
	f.offset += int64(write)
	return offset, nil
}

func (f *FileDao) Get(bytesToRead int, offset int64) ([]byte, error) {
	seekOffset, err := f.reader.Seek(offset, io.SeekStart)
	if err != nil {
		fmt.Println("error while seeking to file: ", err)
		return nil, err
	}

	b := make([]byte, bytesToRead)
	_, err = f.reader.ReadAt(b, seekOffset)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (f *FileDao) GetOffset() int64 {
	return f.offset
}

func (f *FileDao) GetFileName() string {
	return f.reader.Name()
}
