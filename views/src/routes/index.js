// import React from 'react';
import { RegisterModel } from '../utils/common';

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
      require.ensure([], (require) => {
        RegisterModel(require('../models/app'));
        callback(null, require('../components').default);
      }, 'app');
    },
  },
  {
    path: '*',
    getComponent(location, callback) {
      require.ensure([], (require) => {
        RegisterModel(require('../models/test'));
        callback(null, require('../components').NotFound);
      }, 'test');
    },
  },
];

export default RootRouter;
