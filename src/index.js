import './index.less';

import React from 'react';
import ReactDOM from 'react-dom';
import superagent from 'superagent';
import { root } from 'baobab-react/higher-order';
import { Router, browserHistory } from 'react-router';
import { events } from './actions/events';
import { tree } from './actions/tree';
import { routes } from './routes';


window.tree = tree;
window.events = events;

class App extends React.Component {
  render() {
    return (
      <Router history={browserHistory}>
        {routes}
      </Router>
    );
  }
}

ReactDOM.render(
  React.createElement(root(tree, App)),
  document.getElementById('umschlag')
);
