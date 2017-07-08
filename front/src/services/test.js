import request from '../utils/request';

export function GetUser() {
  return request('/public/index.php');
}
