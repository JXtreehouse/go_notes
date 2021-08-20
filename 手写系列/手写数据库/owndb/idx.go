/*
 * @Author: your name
 * @Date: 2021-07-09 17:48:52
 * @LastEditTime: 2021-07-09 18:09:57
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/手写系列/手写数据库/owndb/idx.go
 */
package owndb

import (
	"owndb/ds/list"
	"owndb/index"
	"owndb/storage"
	"owndb/utils"
	"io"
	"log"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// DataType Define the data structure type
type DataType = unit16

// Five different data types, support String, List, Hash, Set, Sorted Set right now

const (
	String DataType = iota
	List
	Hash
	Set
	ZSet
)

// The operation of a  String Type , willl be a part of Entry, the same for the other four types
const (
	StringSet unit16 = iota
	StringRem
	StringExpire
	StringPersist
)

// The operations of List
const (
	ListLPush unit16 = iota
	ListRPush 
	ListLPop
	ListRPop
	ListLRem
	ListLInsert
	ListLSet
	ListTrim
	ListLClear
	ListLExpire

)