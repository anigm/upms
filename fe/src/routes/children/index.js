export default [
  {
    path: '/Register',
    getComponent(location, callback) {
      require.ensure([], (require) => {
        callback(null, require('../../components/UserManager/Register'));
      }, 'Register');
    },
  },
  {
    path: '/users',
    getComponent(location, callback) {
      require.ensure([], (require) => {
        callback(null, require('../../components/UserManager/UserTable'));
      }, 'UserTable');
    },
  },
];
