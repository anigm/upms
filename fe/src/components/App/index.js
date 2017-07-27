import React from 'react';
import { connect } from 'dva';
import { IndexLink, Link } from 'dva/router';
import { Layout, Menu, Icon } from 'antd';

import './index.less';

const { Header, Sider, Content } = Layout;

class App extends React.Component {
  state = {
    collapsed: false,
  };
  toggle = () => {
    this.setState({
      collapsed: !this.state.collapsed,
    });
  }
  render() {
    return (
      <Layout className="app-layout">
        <Sider
          trigger={null}
          collapsible
          collapsed={this.state.collapsed}
        >
          <IndexLink to="/">
            <div className="logo" />
          </IndexLink>
          <Menu theme="dark" mode="inline" defaultSelectedKeys={['1']}>
            <Menu.SubMenu key="user" title={<span><Icon type="user" /><span>用户管理</span></span>}>
              <Menu.Item key="users">
                <Link to="users"> <Icon type="user" /><span>用户列表</span></Link>
              </Menu.Item>
              <Menu.Item key="register">
                <Link to="Register"> <Icon type="user-add" /><span>用户注册</span></Link>
              </Menu.Item>
            </Menu.SubMenu>
            <Menu.Item key="2">
              <Icon type="team" />
              <span>角色管理</span>
            </Menu.Item>
            <Menu.Item key="3">
              <Icon type="solution" />
              <span>权限管理</span>
            </Menu.Item>
          </Menu>
        </Sider>
        <Layout>
          <Header style={{ background: '#fff', padding: 0 }}>
            <Icon
              className="trigger"
              type={this.state.collapsed ? 'menu-unfold' : 'menu-fold'}
              onClick={this.toggle}
            />
          </Header>
          <Content style={{ margin: '24px 16px', padding: 24, background: '#fff', minHeight: 280 }}>
            {this.props.children}
          </Content>
        </Layout>
      </Layout>
    );
  }
}

export default connect()(App);
