/*
 * @Author: AlexZ33
 * @Date: 2021-04-16 17:46:15
 * @LastEditTime: 2021-04-16 18:21:55
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/Web框架/手写框架/TuringGo/turingGo.go
 */
package TuringGo

import (
	"fmt"
	"net/http"
)

// handlerFunc defines the request handler used by turingGo

type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine implement the interface of  ServeHTTP
type Engine struct {
	router map[string] HandlerFunc
}

// New is the constructor of turingGo.Engine
func New() *Engine {
	return &Engine(router: make(map[string] HandlerFunc))
}

func (engine *Engine) addRouter(method string, pattern string, handler HandlerFunc)