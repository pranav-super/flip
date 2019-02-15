import React, {Component} from 'react';
import './Retrieve.css';
import Button from './Button';

class Retrieve extends Component {
  render() {
    return (
      <div className="Retrieve">
        <form>
          <input maxLength="5" type="text" placeholder="00000"/>
          <Button label={Retrieve.name}/>
        </form>
      </div>
    )
  }
}

export default Retrieve