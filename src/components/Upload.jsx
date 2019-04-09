import React, { useState } from 'react';
import './Upload.css';
import Button from './Button';

function Upload({ uploadInput }) {
  const [ key, setKey ] = useState('');
  const [ uploaded, setUploaded ] = useState('select/drag files');

  const handleUpload = (e) => {
    e.preventDefault();

    const data = new FormData();
    // TODO: Change key
    data.append('sampleFile', uploadInput.files[ 0 ]);

    fetch('http://localhost:80/upload', {
      method: 'POST',
      body: data,
    }).then((response) => {
      response.json().then((body) => {
        if (body.Key) {
          setKey("Retrieval key: " + body.Key);
        }
      });
    }).catch(() => {
      setUploaded("Please try again!");
    });
  };

  return (
    <div className="Upload">
      <form onSubmit={ handleUpload }>
        <label htmlFor="upload">
          <span>{ uploaded }</span>
        </label>
        <input id="upload"
               ref={
                 (ref) => {
                   uploadInput = ref;
                 }
               }
               type="file" multiple
        />
        <Button label={ Upload.name }/>
        <span className="key">{ key }</span>
      </form>
      <div id="or"><span>OR</span></div>
    </div>
  )
}

export default Upload