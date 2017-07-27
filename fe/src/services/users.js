import request from '../utils/request';

export async function queryUsers() {
  return request('/api/user');
}
