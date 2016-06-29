import React from 'react';
import { IndexRoute, Route } from 'react-router';

import Page from './layout/page';
import Pages from './pages/index';

export const routes = (
  <Route path="/" component={Page}>
    <IndexRoute components={{
      content: Pages.Dashboard.Content,
    }}/>
    <Route path="/profile" components={{
      content: Pages.Profile.Content,
    }}/>
  </Route>
);
