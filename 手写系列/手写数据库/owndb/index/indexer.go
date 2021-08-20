/*
 * @Author: your name
 * @Date: 2021-07-08 16:09:17
 * @LastEditTime: 2021-07-08 16:17:01
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/手写系列/手写数据库/owndb/index/indexer.go
 */package index
 
 import(
	 "../owndb/storage"
 )

 // Indexer the data index info, stored in skip list

 type Indexer struct {
     Meta *storage.Meta  // metadata info
	 FileId unit32 // the size of entry
	 EntrySize unit32  // the size of entry
	 Offset int64 // entry data query start position
 }