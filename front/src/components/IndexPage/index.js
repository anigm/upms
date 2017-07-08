import React from 'react';
import { injectIntl, defineMessages } from 'react-intl';
import { Icon, DatePicker } from 'antd';
import { connect } from 'dva';
import './index.less';

const messages = defineMessages({
  PickDate: {
    id: 'Index.PickDate',
    defaultMessage: 'Index',
  },
});

function IndexPage({ intl }) {
  return (
    <div
      style={{
        position: 'absolute',
        height: '100%',
        width: '100%',
        textAlign: 'center',
      }}
    >
      <div
        style={{
          position: 'relative', /* 脱离文档流*/
          top: '50%', /* 偏移*/
          transform: 'translateY(-50%)',
          margin: '0 auto',
        }}
      >
        <h1><Icon type="smile" />
          {intl.formatMessage(messages.PickDate)}
        </h1>
        <DatePicker />
      </div>
    </div>
  );
}

export default injectIntl(connect()(IndexPage));
