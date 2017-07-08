// import { RegisterModel } from '../../utils/common';

export default {
  path: 'SectionOne',
  getComponent(location, callback) {
    require.ensure([], (require) => {
      callback(null, require('../../components/SectionOne'));
    }, 'SectionOne');
  },
};
