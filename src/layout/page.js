import React from 'react';
import { connect } from 'react-redux';
import { browserHistory, Link } from 'react-router';
import { Nav, NavItem, Navbar } from 'react-bootstrap';

import './page.less';

class Page extends React.Component {
  render() {
    const {
      content,
      user,
    } = this.props;

    return (
      <div>
        <Navbar staticTop>
          <Navbar.Header>
            <Navbar.Brand>
              <a onClick={() => {browserHistory.push('/')}}>Harbor</a>
            </Navbar.Brand>
          </Navbar.Header>
          <Nav>
            <NavItem eventKey={1} onClick={() => {browserHistory.push('/')}}>Dashboard</NavItem>
            <NavItem eventKey={2} onClick={() => {browserHistory.push('/profile')}}>Account</NavItem>
          </Nav>
        </Navbar>
        {content}
      </div>
    )
  }
}

export default connect(
  (state) => {
    return { user: state.user };
  }
)(Page);
