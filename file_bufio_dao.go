package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type FileBufioDao struct {
	reader *os.File
	writer *bufio.Writer
	offset int64
}

func NewFileBufioDao(path string, fileId int) (*FileBufioDao, error) {
	filePath := path + strconv.FormatInt(int64(fileId), 16)
	writerInstance, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		if errors.Is(err, &os.PathError{}) {
			fmt.Println("file already exists")
			return nil, err
		} else {
			fmt.Println(err)
			return nil, err
		}
	}
	writer := bufio.NewWriter(writerInstance)

	reader, err := os.OpenFile(filePath, os.O_RDONLY, 0755)
	if err != nil {
		if errors.Is(err, &os.PathError{}) {
			fmt.Println("file already exists")
		} else {
			fmt.Println(err)
			return nil, err
		}
	}

	fileDao := &FileBufioDao{
		writer: writer,
		reader: reader,
		offset: 0,
	}
	return fileDao, nil
}

func (f *FileBufioDao) Save(data []byte) (int64, error) {
	write, err := f.writer.Write(data)
	if err != nil {
		return 0, err
	}
	offset := f.offset
	f.offset += int64(write)
	return offset, nil
}

func (f *FileBufioDao) SaveString(data string) (int64, error) {
	write, err := f.writer.WriteString(data)
	if err != nil {
		return 0, err
	}
	offset := f.offset
	f.offset += int64(write)
	return offset, nil
}

func (f *FileBufioDao) Get(bytesToRead int, offset int64) ([]byte, error) {
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

func (f *FileBufioDao) GetOffset() int64 {
	return f.offset
}
