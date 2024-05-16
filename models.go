package main

import (
	"bytes"
	"encoding/binary"
	"time"
)

type Record struct {
	Header Header
	Key    string
	Value  string
}

type Header struct {
	KeySize   int
	ValueSize int
	TimeStamp time.Time
}

func (h *Header) encode(buf *bytes.Buffer) error {
	return binary.Write(buf, binary.LittleEndian, h)
}

type MetaData struct {
	FileID    string
	ValueSize int
	ValuePos  int64
	TimeStamp time.Time
}
