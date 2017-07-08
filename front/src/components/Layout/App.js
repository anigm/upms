import React from 'react';
import { IntlProvider } from 'react-intl';
import { Layout, LocaleProvider } from 'antd';
import { connect } from 'dva';
import Header from './Header';
import Footer from './Footer';
import './App.less';

function App({ app, children }) {
  return (
    <LocaleProvider locale={app.appLocale.antd}>
      <IntlProvider locale={app.appLocale.locale} messages={app.appLocale.messages}>
        <Layout>
          <Layout.Header>
            <Header />
          </Layout.Header>
          <Layout.Content>
            {children}
          </Layout.Content>
          <Layout.Footer>
            <Footer />
          </Layout.Footer>
        </Layout>
      </IntlProvider>
    </LocaleProvider>
  );
}

export default connect(({ app }) => ({ app }))(App);
