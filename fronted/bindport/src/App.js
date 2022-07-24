/**
 * Copyright (c) OpenSpug Organization. https://github.com/openspug/spug
 * Copyright (c) <spug.dev@gmail.com>
 * Released under the AGPL-3.0 License.
 */
import React, { Component } from 'react';
import { Switch, Route } from 'react-router-dom';
import Home from './pages/home';

class App extends Component {
  render() {
    return (
      <Switch>
        <Route path="/" exact component={Home} />
      </Switch>
    );
  }
}

export default App;
