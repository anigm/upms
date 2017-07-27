import React from 'react';
import { connect } from 'dva';
import { Table } from 'antd';
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
      <Table dataSource={this.state.users}>
        <Column
          title="ID"
          dataIndex="ID"
          key="ID"
        />
        <Column
          title="昵称"
          dataIndex="Name"
          key="Name"
        />
      </Table>
    );
  }
}

export default connect()(UserTable);
