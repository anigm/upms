import { addLocaleData } from 'react-intl';

function getZhLocale() {
  return {
    messages: require('../locales/zh.json'),
    antd: null,
    locale: 'zh-Hans-CN',
    data: require('react-intl/locale-data/zh'),
  };
}

function getEnLocale() {
  return {
    messages: require('../locales/en.json'),
    antd: require('antd/lib/locale-provider/en_US'),
    locale: 'en',
    data: require('react-intl/locale-data/en'),
  };
}

export default {
  namespace: 'app',
  state: {
    user: {},
    login: false,
    antd: null,
    appLocale: undefined,
  },
  reducers: {
    switchLocale(state) {
      console.log('switchLocale');
      const language = (navigator.language || navigator.browserLanguage).toLowerCase();
      addLocaleData(require('react-intl/locale-data/zh'), require('react-intl/locale-data/en'));
      if ((!state.appLocale && (language === 'zh-cn' || language === 'zh')) || (state.appLocale && state.appLocale.locale && state.appLocale.locale === 'en')) {
        return { ...state, ...{ appLocale: getZhLocale() },
        };
      } else {
        return { ...state, ...{ appLocale: getEnLocale() } };
      }
    },
  },
  effects: {
    // *locale() {
    // },
  },
  subscriptions: {
    setup({ dispatch }) {
      dispatch({ type: 'switchLocale' });
    },
  },
};
