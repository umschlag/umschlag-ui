import React from 'react';
import { branch } from 'baobab-react/higher-order';

import './index.less';

class Body extends React.Component {
  render() {
    return (
      <div className="container">
        <h1>
          Profile
        </h1>
      </div>
    );
  }
}

export default branch({

}, Body);
