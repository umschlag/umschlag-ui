import React from 'react';
import FontAwesome from 'react-fontawesome';
import { branch } from 'baobab-react/higher-order';
import { browserHistory, Link } from 'react-router';
import { Navbar, Nav, NavItem } from 'react-bootstrap';
import Image from '../../components/image';
import { events } from '../../actions/events';

import './index.less';

class Page extends React.Component {
  render() {
    const {
      body,
    } = this.props;

    return (
      <div>
        <Navbar>
          <Navbar.Header>
            <Navbar.Brand>
              <Link to="/">
                <Image src="/images/logo.svg" alt="Umschlag"/>
              </Link>
            </Navbar.Brand>
          </Navbar.Header>
          <Navbar.Header className="pull-right">
            <Nav>
              <NavItem eventKey={1} onClick={() => {browserHistory.push('/')}}>
                <FontAwesome name="dashboard" size="lg"/>
                {' '}
                <span className="hidden-xs">Dashboard</span>
              </NavItem>
              <NavItem eventKey={2} onClick={() => {browserHistory.push('/profile')}}>
                <FontAwesome name="user" size="lg"/>
                {' '}
                <span className="hidden-xs">Profile</span>
              </NavItem>
            </Nav>
          </Navbar.Header>
        </Navbar>
        { body }
      </div>
    )
  }
}

export default branch({
  state: ['pages']
}, Page);
