package data

import "bitcask-go/fio"

const DataFileNameSuffix = ".data"

// DataFile 数据文件
type DataFile struct {
	FileId    uint32        // 文件id
	WriteOff  int64         // 文件写到了哪个位置
	IoManager fio.IOManager // io 读写管理
}

// OpenDataFile 打开新的数据文件，应该要根据传入的fileId打开不同的文件吧
func OpenDataFile(dirPath string, fileId uint32) (*DataFile, error) {
	return nil, nil
}

func (df *DataFile) ReadLogRecord(offset int64) (*LogRecord, int64, error) {
	return nil, 0, nil
}

func (df *DataFile) Write(buf []byte) error {
	return nil
}

func (df *DataFile) Sync() error {
	return nil
}
