/*
 * @Author: your name
 * @Date: 2021-07-02 16:37:51
 * @LastEditTime: 2021-07-02 16:49:42
 * @LastEditors: Please set LastEditors
 * @Description: https://tutorialedge.net/projects/building-security-tools-in-go/building-port-scanner-go/
 * @FilePath: /go_notes/base/go_scan_a_port.go
 */
package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func scanPort(protocol, hostname string , port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60 * time.Second)

	if err != nil {
		return false
	}
}