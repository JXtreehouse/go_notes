/*
 * @Author: your name
 * @Date: 2021-11-10 15:53:59
 * @LastEditTime: 2021-11-10 15:57:44
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_notes/src/unzip/unzip.go
 */
package unzip

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	Silence bool
)

func IsDir (path string) bool {
	
}
func Do(zipfile, targetDir string)  error {
	if targetDir == "" {
		targetDir = "./" //默认在本目录
	}

	if targetDir != "./" {
		if !IsDir()
	}
 }