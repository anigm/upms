import React, { PropTypes } from 'react';
import { injectIntl, defineMessages } from 'react-intl';
import { Menu, Icon, Dropdown, Button, Row } from 'antd';
import { IndexLink, Link } from 'dva/router';
import { connect } from 'dva';
import './Header.less';


const messages = defineMessages({
  Index: {
    id: 'Header.Index',
    defaultMessage: 'Index',
  },
  SectionOne: {
    id: 'Header.SectionOne',
    defaultMessage: 'SectionOne',
  },
  SectionTwo: {
    id: 'Header.SectionTwo',
    defaultMessage: 'SectionTwo',
  },
  Logo: {
    id: 'Header.Logo',
  },
});

function Header({ intl }) {
  function NavMenu(props) {
    return (
      <Menu {...props}>
        <Menu.Item key="mail">
          <IndexLink to="/"><Icon type="mail" />
            {intl.formatMessage(messages.Index)}
          </IndexLink>
        </Menu.Item>
        <Menu.Item key="SectionOne">
          <Link to="/SectionOne"><Icon type="appstore" />
            {intl.formatMessage(messages.SectionOne)}
          </Link>
        </Menu.Item>
        <Menu.Item key="SectionTwo">
          <Link to="/SectionTwo"><Icon type="appstore" />
            {intl.formatMessage(messages.SectionTwo)}
          </Link>
        </Menu.Item>
        {/* <Menu.SubMenu title={<span><Icon type="setting" />Navigation Three - Submenu</span>}>
        <Menu.ItemGroup title="Item 1">
          <Menu.Item key="setting:1">Option 1</Menu.Item>
          <Menu.Item key="setting:2">Option 2</Menu.Item>
        </Menu.ItemGroup>
        <Menu.ItemGroup title="Item 2">
          <Menu.Item key="setting:3">Option 3</Menu.Item>
          <Menu.Item key="setting:4">Option 4</Menu.Item>
        </Menu.ItemGroup>
      </Menu.SubMenu>*/}
        <Menu.Item key="alipay">
          <a href="https://ant.design" target="_blank" rel="noopener noreferrer">ant-design</a>
        </Menu.Item>
      </Menu>
    );
  }


  return (
    <Row type="flex" justify="space-between" align="middle">
      <div className="nav-logo">
        <IndexLink to="/" className="nav-logo-link">
          <h1>{intl.formatMessage(messages.Logo)}</h1>
        </IndexLink>
      </div>
      {NavMenu({ mode: 'horizontal', className: 'nav-menu' })}
      <div className="nav-dropdown">
        <Dropdown overlay={NavMenu()} placement="bottomLeft" trigger={['click']}>
          <Button><Icon type="bars" /></Button>
        </Dropdown>
      </div>
    </Row>
  );
}

Header.propTypes = {
  intl: PropTypes.object.isRequired,
};

export default injectIntl(connect()(Header));
