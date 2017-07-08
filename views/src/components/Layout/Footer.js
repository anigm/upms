import React from 'react';
import { injectIntl, FormattedMessage } from 'react-intl';
import { Row, Button } from 'antd';
import { connect } from 'dva';

function Footer({ dispatch }) {
  return (
    <Row
      style={{
        textAlign: 'center',
        margin: '2%',
      }}
    >
      <h4>@2017   Enorzw</h4>
      <Button style={{ float: 'right' }} onClick={() => { dispatch({ type: 'app/switchLocale' }); }}>
        <FormattedMessage id="Footer.Language" />
      </Button>
    </Row>
  );
}

export default injectIntl(connect()(Footer));
