// import React from 'react';
// import { RegisterModel } from '../utils/common';

const RootRouter = [
  {
    path: '/',
    getChildRoutes(location, callback) {
      require.ensure([], (require) => {
        callback(null, require('./children'));
      }, 'children');
    },
    getIndexRoute(location, callback) {
      require.ensure([], (require) => {
        callback(null, { component: require('../components/IndexPage') });
      }, 'IndexPage');
    },
    getComponent(location, callback) {
      callback(null, require('../components/App'));
    },
  },
];

export default RootRouter;
