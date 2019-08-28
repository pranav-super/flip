import React from 'react';
import Upload from './Upload';
import Retrieve from './Retrieve';

import './App.css';

function Home() {
  return (
    <div className="App" onDragOver={ e => e.preventDefault() }>
      <header className="App-header">
        flip
      </header>
      <Upload/>
      <Retrieve/>
    </div>
  );
}

export default Home;
