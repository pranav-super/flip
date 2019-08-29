import React, { useState, useCallback, useMemo } from 'react';
import './Upload.css';
import Button from './Button';
import { useDropzone } from 'react-dropzone';

const zone = {
  fontStyle: 'italic',
  color: 'rgba(0, 0, 0, 0.5)',
  cursor: 'pointer',
  border: '2px dashed #02C3BD',
  padding: '30px 0',
  height: '20px',
  width: '400px',
  marginBottom: '15px',
  overflow: 'scroll',
  lineHeight: '1.3',
};

const zoneActive = {
  background: 'repeating-linear-gradient(-45deg, white, white 10px, rgba(0, 0, 0, 0.05) 10px, rgba(0, 0, 0, 0.05) 20px)',
};

function Upload({ uploadInput }) {
  const [ key, setKey ] = useState('');
  const [ uploaded, setUploaded ] = useState('select/drag files');
  const [ files, setFiles ] = useState([]);
  const [ hover, setHover ] = useState(false);
  const onDrop = useCallback(acceptedFiles => {
    setUploaded(acceptedFiles.map(f => f.name).join(', '))
    setFiles(acceptedFiles)
  }, []);
  const { getRootProps, isDragActive } = useDropzone({ onDrop });
  const style = useMemo(() => ({
    ...zone,
    ...(isDragActive || hover ? zoneActive : {}),
  }));

  const handleUpload = (e) => {
    e.preventDefault();

    const data = new FormData();
    data.append('file', files[ 0 ]);
    console.log(uploadInput.files)

    fetch('http://35.171.21.255:80/data', {
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
        }) }>
          <span>{ uploaded }</span>
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
        <Button loading={ false } label={ Upload.name }/>
        <span className="key">{ key }</span>
      </form>
      <div id="or"><span>OR</span></div>
    </div>
  )
}

export default Upload