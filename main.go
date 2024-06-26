package main

import (
	"fmt"
)

func main() {
	keyDir := KeyValueStore{map[string]MetaData{}}
	//ds := DiskStore{keyDir: &keyDir}
	//ds.Open("data_store")
	//handler := Handler{
	//	diskStore:     &ds,
	//	inMemoryStore: &keyDir,
	//}
	//err := handler.put("one-1", "Value-1")
	//if err != nil {
	//	return
	//}
	//val1, err := handler.get("one-1")
	//if err == nil {
	//	fmt.Println("value for key: one-1 :", val1)
	//}
	//fmt.Println("===============")
	//
	//err = handler.put("two", "Value-2")
	//if err != nil {
	//	return
	//}
	//val2, err := handler.get("two")
	//if err == nil {
	//	fmt.Println("value for key: two :", val2)
	//}
	//fmt.Println("===============")
	//
	//err = handler.put("two-2", "Value-2")

	// ==========================================

	segmentDB, err := NewSegmentDB("/Users/pallavagarwal/Documents/CaskDB/data_store", 4096, &keyDir)
	if err != nil {
		fmt.Println("error while creating segment DB:", err)
		return
	}
	err = segmentDB.CreateNewFile(NewFileBufioDao)
	if err != nil {
		fmt.Println("error while creating new file:", err)
		return
	}

	handler2 := Handler{
		diskStore:     segmentDB,
		inMemoryStore: &keyDir,
	}

	err = handler2.put("one-1", "Value-1")
	if err != nil {
		fmt.Printf("Error | while putting value for key: %s err: %s\n", "one-1", err)
		return
	}

	val21, err := handler2.get("one-1")
	if err == nil {
		fmt.Println("value for key: one-1: ", val21)
	}
	fmt.Println("===============")

	err = handler2.put("two", "Value-2")
	if err != nil {
		return
	}
	val22, err := handler2.get("two")
	if err == nil {
		fmt.Println("value for key: two: ", val22)
	}
	fmt.Println("===============")

	err = handler2.put("two-2", "Value-2")
	val32, err := handler2.get("two-2")
	if err == nil {
		fmt.Println("value for key: two: ", val32)
	}

}
