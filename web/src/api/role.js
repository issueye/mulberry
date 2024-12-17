// src/api/user.js

import axios from '~/axios';


// 获取用户列表
export function apiGetRoleList(data) {
  return axios.post('/api/v1/role/list', data);
}

// apiAddUser
// 添加用户
export function apiAddRole(data) {
  return axios.post('/api/v1/role/add', data);
}

// apiUpdateUser
// 更新用户
export function apiUpdateRole(data) {
  return axios.put('/api/v1/role/update', data);
}

// apiDeleteUser
// 删除用户
export function apiDeleteRole(id) {
  return axios.delete(`/api/v1/role/delete/${id}`);
}

