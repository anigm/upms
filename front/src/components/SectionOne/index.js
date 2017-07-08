import React from 'react';
import { Icon } from 'antd';
import { connect } from 'dva';

function SectionOne() {
  return (
    <div
      style={{
        position: 'absolute',
        height: '100%',
        width: '100%',
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
      ><Icon type="smile" />  SectionOne</h1>
    </div>
  );
}

export default connect()(SectionOne);
