import request, { form } from '../utils/request';

export async function query() {
  return request('/api/users');
}

export async function register(data) {
  return form('/api/register', data);
}

