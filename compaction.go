package main

func (ds *DiskStore) Compaction() error {
	//tempFile, err := os.Create(ds.directoryName)
	//if err != nil {
	//	fmt.Println("error while creating temp file: ", err)
	//	return err
	//}
	//for key, val := range ds.keyDir.store {
	//	if val.FileID != ds.activeFileID {
	//		value, getErr := ds.Get(MetaData{
	//			FileID:    val.FileID,
	//			ValueSize: val.ValueSize,
	//			ValuePos:  val.ValuePos,
	//			TimeStamp: val.TimeStamp,
	//		})
	//		if getErr != nil {
	//			fmt.Println("Compaction | error while getting value: ", getErr)
	//		}
	//		_, err2 := ds.Put(tempFile.Name(), key, value)
	//		if err2 != nil {
	//			return err2
	//		}
	//	}
	//}
	return nil
}

//
//func (ds *DiskStore) listFilesInDir() ([]os.FileInfo, error) {
//	entries, err := os.ReadDir(ds.directoryName)
//	if err != nil {
//		return nil, err
//	}
//
//	filesInfo := make([]os.FileInfo, len(entries))
//	for _, entry := range entries {
//		file, err := entry.Info()
//		if err != nil {
//			fmt.Println("error while reading file: ", err)
//		}
//		filesInfo = append(filesInfo, file)
//	}
//	return filesInfo, nil
//}
