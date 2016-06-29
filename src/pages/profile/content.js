import React from 'react';
import { connect } from 'react-redux';
import Request from 'superagent';
import { Link } from 'react-router';

import './index.less';

class Content extends React.Component {
  render() {
    return (
      <h1>
        Profile
      </h1>
    )
  }
}

export default connect(
  (state) => {
    return {
      user: state.user,
    };
  }
)(Content);
