import React, { useState } from 'react';
import { MoonLoader } from 'react-spinners';
import './Button.css';

function Button({ label, onClick }) {
  const [ loading, setLoading ] = useState(false);

  // Temp onClick
  return (
    <button className={ loading ? "Load" : "Button" } onClick={ onClick }>
      {
        loading ? <MoonLoader sizeUnit={ "px" }
                              size={ 25 }
                              color={ '#02C3BD' }
        /> : label
      }
    </button>
  );
}

export default Button