package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func ConvertToBytes(p interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		fmt.Println("encode error:", err)
		return nil, err
	}
	return buf.Bytes(), nil
}
