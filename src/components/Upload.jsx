import React, { useState, useCallback, useMemo } from 'react';
import './Upload.css';
import Button from './Button';
import { useDropzone } from 'react-dropzone';

const zone = {
  fontStyle: 'italic',
  color: 'rgba(0, 0, 0, 0.5)',
  cursor: 'pointer',
  border: '2px dashed #02C3BD',
  padding: '50px 0',
  width: '400px',
  marginBottom: '15px',
};

const zoneActive = {
  background: 'repeating-linear-gradient(-45deg, white, white 10px, rgba(0, 0, 0, 0.05) 10px, rgba(0, 0, 0, 0.05) 20px)',
};

function Upload({ uploadInput }) {
  const [ key, setKey ] = useState('');
  const [ uploaded, setUploaded ] = useState('select/drag files');
  const [ hover, setHover ] = useState(false);
  const onDrop = useCallback(acceptedFiles => {
    console.log(acceptedFiles);
  }, []);
  const { getRootProps, getInputProps, isDragActive } = useDropzone({ onDrop });
  const style = useMemo(() => ({
    ...zone,
    ...(isDragActive || hover ? zoneActive : {}),
  }));

  const handleUpload = (e) => {
    e.preventDefault();

    const data = new FormData();
    data.append('file', uploadInput.files[ 0 ]);

    fetch('http://localhost:80/buffer', {
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
        <label htmlFor="upload" { ...getRootProps({
          style,
          onMouseEnter: () => setHover(true),
          onMouseLeave: () => setHover(false),
        }
          ) }>
          <span { ...getInputProps() }>{ uploaded }</span>
        </label>
        <input id="upload"
               ref={
                 (ref) => {
                   uploadInput = ref;
                 }
               }
               type="file"
               multiple
        />
        <Button label={ Upload.name }/>
        <span className="key">{ key }</span>
      </form>
      <div id="or"><span>OR</span></div>
    </div>
  )
}

export default Upload