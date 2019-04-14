import React from 'react';
import './App.css';
import Upload from './Upload'
import Retrieve from './Retrieve'

function App() {
  return (
    <div className="App" onDragOver={e => e.preventDefault()}>
      <header className="App-header">
        flip
      </header>
      <Upload/>
      <Retrieve/>
    </div>
  );
}

export default App;
