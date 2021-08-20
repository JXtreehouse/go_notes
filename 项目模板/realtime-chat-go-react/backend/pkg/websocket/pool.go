/*
 * @Author: your name
 * @Date: 2021-07-02 15:09:01
 * @LastEditTime: 2021-07-02 16:04:15
 * @LastEditors: Please set LastEditors
 * @Description:实现了一个 Pool 机制，这将有效地允许我们跟踪我们有多少连接到我们的 WebSocket 服务器。
 * https://github.com/TutorialEdge/realtime-chat-go-react/blob/master/backend/pkg/websocket/pool.go
 * @FilePath: /realtime-chat-go-react/backend/pkg/websocket/pool.go
 */
package websocket

import "fmt"

type Pool struct {
	Register chan *Client
	Unregister chan *Client
	Clients map[*Client] bool
	Broadcast chan Message
}