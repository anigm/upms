import React from 'react';
import { Layout, Icon } from 'antd';
import { connect } from 'dva';

function Footer() {
  return (
    <Layout>
      <Layout.Content
        style={{
          height: 'inherit',
          textAlign: 'center',
        }}
      >
        <h1
          style={{
            position: 'relative', /* 脱离文档流*/
            top: '50%', /* 偏移*/
            transform: 'translateY(-50%)',
            margin: '0 auto',
          }}
        ><Icon type="frown-o" />  404 Not Found</h1>
      </Layout.Content>
    </Layout>
  );
}

export default connect()(Footer);
