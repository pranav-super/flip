import React, { Component } from 'react';
import './Upload.css';
import Button from './Button';

class Upload extends Component {
  constructor(props) {
    super(props);

    this.state = {
      key: '',
      uploaded: ''
    };

    this.handleUpload = this.handleUpload.bind(this);
  }

  handleUpload(e) {
    e.preventDefault();

    const data = new FormData();
    data.append('sampleFile', this.uploadInput.files[0]);

    fetch('http://localhost:80/upload', {
      method: 'POST',
      body: data,
    }).then((response) => {
      response.json().then((body) => {
        if (body.Key) {
          this.setState({ key: "Retrieval key: " + body.Key });
        }
      });
    }).catch((err) => {
      this.setState({ key: "Please try again!" });
    });
  }

  render() {
    return (
      <div className="Upload">
        <form onSubmit={ this.handleUpload }>
          <label htmlFor="upload">
            <span>select/drag files</span>
            <span>{ this.uploaded }</span>
          </label>
          <input id="upload" ref={
            (ref) => {
              this.uploadInput = ref;
            }
          } type="file" multiple/>
          <Button label={ Upload.name }/>
          <span className="key">{ this.state.key }</span>
        </form>
        <div id="or"><span>OR</span></div>
      </div>
    )
  }
}

export default Upload