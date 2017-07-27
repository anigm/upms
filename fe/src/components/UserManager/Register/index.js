import React from 'react';
import { connect } from 'dva';

import { Form, Input, Tooltip, Icon, Button, message } from 'antd';

import { register } from '../../../services/example';

const FormItem = Form.Item;

class RegistrationForm extends React.Component {
  state = {
    confirmDirty: false,
    autoCompleteResult: [],
  };
  handleSubmit = (e) => {
    e.preventDefault();
    this.props.form.validateFieldsAndScroll((err, values) => {
      if (!err) {
        console.log('Received values of form: ', values);
        let data = '';
        for (const name in values) {
          if (name !== undefined) {
            data += `${name}=${values[name]}&`;
          }
        }
        register(data).then((d) => {
          if (d.status !== 0) {
            message.error(`创建失败：${d.info}`);
          } else {
            message.info('创建成功');
          }
        });
      }
    });
  }
  handleConfirmBlur = (e) => {
    const value = e.target.value;
    this.setState({ confirmDirty: this.state.confirmDirty || !!value });
  }
  checkPassword = (rule, value, callback) => {
    const form = this.props.form;
    if (value && value !== form.getFieldValue('Password')) {
      callback('两次输入的密码不一致');
    } else {
      callback();
    }
  }
  checkConfirm = (rule, value, callback) => {
    const form = this.props.form;
    console.log(form);
    if (value && this.state.confirmDirty) {
      form.validateFields(['confirm'], { force: true });
    }
    callback();
  }
  render() {
    const { getFieldDecorator } = this.props.form;

    const formItemLayout = {
      labelCol: {
        xs: { span: 24 },
        sm: { span: 6 },
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 14 },
      },
    };
    const tailFormItemLayout = {
      wrapperCol: {
        xs: {
          span: 24,
          offset: 0,
        },
        sm: {
          span: 14,
          offset: 6,
        },
      },
    };
    return (
      <Form id="register-form" onSubmit={this.handleSubmit}>
        <FormItem
          {...formItemLayout}
          label="账号"
          hasFeedback
        >
          {getFieldDecorator('Account', {
            rules: [{
              required: true, message: '请输入账号',
            }],
          })(
            <Input />,
          )}
        </FormItem>
        <FormItem
          {...formItemLayout}
          label="密码"
          hasFeedback
        >
          {getFieldDecorator('Password', {
            rules: [{
              required: true, message: '请输入密码',
            }, {
              validator: this.checkConfirm,
            }],
          })(
            <Input type="password" />,
          )}
        </FormItem>
        <FormItem
          {...formItemLayout}
          label="确认密码"
          hasFeedback
        >
          {getFieldDecorator('confirm', {
            rules: [{
              required: true, message: '请确认密码!',
            }, {
              validator: this.checkPassword,
            }],
          })(
            <Input type="password" onBlur={this.handleConfirmBlur} />,
          )}
        </FormItem>
        <FormItem
          {...formItemLayout}
          label={(
            <span>
              昵称&nbsp;
              <Tooltip title="其他人看到的名称">
                <Icon type="question-circle-o" />
              </Tooltip>
            </span>
          )}
          hasFeedback
        >
          {getFieldDecorator('Name', {
            rules: [{ required: true, message: '请输入昵称!', whitespace: true }],
          })(
            <Input />,
          )}
        </FormItem>
        <FormItem {...tailFormItemLayout}>
          <Button type="primary" htmlType="submit">创建</Button>
        </FormItem>
      </Form>
    );
  }
}

export default connect()(Form.create()(RegistrationForm));
