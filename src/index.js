import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import { Router, browserHistory } from 'react-router';
import { createStore, combineReducers, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';
import createLogger from 'redux-logger';
import { syncHistoryWithStore, routerReducer } from 'react-router-redux';

import './index.less';

import { routes } from './routes';
import { harbor } from './data/reducers';

function configureStore(initialState) {
  const logger = createLogger();

  const createStoreWithMiddleware = applyMiddleware(
    thunk,
    logger
  )(createStore);

  let combinedReducers = combineReducers({
    harbor,
    routing: routerReducer,
  });

  return createStoreWithMiddleware(combinedReducers, initialState);
}

const store = configureStore();
const history = syncHistoryWithStore(browserHistory, store);

let app = (
  <Provider store={store}>
    <Router history={history}>
      {routes}
    </Router>
  </Provider>
);

ReactDOM.render(app, document.querySelector('#app'));
