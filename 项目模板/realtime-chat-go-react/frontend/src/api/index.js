/*
 * @Author: your name
 * @Date: 2021-07-01 17:54:41
 * @LastEditTime: 2021-07-01 21:14:33
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /realtime-chat-go-react/frontend/src/api/index.js
 */

var socket = new WebSocket("ws://localhost:8080/ws");

let connect = () => {
    console.log("尝试连接...")

    socket.onopen= () => {
        console.log("连接成功")
    };

    socket.onmessage = msg => {
        console.log(msg)
    }

    socket.onclose = event => {
        console.log("连接关闭", event)
    }

    socket.onerror = error => {
        console.log("连接错误", error)
    }
};

let sendMsg = msg => {
   console.log("sending msg:", msg);
//    console.log(msg)
   socket.send(msg)
}

export {
    connect,
    sendMsg
}


