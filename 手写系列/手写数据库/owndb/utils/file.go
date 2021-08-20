/*
 * @Author: your name
 * @Date: 2021-08-20 14:15:06
 * @LastEditTime: 2021-08-20 15:32:34
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/手写系列/手写数据库/owndb/utils/file.go
 */

 package utils
 
 import(
	 "io"
	 "io/ioutil"
	 "os"
	 "path"
 )

 // exist check if the directory or file exits
 func CopyDir(src string,  dst string) error {
	 var {
		 err error
		 dir []os.FileInfo
		 srcInfo os.FileInfo
	 }

	 if srcInfo, err = os.Stat(src); err !=nil {
		 return err
	 }

	 if err=os.MkdirAll(dst, srcInfo.Mode()); err!=nil {
		 return err
	 }

	 if dir, err = ioutil.ReadDir(src); err !=nil {
		 return err
	 }

	 for _, fd := range dir {
		 srcPath := path.Join(src, fd.Name())
		 dstPath := path.Join(dst, fd.Name())

		 if fd.IsDir() {
			 if err = CopyDir(srcPath, dstPath); err!=nil {
				 return err
			 }
		 } else {
			 if err = CopyFile(srcPath, dstPath); err !=nil {
				 return err
			 }
		 }
	 }

	 return nil
 }

 // CopyFIle copy as single file 
 func CopyFile(src, dst string) error {
	 var (
		 err error
		 srcFile *os.File
		 dstFile *os.File
		 srcInfo os.FileInfo
	 )

	 if srcFile, err = os.Open(src); err !=nil {
		 return err
	 }

	 defer srcFile.Close()

	 if dstFile,err = os.Create(dst); err != nil {
		 return err
	 }

	 defer dstFile.Close()

	 if_, err = io.Copy(dstFile, srcFile); err!=nil {
		 return err
	 }

	 if srcInfo, err = os.Stat(src); err!=nil {
		 return err
	 }

	 return os.Chmod(dst, srcInfo.Mode())
 }