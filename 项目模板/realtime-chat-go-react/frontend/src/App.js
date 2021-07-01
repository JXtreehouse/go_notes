/*
 * @Author: your name
 * @Date: 2021-07-01 17:21:16
 * @LastEditTime: 2021-07-01 21:07:10
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /realtime-chat-go-react/frontend/src/App.js
 */
import logo from './logo.svg';
import './App.css';

import {connect, sendMsg} from './api'



function App() {
  connect()
   return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
        <button onClick={send}> 点击</button>
      </header>
      
    </div>
  );
}

function send () {
  console.log("hello");
  sendMsg("hello")
}

export default App;
