package data

// LogRecordType 代表这个记录的类型，因为只往磁盘中追加记录，所以需要标识这个记录的类型
type LogRecordType = byte

const (
	LogRecordNormal LogRecordType = iota
	LogRecordDeleted
)

// LogRecordPos 这个代表的是内存中的索引信息，里面的字段指示了磁盘的位置，而真实存储在磁盘上的结构体不是这个
type LogRecordPos struct {
	Fid    uint32
	Offset int64 // 文件偏移量
}

// LogRecord 真实存储在磁盘中的结构体
type LogRecord struct {
	Key   []byte
	Value []byte
	Type  LogRecordType
}

// EncodeLogRecord 对 LogRecord 进行编码，返回字节数组及长度
func EncodeLogRecord(logRecord *LogRecord) ([]byte, int64) {
	return nil, 0
}