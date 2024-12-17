// src/api/user.js

import axios from '~/axios';


// 获取用户列表
export function apiGetDBList(data) {
  return axios.post('/api/v1/db/list', data);
}

// apiAddUser
// 添加用户
export function apiAddDB(data) {
  return axios.post('/api/v1/db', data);
}

// apiUpdateUser
// 更新用户
export function apiUpdateDB(data) {
  return axios.put('/api/v1/db', data);
}

// apiDeleteUser
// 删除用户
export function apiDeleteDB(id) {
  return axios.delete(`/api/v1/db/${id}`);
}