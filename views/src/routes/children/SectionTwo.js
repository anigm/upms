// import React from 'react';
// import { RegisterModel } from '../../utils/common';

export default {
  path: 'SectionTwo',
  getComponent(location, callback) {
    require.ensure([], (require) => {
      callback(null, require('../../components/SectionTwo'));
    }, 'SectionTwo');
  },
};
