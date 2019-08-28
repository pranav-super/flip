import React, { useState } from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

import './Retrieve.css';
import Button from './Button';
import { updateRetrieved } from '../reducers/retrieved';

function Retrieve({ retrieve, history }) {
  const [ loading, setLoading ] = useState(false);
  const [ token, setToken ] = useState('');

  return (
    <div className="Retrieve">
      <input maxLength="6" type="text" placeholder="000000" onChange={ e => setToken(e.target.value) }/>
      <Button loading={ loading } label={ Retrieve.name } onClick={ () => {
        setLoading(true);
        retrieve(token).then(() => {
          setLoading(false);
          history.push(token);
        });
      } }/>
    </div>
  )
}

const mapDispatchToProps = ({
  retrieve: updateRetrieved,
});

export default withRouter(connect(null, mapDispatchToProps)(Retrieve));