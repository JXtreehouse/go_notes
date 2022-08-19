/*
 * @Author: your name
 * @Date: 2021-05-26 15:34:00
 * @LastEditTime: 2021-05-26 15:39:22
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/base/sync_lock_unlock.go
 */
 package bufferChannel
 
 import (
	 "time"
	 "fmt"
	 "sync"
 )

 func main() {
	 var mutex sync.Mutex
	 fmt.Println("Lock the lock")
	 mutex.Lock()
	 fmt.Println("The lock is locked")
	 channels := make(chan, int)
 }
