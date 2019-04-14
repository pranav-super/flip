import React, { useState } from 'react';
import './Retrieve.css';
import Button from './Button';

function Retrieve() {
  const [ token, setToken ] = useState('');

  return (
    <div className="Retrieve">
      <input maxLength="5" type="text" placeholder="00000" onChange={ e => {
        setToken(e.target.value);
      } }/>
      <Button label={ Retrieve.name } onClick={ e => {
        console.log(token);
      }
      }/>
    </div>
  )
}

export default Retrieve