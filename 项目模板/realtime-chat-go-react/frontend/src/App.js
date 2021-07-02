/*
 * @Author: your name
 * @Date: 2021-07-01 17:21:16
 * @LastEditTime: 2021-07-02 14:53:37
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /realtime-chat-go-react/frontend/src/App.js
 */
import React, { Component } from "react";

import './App.css';

import {connect, sendMsg} from './api'
import Header from './components/Header/Header';
import ChatInput from './components/ChatInput/ChatInput';
import ChatHistory from './components/ChatHistory/ChatHistory';

class App extends Component {
  constructor(props) {
      super(props);
      this.state = {
          chatHistory: []
      }
  }

  componentDidMount() {
      connect((msg) => {
          console.log("New Message")
      this.setState(prevState => ({
          chatHistory: [...this.state.chatHistory, msg]
  }))
      console.log(this.state);
  });
  }

  send(event) {
      if(event.keyCode === 13) {
          sendMsg(event.target.value);
          event.target.value = "";
      }
  }

  render() {
      return (
      <div className="App">
          <Header />
          <ChatHistory chatHistory={this.state.chatHistory} /> 
          <ChatInput send={this.send} /> 
      </div>
  );
  }
}

export default App;
