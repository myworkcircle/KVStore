package main

import (
	"bufio"
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

func NewFileBufioDao(path string, fileId int) (FileRespository, error) {
	filePath := path + strconv.FormatInt(int64(fileId), 16)
	writerInstance, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("NewFileBufioDao | Error while opening file: ", err)
		return nil, err
	}
	writer := bufio.NewWriterSize(writerInstance, 4096)

	reader, err := os.OpenFile(filePath, os.O_RDONLY, 0755)
	if err != nil {
		fmt.Println("NewFileBufioDao | Error while opening file: ", err)
		return nil, err
	}

	fileInfo, err := reader.Stat()
	if err != nil {
		fmt.Println("NewFileBufioDao | Error while reading stat of file: ", err)
	}
	fileBufioDao := &FileBufioDao{
		writer: writer,
		reader: reader,
		offset: fileInfo.Size(),
	}
	return fileBufioDao, nil
}

func (f *FileBufioDao) Save(data []byte, toFlush bool) (int64, error) {
	write, err := f.writer.Write(data)
	if err != nil {
		return 0, err
	}
	offset := f.offset
	f.offset += int64(write)
	if toFlush {
		err = f.writer.Flush()
		if err != nil {
			fmt.Printf("Save | error while flushing data to disk: %s\n", err)
			return 0, err
		}
	}
	return offset, nil
}

func (f *FileBufioDao) SaveString(data string, toFlush bool) (int64, error) {
	write, err := f.writer.WriteString(data)
	if err != nil {
		return 0, err
	}
	offset := f.offset
	f.offset += int64(write)
	if toFlush {
		err = f.writer.Flush()
		if err != nil {
			fmt.Printf("Save | error while flushing data to disk: %s\n", err)
			return 0, err
		}
	}
	return offset, nil
}

func (f *FileBufioDao) Get(bytesToRead int, offset int64) ([]byte, error) {
	seekOffset, err := f.reader.Seek(offset, io.SeekStart)
	if err != nil {
		fmt.Println("Get | error while seeking to file: ", err)
		return nil, err
	}
	fmt.Printf("file size: %d\nfile offset: %d\n", f.GetFileSize(), offset)
	b := make([]byte, bytesToRead)
	_, err = f.reader.ReadAt(b, seekOffset)
	if err != nil {
		fmt.Println("Get | error while reading: ", err)
		return nil, err
	}
	return b, nil
}

func (f *FileBufioDao) GetOffset() int64 {
	return f.offset
}

func (f *FileBufioDao) GetFileName() string {
	return f.reader.Name()
}

func (f *FileBufioDao) GetFileSize() int64 {
	fileInfo, _ := f.reader.Stat()
	return fileInfo.Size()
}
