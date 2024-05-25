package main

import (
	"bufio"
	"os"
)

type FileHandler struct {
	reader *os.File
	writer *bufio.Writer
	offset int64
}

func NewFileHandler(reader *os.File, writer *bufio.Writer, offset int64) *FileHandler {
	return &FileHandler{reader: reader, writer: writer, offset: offset}
}

func (f *FileHandler) put(data []byte) (int64, error) {
	write, err := f.writer.Write(data)
	if err != nil {
		return 0, err
	}
	offset := f.offset
	f.offset += int64(write)
	return offset, nil
}

func (f *FileHandler) get(offset int64) ([]byte, error) {
	var b []byte
	_, err := f.reader.ReadAt(b, offset)
	if err != nil {
		return nil, err
	}
	return b, nil
}
