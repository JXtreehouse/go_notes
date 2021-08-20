/*
 * @Author: your name
 * @Date: 2021-07-08 15:57:50
 * @LastEditTime: 2021-07-09 17:48:09
 * @LastEditors: Please set LastEditors
 * @Description: 外层操作接口 - 字符串操作接口
 * @FilePath: /go_notes/手写系列/手写数据库/owndb/db_str.go
 */

 package owndb
 
 import (
	 "bytes"
	 "/owndb/storage"
 )

// Set 设置value以保存字符串值。 如果 key 已经包含一个值，它会被覆盖。
 func (db *OwnDB) Set (key, value []byte) error {
	 return db.doSet(key, value)
 }

 func (db *OwnDB) doSet(key, value []byte) (err error) {
	 // 检查key value是否合法
	 if err = db.checkKeyValue(key, value); err != nil {
		 return err
	 }
      // If the existed value is the same as the set value, nothing will be done.
	 if db.config.IdxMode == KeyValueMemMode {
		 if existVal, _ := db.Get(key); existVal!nil && bytes.Compare(existVal, value) == 0 {
			 return 
		 }
	 }
	 db.strIndex.mu.Lock()
	 defer db.strIndex.mu.Unlock()
     
	 // string  标识数据类型
	 // StringSet  标识数据的操作类型
	 // 在idx.go文件中有数据类型和操作类型的枚举定义
	 e:= storage.NewEntryNoExtra(key, value, String, StringSet)
	 if err:= db.store(e); err!=nil {
		 return err
	 }
	 
 }