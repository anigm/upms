import React from 'react';
import { connect } from 'dva';
import { Link } from 'dva/router';
import { Table, Popconfirm, Button, Icon } from 'antd';
import { queryUsers } from '../../../services/users';

const { Column } = Table;

class UserTable extends React.Component {
  state={
    users: undefined,
  }
  componentDidMount() {
    queryUsers().then((data) => {
      console.log(data);
      this.setState(() => {
        return { users: data.data };
      });
    });
  }
  render() {
    return (
      <div>
        <div>
          <Button>
            <Link to="/Register"><Icon type="plus" />添加用户</Link>
          </Button>
        </div>
        <Table dataSource={this.state.users}>
          <Column
            title="ID"
            dataIndex="ID"
            key="ID"
          />
          <Column
            title="账号"
            dataIndex="Account"
            key="Account"
          />
          <Column
            title="昵称"
            dataIndex="Name"
            key="Name"
          />
          <Column
            title="删除"
            key="delete"
            render={(text, record) => ( // eslint-disable-line 
              <span>
                <Popconfirm record={record} title="确认删除这个用户吗?" okText="是" cancelText="否">
                  <a>删除</a>
                </Popconfirm>
              </span>
            )}
          />
          <Column
            title="修改"
            key="modify"
            render={(text, record) => ( // eslint-disable-line 
              <span>
                <a >修改</a>
              </span>
            )}
          />
        </Table>

      </div>
    );
  }
}

export default connect()(UserTable);
