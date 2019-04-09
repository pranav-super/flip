import React from 'react';
import './Retrieve.css';
import Button from './Button';

function Retrieve() {
  return (
    <div className="Retrieve">
      <form>
        <input maxLength="5" type="text" placeholder="00000"/>
        <Button label={ Retrieve.name }/>
      </form>
    </div>
  )
}

export default Retrieve