import React from 'react';
import { Provider } from 'react-redux';
import { createStore, applyMiddleware } from 'redux';
import { BrowserRouter, Route } from 'react-router-dom';
import Home from './Home';
import Download from './Download';
import thunk from 'redux-thunk';

import rt from '../reducers';

const store = createStore(rt, applyMiddleware(thunk));

function App() {
  return (
    <BrowserRouter>
      <Provider store={ store }>
        <Route exact path={ "/" } component={ Home }/>
        <Route exact path={ "/:token" } component={ Download }/>
      </Provider>
    </BrowserRouter>
  );
}

export default App;
