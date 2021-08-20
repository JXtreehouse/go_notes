/*
 * @Author: your name
 * @Date: 2021-08-20 15:40:41
 * @LastEditTime: 2021-08-20 16:07:23
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/手写系列/手写数据库/owndb/utils/str.go
 */
package utils

import "strconv"

// Float64ToStr Convert type float64 to string
func Float64ToStr(val float64) string {
	return strconv.FormatFloat(val, 'f', -1, 64)
}

// StrToFloat64 convert type string to float64
func StrToFloat64(val string) (float64, error ) {
	return strconv.ParseFloat(val, 64)
}