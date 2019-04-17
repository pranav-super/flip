import React from 'react';

function Download({ match }) {
  return <div>{match.params.token}</div>
}

export default Download;