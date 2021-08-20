/*
 * @Author: your name
 * @Date: 2021-07-08 16:47:33
 * @LastEditTime: 2021-07-08 17:00:05
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/手写系列/手写数据库/owndb/storage/entry.go
 */



func newInternal(key, value, extra []byte, state unit16, timestamp unit64) *Entry {
	return &Entry {
		state: state, 
		Timestamp: timestamp,
		Meta: &Meta{
			Key:       key,
			Value:     value,
			Extra:     extra,  // 操作Entry所需的额外信息
			KeySize:   uint32(len(key)),
			ValueSize: uint32(len(value)),
			ExtraSize: uint32(len(extra)),
		}
	}
} 
 // NewEntryNoExtra create a new entry without extra info 
 func NewEntryNoExtra(key, value []byte, t, mark uint16) *Entry {
	return NewEntry(key, value, nil, t, mark)
}

