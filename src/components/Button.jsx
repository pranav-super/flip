import React from 'react';
import './Button.css';

function Button({ label }) {
  return <button className="Button">{ label }</button>;
}

export default Button