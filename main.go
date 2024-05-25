package main

import (
	"fmt"
)

func main() {
	keyDir := KeyValueStore{map[string]MetaData{}}
	ds := DiskStore{keyDir: &keyDir}
	ds.Open("data_store")
	handler := Handler{
		diskStore:     &ds,
		inMemoryStore: &keyDir,
	}
	err := handler.put("one-1", "Value-1")
	if err != nil {
		return
	}
	val1, err := handler.get("one-1")
	if err == nil {
		fmt.Println("value for key: one-1 :", val1)
	}
	fmt.Println("===============")

	err = handler.put("two", "Value-2")
	if err != nil {
		return
	}
	val2, err := handler.get("two")
	if err == nil {
		fmt.Println("value for key: two :", val2)
	}
	fmt.Println("===============")

	err = handler.put("two-2", "Value-2")
	if err != nil {
		//
	}
	//testing.Benchmark(BenchmarkAppendToFile)
	//testing.Benchmark()
}
