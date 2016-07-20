import React from 'react';
import { IndexRoute, Route } from 'react-router';
import Page from './layout/page';
import Pages from './pages/index';

export const routes = (
  <Route path="/" component={Page}>
    <IndexRoute components={{
      body: Pages.Dashboard.Body,
    }}/>
    <Route path="/profile" components={{
      body: Pages.Profile.Body,
    }}/>
  </Route>
);
