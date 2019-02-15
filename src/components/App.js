import React, { Component } from 'react';
import './App.css';
import Upload from './Upload'
import Retrieve from './Retrieve'

class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          flip
        </header>
        <Upload/>
        <Retrieve/>
      </div>
    );
  }
}

export default App;
